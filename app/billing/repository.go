package billing

import (
	"gin_auto/pkg/database"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type IRepository interface {
}

type Repository struct {
}

var BillingRepository = new(Repository)

func (r *Repository) InsertOrder(c *gin.Context, orderDetail OrderDetail) (err error) {
	log.Infof("Insert Orderid")

	result := database.StoreDB.Debug().Table("order_detail").Create(&orderDetail)

	if result.Error != nil {
		log.Errorf("InsertOrder error, err_msg:%s", result.Error)
		return
	}

	log.Infof("Insert Orderid Success, orderDetail :%v", orderDetail)
	return
}

func (r *Repository) GetOrderInfo(c *gin.Context, req BillingInfoReq) (billingRespList []*BillingInfoResp, err error) {
	log.Infof("GetOrder ")
	var result *gorm.DB

	log.Infof("Openid:%s, order_type:%s, out_trade_no:%s, begin_date:%s, end_date:%s", req.OpenId, req.OrderType, req.OutTradeNo, req.BeginDate, req.EndDate)
	if req.OpenId != "" && req.OutTradeNo != "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Where("openid=? and out_trade_no=? and book_begin_date=? and book_end_date=?",
				req.OpenId, req.OutTradeNo, req.BeginDate, req.EndDate).Order("created_at desc").
			Find(&billingRespList)
	} else if req.OpenId != "" && req.OrderType != "" && req.OutTradeNo == "" && req.BeginDate == "" && req.EndDate == "" {
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Where("openid=? and order_type=?", req.OpenId, req.OrderType).Order("created_at desc").
			Find(&billingRespList)
	} else if req.OpenId != "" && req.OrderType == "" && req.OutTradeNo == "" && req.BeginDate == "" && req.EndDate == "" {
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Where("openid=?", req.OpenId).Order("created_at desc").
			Find(&billingRespList)
	} else if req.OpenId != "" && req.OutTradeNo == "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Where("openid=? and  book_begin_date=? and book_end_date=?", req.OpenId, req.BeginDate, req.EndDate).
			Order("created_at desc").
			Find(&billingRespList)
	} else if req.OpenId == "" && req.OutTradeNo != "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Where("out_trade_no=? and book_begin_date=? and book_end_date=?", req.OutTradeNo, req.BeginDate, req.EndDate).
			Order("created_at desc").
			Find(&billingRespList)
	} else if req.OpenId != "" && req.OutTradeNo != "" && req.BeginDate == "" && req.EndDate == "" {
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Where("openid=? and out_trade_no=?", req.OpenId, req.OutTradeNo).Order("created_at desc").
			Find(&billingRespList)
	} else {
		log.Errorf("can not support select condition, openid:%s, out_trade_no:%s, begin_date:%s, end_date:%s",
			req.OpenId, req.OutTradeNo, req.BeginDate, req.EndDate)
	}

	if result.Error != nil {
		log.Errorf("GetOrderInfo error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetOrderInfo Null, can not find billing order info")
	}

	log.Infof("result.RowsAffected:%d", result.RowsAffected)
	log.Infof("billingRespList len:%d", len(billingRespList))
	log.Infof("billingRespList: %s", billingRespList)

	return
}

func (r *Repository) GetOrderInfoDivideByTime(c *gin.Context, req BillingInfoDivideByTimeReq) (billingRespList []*BillingInfoResp, err error) {
	log.Infof("GetOrderInfoDivideByTime ")

	var result *gorm.DB
	timeUnix := time.Now().Unix()

	// 设置查询分页
	/*
		if req.Page >= 0 && req.PageSize > 0 {
			log.Infof("enter set db limit - req.PageSize:%d, req.Page:%d", req.PageSize, req.Page)

			database.StoreDB = database.StoreDB.Limit(req.PageSize).Offset((req.Page) * req.PageSize)
		}

	*/

	if req.FutureFlag == "0" {
		// 历史预定记录：订单结束时间 小于 当前时间
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).Where("openid=? and end_unix_time < ?",
			req.OpenId, timeUnix).Order("end_unix_time desc").Find(&billingRespList)

	} else {
		result = database.StoreDB.Debug().Table("order_detail").Limit(req.PageSize).Offset((req.Page)*req.PageSize).Where("openid=? and end_unix_time >= ?",
			req.OpenId, timeUnix).Order("end_unix_time desc").Find(&billingRespList)
		//
	}
	if result.Error != nil {
		log.Errorf("GetOrderInfoDivideByTime error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetOrderInfoDivideByTime Null, can not find billing order info")
	}

	log.Infof("result.RowsAffected:%d", result.RowsAffected)
	log.Infof("billingRespList len:%d", len(billingRespList))
	log.Infof("billingRespList: %s", billingRespList)

	return
}

func (r *Repository) UpdateOrderInfo(c *gin.Context, req UpdateBillingReq) (err error) {
	log.Infof("UpdateOrderInfo - ")
	var result *gorm.DB
	//result = database.StoreDB.Debug().Table("order_detail").Model(&OrderDetail{}).Where("out_trade_no = ?", req.OutTradeNo).
	//	Updates(OrderDetail{ChannelOrderNo: req.ChannelOrderNo, NotifyStatus: req.NotifyStatus}).Limit(1)

	result = database.StoreDB.Debug().Table("order_detail").Model(&OrderDetail{}).Updates(map[string]interface{}{"ChannelOrderNo": req.ChannelOrderNo, "NotifyStatus": req.NotifyStatus, "PayStatus": req.PayStatus}).Where("out_trade_no = ?", req.OutTradeNo)
	//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	if result.Error != nil {
		log.Errorf("UpdateOrder error, out_trade_no:%s, channel_order_no:%s, notify_status=%s",
			req.OutTradeNo, req.ChannelOrderNo, req.NotifyStatus)
		err = result.Error
		return
	}
	log.Infof("result:%+v", result)

	if result.RowsAffected == 0 {
		log.Errorf("UpdateOrder error, effect row is 0, out_trade_no:%s, channel_order_no:%s, notify_status=%s",
			req.OutTradeNo, req.ChannelOrderNo, req.NotifyStatus)
		err = result.Error
		return
	}

	log.Infof("UpdateOrder Success, out_trade_no:%s, channel_order_no:%s, notify_status=%s",
		req.OutTradeNo, req.ChannelOrderNo, req.NotifyStatus)

	return
}
