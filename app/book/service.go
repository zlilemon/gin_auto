package book

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/app/account"
	"github.com/zlilemon/gin_auto/app/billing"
	"github.com/zlilemon/gin_auto/app/pay"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"strconv"
	"time"
)

type IService interface {
	// SCreateOrderId 生成订单号
	SCreateOrderId() (err error)

	// SOrder 下单
	SOrder() (err error)
	SPrepay(c *gin.Context, openid string, tranAmt int64) (err error)
	SAutoShutDownDevice() (err error)
}

type Service struct {
	//repo  		IRepository
}

var BookService = new(Service)

func (s *Service) SCreateOrderId(c *gin.Context) (orderId string, err error) {
	log.Infof("SCreateOrderId")
	currentTime := time.Now()
	orderId = strconv.FormatInt(currentTime.Unix(), 10)
	log.Infof("create orderId:%s", orderId)
	return
}

func (s *Service) SOrder(c *gin.Context, req BookReq, resp *comm.Result) (err error) {
	log.Infof("SOrder - ")

	// 座位状态
	seatOrderStatusReq := SeatOrderStatusReq{}

	// 生成内部订单号
	outTradeNo, err := BookService.SCreateOrderId(c)

	if err != nil {
		log.Errorf("create order id error, errMsg:%s", err)
		return
	}

	payMethod := req.PayMethod
	log.Infof("payMethod : %s", payMethod)

	// 先创建订单
	orderInsertReq := billing.InsertBillingReq{}
	orderInsertResp := billing.InsertBillingResp{}

	orderInsertReq.OutTradeNo = outTradeNo
	orderInsertReq.ChannelOrderNo = ""
	orderInsertReq.OrderType = req.OrderType
	orderInsertReq.PayMethod = req.PayMethod
	orderInsertReq.OpenId = req.OpenId
	orderInsertReq.StoreId = req.ShopId
	orderInsertReq.StoreName = req.ShopName
	orderInsertReq.SeatId = req.SeatId
	orderInsertReq.SeatName = req.SeatName
	orderInsertReq.SeatType = req.SeatType
	orderInsertReq.BookBeginDate = req.BookBeginDate
	orderInsertReq.BookEndDate = req.BookEndDate
	orderInsertReq.BookBeginTime = req.BookBeginTime
	orderInsertReq.BookEndTime = req.BookEndTime
	orderInsertReq.BookDuration = req.BookDuration
	orderInsertReq.Currency = req.Currency
	orderInsertReq.Amount = req.Amount
	orderInsertReq.TranTime = req.TranTime
	orderInsertReq.BillStatus = "VALID"
	orderInsertReq.PayStatus = "INIT"
	orderInsertReq.PayInfo = req.PayInfo

	err = billing.BillingService.SCreateBilling(c, orderInsertReq, &orderInsertResp)
	if err != nil {
		log.Errorf("failed to insert order - %+v", err)
		return
	}
	log.Infof("SCreateBilling Success, outTradeNo:%s, openid:%s", outTradeNo, req.OpenId)

	// 预定座位类型，先锁住座位的状态信息，状态为 lock
	if req.OrderType == "bookConsume" {
		seatOrderStatusReq.StoreId = req.ShopId
		seatOrderStatusReq.StoreName = req.ShopName
		seatOrderStatusReq.SeatId = req.SeatId
		seatOrderStatusReq.SeatName = req.SeatName
		seatOrderStatusReq.BookBeginTime = fmt.Sprintf("%s %s", req.BookBeginDate, req.BookBeginTime)
		seatOrderStatusReq.BookEndTime = fmt.Sprintf("%s %s", req.BookEndDate, req.BookEndTime)
		seatOrderStatusReq.ToStatus = "LOCK"

		err = BookService.SUpdateSeatStatus(c, seatOrderStatusReq)
		if err != nil {
			log.Errorf("SUpdateSeatStatus error, errMsg:%s", err)
			return
		}
	}

	// 更新订单状态为成功
	// 更新对应订单号为支付成功，并更新渠道回调订单号
	bookOrderUpdateReq := billing.UpdateBillingReq{}
	bookOrderUpdateReq.OutTradeNo = outTradeNo
	bookOrderUpdateReq.OpenId = req.OpenId

	// 根据支付方式，选择调起不同的支付接口
	wxpayReq := pay.WxPayReq{}
	wxpayResp := pay.WxPayResp{}
	if payMethod == "wechat" {
		// 通过微信支付
		wxpayReq.OpenId = req.OpenId
		wxpayReq.Amount = req.Amount
		wxpayReq.PayInfo = req.PayInfo
		wxpayReq.OutTradeNo = outTradeNo

		// 调用微信支付接口
		err = pay.PayService.SWxPayV3(c, wxpayReq, &wxpayResp)
		//if err != nil {
		//	log.Errorf("SWxPayV3 error, errMsg:%s", err)
		//	return
		//}
	} else if payMethod == "account" {
		// 先判断账户余额是否足够，不足够的直接返回失败
		accountBalanceReq := account.BalanceReq{}
		accountBalanceResp := account.BalanceResp{}

		accountBalanceReq.OpenId = req.OpenId
		err = account.AccountService.SAccountGet(c, accountBalanceReq, &accountBalanceResp)
		if err != nil {
			log.Errorf("SAccountGet error, errMsg:%s", err)
			bookOrderUpdateReq.PayStatus = "FAIL"
			updateErr := billing.BillingService.SUpdateBilling(c, bookOrderUpdateReq)
			if updateErr != nil {
				log.Errorf("SUpdateOrder error, err_msg:%v", updateErr)
				return
			}
			return
		}

		accountBalance := accountBalanceResp.Amount
		if accountBalance/100 < req.Amount {
			errMsg := fmt.Sprintf("账户余额不足，请选择其它支付方式。账户余额:[%d]， 支付金额:[%d]",
				accountBalance/100, accountBalanceResp.Amount)
			log.Errorf("%s", errMsg)

			err = errors.New(errMsg)
			bookOrderUpdateReq.PayStatus = "FAIL"
			updateErr := billing.BillingService.SUpdateBilling(c, bookOrderUpdateReq)
			if updateErr != nil {
				log.Errorf("SUpdateOrder error, err_msg:%v", updateErr)
				return
			}
			return
		}

		// 通过余额支付
		accountPayReq := account.ConsumeReq{}
		accountPayResp := account.ConsumeResp{}

		accountPayReq.OpenId = req.OpenId
		accountPayReq.Currency = req.Currency
		accountPayReq.Amount = req.Amount
		accountPayReq.OutTradeNO = outTradeNo

		err = account.AccountService.SAccountConsume(c, accountPayReq, &accountPayResp)
	}

	if err != nil {
		// 有异常，退出处理
		log.Errorf("SOrder -- pay error, err:%s", err)
		// 更新订单为失败
		if req.OrderType == "bookConsume" {
			bookOrderUpdateReq.PayStatus = "FAIL"
		}
		// 把座位的状态释放掉
		seatOrderStatusReq.ToStatus = "UNLOCK"

		err = BookService.SUpdateSeatStatus(c, seatOrderStatusReq)
		if err != nil {
			log.Errorf("SUpdateSeatStatus error, errMsg:%s", err)
			return
		}

	} else {
		bookOrderUpdateReq.PayStatus = "SUCCESS"
	}

	//  根据充值结果，跟进订单的同步充值状态
	err = billing.BillingService.SUpdateBilling(c, bookOrderUpdateReq)
	if err != nil {
		log.Errorf("SUpdateOrder error, err_msg:%v", err)
		return
	}

	// 把座位的状态改成SUCCESS（成功预定）
	seatOrderStatusReq.ToStatus = "SUCCESS"

	err = BookService.SUpdateSeatStatus(c, seatOrderStatusReq)
	if err != nil {
		log.Errorf("SUpdateSeatStatus error, errMsg:%s", err)
		return
	}

	log.Infof("SOrder Success")
	resp.Data = wxpayResp
	return
}

