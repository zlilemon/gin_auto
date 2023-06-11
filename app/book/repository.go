package book

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zlilemon/gin_auto/pkg/database"
	"github.com/zlilemon/gin_auto/pkg/log"
)

type IRepository interface {
}

type Repository struct {
}

var BookRepository = new(Repository)

func (r *Repository) GetSeatOrderStatus(c *gin.Context, req SeatOrderStatusReq) (seatOrderStatus []*BookSeatOrderStatus, err error) {
	log.Infof("GetSeatOrderStatus ")
	var result *gorm.DB

	log.Infof("Openid:%s, store_id:%s, store_name:%s, seat_id:%s, seat_name:%s, book_begin_time:%s, book_end_time:%s",
		req.OpenId, req.StoreId, req.StoreName, req.SeatId, req.SeatName, req.BookBeginTime, req.BookEndTime)

	result = database.StoreDB.Debug().Table("seat_order_status").
		Where("store_id=? and seat_id=? and book_begin_time=? and book_end_time=?",
			req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime).Find(&seatOrderStatus)

	if result.Error != nil {
		log.Errorf("GetSeatOrderStatus error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetSeatOrderStatus Null, can not find seat order status")
	}

	log.Infof("result.RowsAffected:%d", result.RowsAffected)

	return
}

func (r *Repository) InsertSeatOrderStatus(c *gin.Context, req SeatOrderStatusReq) (err error) {
	log.Infof("InsertSeatOrderStatus ")
	var result *gorm.DB

	log.Infof("Openid:%s, store_id:%s, store_name:%s, seat_id:%s, seat_name:%s, book_begin_time:%s, book_end_time:%s, to_status:%s",
		req.OpenId, req.StoreId, req.StoreName, req.SeatId, req.SeatName, req.BookBeginTime, req.BookEndTime, req.ToStatus)

	seatOrderStatus := BookSeatOrderStatus{}
	seatOrderStatus.OpenId = req.OpenId
	seatOrderStatus.StoreId = req.StoreId
	seatOrderStatus.StoreName = req.StoreName
	seatOrderStatus.SeatId = req.SeatId
	seatOrderStatus.SeatName = req.SeatName
	seatOrderStatus.BookBeginTime = req.BookBeginTime
	seatOrderStatus.BookEndTime = req.BookEndTime
	seatOrderStatus.Status = req.ToStatus

	result = database.StoreDB.Debug().Table("seat_order_status").Create(&seatOrderStatus)

	if result.Error != nil {
		log.Errorf("InsertSeatOrderStatus error, err_msg:%s", result.Error)
		return
	}

	log.Infof("Insert seat_order_status Success, orderDetail :%v", seatOrderStatus)
	return

	result = database.StoreDB.Debug().Table("seat_order_status").
		Where("store_id=? and seat_id=? and book_begin_time=? and book_end_time=?",
			req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime).Find(&seatOrderStatus)

	if result.Error != nil {
		log.Errorf("GetSeatOrderStatus error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetSeatOrderStatus Null, can not find seat order status")
	}

	log.Infof("result.RowsAffected:%d", result.RowsAffected)

	return
}

