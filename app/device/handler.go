package device

import (
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"net/http"
)

/*
func GetDeviceInfo(c *gin.Context) error {
	log.Info("start go GetDeviceInfo")

	var deviceInfoModel []*DeviceInfo
	deviceInfoModel, err := DeviceService.SGetDeviceInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": deviceInfoModel,
	})

	return err
}

*/

func DeviceOperation(c *gin.Context) error {
	log.Infof("start to DeviceOperation - ")
	var req OperationReq
	var deviceControlResp OperationResp
	var resp comm.Result

	c.BindJSON(&req)

	err := DeviceService.SDeviceOperation(c, req, &deviceControlResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		//resp.Data = phoneNumber
	}

	c.JSON(http.StatusOK, resp)

	return err

}

func XiaoMiDeviceOperation(c *gin.Context) error {
	log.Infof("start to XiaoMiDeviceOperation - ")
	var req OperationReq
	var deviceControlResp OperationResp
	var resp comm.Result

	c.BindJSON(&req)

	err := DeviceService.SXiaomiDeviceOperation(c, req, &deviceControlResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		//resp.Data = phoneNumber
	}

	c.JSON(http.StatusOK, resp)

	return err

}