func (s *Service) SGetOrder(c *gin.Context, req BookOrderReq, resp *comm.Result) (err error) {
	log.Infof("SOrder - ")
	//bookOrderRespModel, err := BookRepository.GetOrder(c, req)

	//orderResp := make([]*BookOrderResp, 0)
	orderReq := billing.BillingInfoReq{}
	orderResp := make([]*billing.BillingInfoResp, 0)

	orderReq.OpenId = req.OpenId
	orderReq.OrderType = req.OrderType
	orderReq.OutTradeNo = req.OutTradeNo
	orderReq.BeginDate = req.BeginDate
	orderReq.EndDate = req.EndDate
	orderReq.NotifyStatus = req.NotifyStatus
	orderReq.Page = req.Page
	orderReq.PageSize = req.PageSize

	//orderResp, err = BookRepository.GetOrder(c, req)
	orderResp, err = billing.BillingRepository.GetOrderInfo(c, orderReq)
	if err != nil {
		log.Errorf("failed to SGetOrder - %+v", err)
		return
	}

	resp.Data = orderResp
	//bookOrderRespList := make([]OrderDetail, 0)
	/*
		for _, v := range bookOrderRespModel {
			newItem := OrderDetail{}
			newItem.OpenId = v.OpenId
			newItem.ChannelOrderNo = v.ChannelOrderNo
			newItem.OutTradeNo = v.OutTradeNo
			newItem.StoreId = v.ShopId
			newItem.SeatId = v.SeatId
			newItem.SeatType = v.SeatType
			newItem.BookBeginDate = v.BookBeginDate
			newItem.BookEndDate = v.BookEndDate
			newItem.BookBeginTime = v.BookBeginTime
			newItem.BookEndTime = v.BookEndTime
			newItem.BookDuration = v.BookDuration
			newItem.PayMethod = v.PayMethod
			newItem.Amount = v.Amount
			newItem.PayInfo = v.PayInfo

			log.Infof("shop_id : %s", newItem.OpenId)
			bookOrderRespList = append(bookOrderRespList, newItem)
		}

		resp.Data = bookOrderRespList
	*/
	return
}

