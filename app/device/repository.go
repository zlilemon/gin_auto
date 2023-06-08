package device

import (
	"gin_auto/pkg/database"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IRepository interface {
}

type Repository struct {
	dbProxy *gorm.DB
}

var DeviceRepository = new(Repository)

/*
func (r *Repository) GetDeviceInfo(c *gin.Context, deviceInfoModel []*DeviceInfo,) (err error) {
	log.Infof("GetDeviceInfo ")

	//var deviceInfoModel []DeviceInfo

	//num = r.dbProxy.Table("access_token").Find(&accessTokenModel).RowsAffected
	database.DB.Table("device_info").Find(&deviceInfoModel)

	log.Infof("GetDeviceInfo done")

	return nil
}
*/

func (r *Repository) GetDeviceMapping(c *gin.Context, req OperationReq) (deviceMapping []*DeviceMapping, err error) {
	log.Infof("GetDeviceMapping ")

	var result *gorm.DB

	result = database.StoreDB.Table("device_mapping").Debug().Where("store_id=? and seat_id=? "+
		"and device_function_type=?",
		req.StoreId, req.SeatId, req.DeviceFunctionType).Find(&deviceMapping)

	if result.Error != nil {
		log.Errorf("GetDeviceMapping error, msg:%s", result.Error)

		err = result.Error
		return
	}

	log.Infof("GetDeviceMapping Success - ")

	return
}
