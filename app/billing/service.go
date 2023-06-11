package billing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/log"
	"time"
)

type IService interface {
	// SCreateBilling 生成订单号
	SCreateBilling() (err error)

	// SGetBillingInfo 查询订单号具体信息
	SGetBillingInfo() (err error)
}

type Service struct {
	repo IRepository
}

var BillingService = new(Service)

func (s *Service) SCreateBilling(c *gin.Context, req InsertBillingReq, resp *InsertBillingResp) (err error) {
	log.Infof("SCreateBilling - ")
	timeTemplate1 := "2006-01-02 15:04:05"

	orderDetail := OrderDetail{}
	orderDetail.OutTradeNo = req.OutTradeNo
	orderDetail.ChannelOrderNo = req.ChannelOrderNo
	orderDetail.OrderType = req.OrderType
	orderDetail.PayMethod = req.PayMethod
	orderDetail.OpenId = req.OpenId
	orderDetail.StoreId = req.StoreId
	orderDetail.StoreName = req.StoreName
	orderDetail.SeatId = req.SeatId
	orderDetail.SeatName = req.SeatName
	orderDetail.SeatType = req.SeatType
	orderDetail.BookBeginDate = req.BookBeginDate
	orderDetail.BookEndDate = req.BookEndDate
	orderDetail.BookBeginTime = req.BookBeginTime
	orderDetail.BookEndTime = req.BookEndTime
	orderDetail.PayInfo = req.PayInfo

	if req.OrderType == "bookConsume" {
		// 直接预定座位类型，订单中需要带上该笔订单的预定开始时间和预定结束时间
		tmpBeginTime := fmt.Sprintf("%s %s", req.BookBeginDate, req.BookBeginTime)
		tmpEndTime := fmt.Sprintf("%s %s", req.BookEndDate, req.BookEndTime)
		beginStamp, _ := time.ParseInLocation(timeTemplate1, tmpBeginTime, time.Local)
		endStamp, _ := time.ParseInLocation(timeTemplate1, tmpEndTime, time.Local)
		orderDetail.BeginUnixTime = beginStamp.Unix()
		orderDetail.EndUnixTime = endStamp.Unix()
	}

	orderDetail.BookDuration = req.BookDuration
	orderDetail.Currency = req.Currency
	orderDetail.Amount = req.Amount
	orderDetail.TranTime = req.TranTime
	orderDetail.BillStatus = req.BillStatus

	err = BillingRepository.InsertOrder(c, orderDetail)
	if err != nil {
		log.Errorf("failed to InsertOrder - %+v", err)
		return
	}
	resp.OutTradeNo = req.OutTradeNo

	log.Infof("SCreateBilling Success, out_trade_no:%s", resp.OutTradeNo)
	return
}

func (s *Service) SGetBillingInfo(c *gin.Context, req BillingInfoReq) (resp []*BillingInfoResp, err error) {
	log.Infof("SGetBillingInfo - ")

	//orderResp := make([]*BillingInfoResp, 0)
	resp, err = BillingRepository.GetOrderInfo(c, req)

	if err != nil {
		log.Errorf("failed to GetOrderInfo - %+v", err)
		return
	}

	log.Infof("resp : %+v", resp)
	log.Infof("orderResp len:%d", len(resp))
	//resp = &orderResp
	return
}

func (s *Service) SUpdateBilling(c *gin.Context, req UpdateBillingReq) (err error) {
	log.Infof("SUpdateBilling - ")

	err = BillingRepository.UpdateOrderInfo(c, req)
	if err != nil {
		log.Errorf("failed to update billing - %+v", err)
		return
	}

	log.Infof("SUpdateBilling Success")
	return
}