func (s *Service) SGetOrderDivideByTime(c *gin.Context, req BookOrderDivideByTimeReq, resp *comm.Result) (err error) {
	log.Infof("SGetOrderDivideByTime - ")
	orderReq := billing.BillingInfoDivideByTimeReq{}
	orderResp := make([]*billing.BillingInfoResp, 0)

	orderReq.OpenId = req.OpenId
	orderReq.FutureFlag = req.FutureFlag
	orderReq.Page = req.Page
	orderReq.PageSize = req.PageSize

	orderResp, err = billing.BillingRepository.GetOrderInfoDivideByTime(c, orderReq)
	if err != nil {
		log.Errorf("failed to GetOrderInfoDivideByTime - %+v", err)
		return
	}
	resp.Data = orderResp
	return
}

/*
func (s *Service) SUpdateOrder(c *gin.Context, req BookOrderUpdateReq) (err error) {
	log.Infof("SUpdateOrder ")

	var updateOrderReq = billing.UpdateBillingReq{}
	updateOrderReq
	err = BookRepository.UpdateOrder(c, req)
	err = billing.BillingRepository.UpdateOrderInfo(c, req)
	if err != nil {
		log.Errorf("failed to inert order - %+v", err)
		return
	}

	log.Infof("SUpdateOrder Success")

	return
}

*/

func (s *Service) SAutoShutDownDevice() (err error) {
	// 通过定时任务出发，查询当前订单状态，到时间了就 自动关闭设备
	fmt.Println("hello SAutoShutDownDevice")
	// 获取订单的最晚结束时间

	// 根据订单，找到要关闭的设备
	// 关闭设备异常要记录告警信息

	return
}

// 更新座位状态
func (s *Service) SUpdateSeatStatus(c *gin.Context, seatOrderStatusReq SeatOrderStatusReq) (err error) {
	log.Infof("SUpdateSeatStatus - ")

	//orderResp := make([]*BookSeatOrderStatus, 0)

	//先查询该预定信息，在预定信息表中是否存在
	seatOrderStatusModel, err := BookRepository.GetSeatOrderStatus(c, seatOrderStatusReq)
	lenSeatOrderStatus := len(seatOrderStatusModel)
	if lenSeatOrderStatus == 0 {
		// 没有预定信息，新增插入预定信息
		err = BookRepository.InsertSeatOrderStatus(c, seatOrderStatusReq)
	} else if lenSeatOrderStatus == 1 {
		// 更新预定信息
		err = BookRepository.UpdateSeatOrderStatus(c, seatOrderStatusReq)
	} else if lenSeatOrderStatus > 1 {
		// 有异常，需要告警

	}

	return
}
