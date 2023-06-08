package store

import (
	"gin_auto/pkg/database"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type IRepository interface {
}

type Repository struct {
}

var StoreRepository = new(Repository)

func (r *Repository) GetStoreInfoList(c *gin.Context, city string) (storeInfoModel []*StoreInfo, err error) {
	log.Infof("GetStoreInfoList ")

	result := database.StoreDB.Table("store_info").Where("city=?", city).Find(&storeInfoModel)
	if result.Error != nil {
		log.Errorf("GetStoreInfoList error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetStoreInfoList Null, can not find store info")
	}
	return
}

func (r *Repository) GetStoreInfo(c *gin.Context, storeID string) (storeInfoModel []*StoreInfo, err error) {
	log.Infof("GetStoreInfo ")

	log.Infof("store_id:%s", storeID)

	result := database.StoreDB.Table("store_info").Where("store_id=?", storeID).Find(&storeInfoModel)
	if result.Error != nil {
		log.Errorf("GetStoreInfo error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetStoreInfo Null, can not find store info, store_id:%s", storeID)
	}
	return
}

func (r *Repository) GetSeatInfoByStore(c *gin.Context, storeID string) (seatInfoModel []*TbSeatInfo, err error) {
	log.Infof("GetSeatInfoByStore ")

	result := database.StoreDB.Table("seat_info").Where("store_id=?", storeID).Find(&seatInfoModel)
	if result.Error != nil {
		log.Errorf("GetSeatInfoByStore error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetSeatInfoByStore Null, can not find seat info, store_id:%s", storeID)
	}
	return
}

func (r *Repository) GetBookOrderStatusBySeatId(c *gin.Context, storeID string, seatID string) (
	seatOrderStatusModel []*TbSeatOrderStatus, err error) {
	log.Infof("GetBookOrderStatusBySeatId ")

	result := database.StoreDB.Debug().Table("seat_order_status").Order("book_begin_time").Where(
		"store_id=? and seat_id=? and status in ('SUCCESS', 'LOCK')", storeID, seatID).Find(&seatOrderStatusModel)
	if result.Error != nil {
		log.Errorf("GetBookOrderStatusBySeatId error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetBookOrderStatusBySeatId Null, can not find seat order status, store_id:%s, seat_id:%s",
			storeID, seatID)
	}
	return
}

func (r *Repository) GetStorePrice(c *gin.Context, storeID string) (priceInfoModel []*PriceModel, err error) {
	log.Infof("GetStorePrice ")

	result := database.StoreDB.Table("price").Where("store_id=?", storeID).Find(&priceInfoModel)
	if result.Error != nil {
		log.Errorf("GetStorePrice error, err:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetStorePrice Null, can not find price info, store_id:%s", storeID)
	}

	return
}

func (r *Repository) GetStoreNotice(c *gin.Context) (storeNotice []*StoreNoticeModel, err error) {
	log.Infof("GetStoreNotice ")

	result := database.StoreDB.Debug().Table("store_notice").Find(&storeNotice)
	if result.Error != nil {
		log.Errorf("GetStoreNotice error, err:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetStoreNotice Null, can not find store notice")
	}

	return
}

func (r *Repository) GetStoreQuestion(c *gin.Context) (questionModel []*QuestionModel, err error) {
	log.Infof("GetStoreQuestion ")

	result := database.StoreDB.Debug().Table("store_question").Find(&questionModel)
	if result.Error != nil {
		log.Errorf("GetStoreQuestion error, err:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetStoreQuestion Null, can not find store notice")
	}

	return
}

func (r *Repository) GetBulletinList(c *gin.Context) (bulletinModel []*BulletinSummary, err error) {
	log.Infof("GetBulletinList ")

	// 获取当前时间戳
	currentTime := time.Now()
	currentUnixTime := strconv.FormatInt(currentTime.Unix(), 10)

	result := database.StoreDB.Debug().Table("bulletin").Where("publish_time_unix_time>=? and status='VALID'", currentUnixTime).Find(&bulletinModel)
	if result.Error != nil {
		log.Errorf("GetStoreQuestion error, err:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetStoreQuestion Null, can not find store notice")
	}

	return
}

func (r *Repository) GetBulletinDetail(c *gin.Context, req BulletinDetailReq) (bulletinModel []*Bulletin, err error) {
	log.Infof("GetBulletinDetail ")

	result := database.StoreDB.Debug().Table("bulletin").Where("id=? and status='VALID'", req.Id).Find(&bulletinModel)
	if result.Error != nil {
		log.Errorf("GetBulletinDetail error, err:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetBulletinDetail Null, can not find bulletin detail")
	}

	log.Infof("GetBulletinDetail Success, bulletin_id:%d", req.Id)

	return
}