func (r *Repository) UpdateSeatOrderStatus(c *gin.Context, req SeatOrderStatusReq) (err error) {
	log.Infof("UpdateSeatOrderStatus ")
	//var result *gorm.DB

	log.Infof("Openid:%s, store_id:%s, store_name:%s, seat_id:%s, seat_name:%s, book_begin_time:%s, book_end_time:%s, to_status:%s",
		req.OpenId, req.StoreId, req.StoreName, req.SeatId, req.SeatName, req.BookBeginTime, req.BookEndTime, req.ToStatus)

	seatOrderStatus := BookSeatOrderStatus{}
	seatOrderStatus.OpenId = req.OpenId
	seatOrderStatus.StoreId = req.StoreId
	seatOrderStatus.StoreName = req.StoreName
	seatOrderStatus.SeatId = req.SeatId
	seatOrderStatus.SeatName = req.SeatName
	seatOrderStatus.BookBeginTime = req.BookBeginTime
	seatOrderStatus.BookEndTime = req.BookEndTime
	seatOrderStatus.Status = req.ToStatus

	// 更新账户余额
	updateSql := "update store.seat_order_status set status=? where store_id=? and seat_id=? and book_begin_time=? and book_end_time=? limit 1"
	log.Infof("update account_balance sql : %s", updateSql)

	if error := database.StoreDB.Debug().Exec(updateSql, req.ToStatus, req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime); error.Error != nil {
		log.Errorf("seat_order_status update to db error, err_smg:%s, error")
		err = error.Error
		return error.Error
	} else {
		log.Infof("Update seat_order_status success, openid:%s, store_id:%s, seat_id=%s, book_begin_time:%s, book_end_time:%s",
			req.OpenId, req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime)
	}
	/*
		result = database.StoreDB.Debug().Table("seat_order_status").Model(&BookSeatOrderStatus{}).
			Updates(map[string]interface{}{"Status": req.ToStatus}).
			Where("store_id = ? and seat_id=? and book_begin_time=? and book_end_time=?",
				req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime)
		//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
		if result.Error != nil {
			log.Errorf("UpdateSeatOrderStatus error, store_id:%s, seat_id:%s, book_begin_time=%s, book_end_time=%s",
				req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime)
			err = result.Error
			return
		}

		log.Infof("result:%+v", result)

		if result.RowsAffected == 0 {
			log.Errorf("UpdateSeatOrderStatus error, effect row is 0, store_id:%s, seat_id:%s, book_begin_time=%s, book_end_time=%s",
				req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime)
			err = result.Error
			return
		}
	*/
	log.Infof("UpdateSeatOrderStatus Success, store_id:%s, seat_id:%s, book_begin_time=%s, book_end_time=%s, to_status:%s",
		req.StoreId, req.SeatId, req.BookBeginTime, req.BookEndTime, req.ToStatus)

	return
}

/*
func (r *Repository) InsertOrder(c *gin.Context, orderDetail OrderDetail) (err error) {
	log.Infof("InsertOrderid ")

	result := database.StoreDB.Table("order_detail").Create(&orderDetail)

	if result.Error != nil {
		log.Errorf("InsertOrder error, err_msg:%s", result.Error)
	}

		//if result.RowsAffected == 0 {
		//	log.Infof("InsertOrder Null, can not find store info")
		//}

	return result.Error
}
*/

/*
func (r *Repository) GetOrder(c *gin.Context, req BookOrderReq) (bookOrderRespList []*BookOrderResp, err error) {
	log.Infof("GetOrder ")
	var result *gorm.DB

	if req.OpenId != "" && req.OutTradeNo != "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.StoreDB.Table("order_detail").Where("openid=? and out_trade_no=? and "+
			"book_begin_date=? and book_end_date=?", req.OpenId, req.OutTradeNo, req.BeginDate, req.EndDate).
			Find(&bookOrderRespList)
	} else if req.OpenId != "" && req.OutTradeNo == "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.StoreDB.Table("order_detail").Where("openid=? and  "+
			"book_begin_date=? and book_end_date=?", req.OpenId, req.BeginDate, req.EndDate).
			Find(&bookOrderRespList)
	} else if req.OpenId == "" && req.OutTradeNo != "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.StoreDB.Table("order_detail").Where("out_trade_no=? and  "+
			"book_begin_date=? and book_end_date=?", req.OutTradeNo, req.BeginDate, req.EndDate).
			Find(&bookOrderRespList)
	} else {
		log.Errorf("can not support select condition, openid:%s, out_trade_no:%s, begin_date:%s, end_date:%s",
			req.OpenId, req.OutTradeNo, req.BeginDate, req.EndDate)
	}

	if result.Error != nil {
		log.Errorf("GetStoreInfoList error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetStoreInfoList Null, can not find store info")
	}

	return
}

*/

/*
func (r *Repository) UpdateOrder(c *gin.Context, req BookOrderUpdateReq) (err error) {
	log.Infof("UpdateOrder ")
	var result *gorm.DB
	result = database.StoreDB.Model(&OrderDetail{}).Where("out_trade_no = ?", req.OutTradeNo).
		Updates(OrderDetail{ChannelOrderNo: req.ChannelOrderNo, NotifyStatus: req.NotifyStatus}).Limit(1)

	if result.Error != nil {
		log.Errorf("UpdateOrder error, out_trade_no:%s, channel_order_no:%s, notify_status=%s",
			req.OutTradeNo, req.ChannelOrderNo, req.NotifyStatus)
		err = result.Error
		return
	}
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

*/
