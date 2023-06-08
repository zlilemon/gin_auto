package device

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gin_auto/app/billing"
	"gin_auto/app/user"
	"gin_auto/pkg/config"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type IService interface {
	SGetDeviceInfo(c *gin.Context) (err error)
}

type Service struct {
	repo IRepository
}

var DeviceService = new(Service)

/*
func (s *Service) SGetDeviceInfo(c *gin.Context) (deviceInfoModel []*DeviceInfo, err error) {
	log.Infof("SGetDeviceInfo - ")

	err = DeviceRepository.GetDeviceInfo(c, deviceInfoModel)
	if err != nil {
		log.Errorf("failed to refress token - %+v", err)
		return
	}
	return
}

*/
func (s *Service) SDeviceOperation(c *gin.Context, req OperationReq, resp *OperationResp) (err error) {
	log.Infof("SDeviceOperation - ")

	// 校验订单是否存在，并在时间范围内
	billReq := billing.BillingInfoReq{}
	billReq.OpenId = req.OpenId
	billReq.OutTradeNo = req.OutTradeNo

	billResp, err := billing.BillingService.SGetBillingInfo(c, billReq)
	if err != nil {
		log.Errorf("failed to SGetBillingInfo, err:%s", err.Error())
		return
	}
	billItem := billResp[0]

	if billItem.ShopId != req.StoreId || billItem.SeatId != req.SeatId {
		errMsg := fmt.Sprintf("bill info not match, openid:%s, out_trade_no:%s, "+
			"bill_store_id:%s, bill_seat_id:%s, "+
			"req_store_id:%s, req_seat_id:%s",
			req.OpenId, req.OutTradeNo,
			billItem.ShopId, billItem.SeatId,
			req.StoreId, req.SeatId)

		err = errors.New(errMsg)
		log.Errorf("%s", errMsg)

		return
	}

	// 订单中的时间，跟当前时间的校验
	beginUnixTime := billItem.BeginUnixTime
	endUnixTime := billItem.EndUnixTime

	currencyTimeUnix := time.Now().Unix()
	currencyDateTime := time.Unix(currencyTimeUnix, 0).Format("2006-01-02 15:04:05")

	beginDateTime := time.Unix(beginUnixTime, 0).Format("2006-01-02 15:04:05")
	endDateTime := time.Unix(endUnixTime, 0).Format("2006-01-02 15:04:05")

	if currencyTimeUnix <= beginUnixTime || currencyTimeUnix >= endUnixTime {
		log.Errorf("billTime expire, out_trade_no:%s, openid:%s, currency_unix_time:%d, currency_time:%s, "+
			"bill_begin_unix_time:%d, bill_end_unix_time:%d, bill_begin_time:%s, bill_end_time:%s",
			req.OutTradeNo, req.OpenId, currencyTimeUnix, currencyDateTime,
			beginUnixTime, endUnixTime, beginDateTime, endDateTime)

		err = errors.New("billTime expire")
		return
	}
	// 获取对应的硬件设备id
	var deviceMappingModel []*DeviceMapping
	deviceMappingModel, err = DeviceRepository.GetDeviceMapping(c, req)
	if err != nil {
		log.Errorf("GetDeviceMapping err, msg:%s", err.Error())
		return
	}
	deviceMappingItem := deviceMappingModel[0]
	deviceId := deviceMappingItem.DeviceId
	deviceCategory := deviceMappingItem.DeviceCategory
	deviceBrand := deviceMappingItem.DeviceBrand

	req.DeviceId = deviceId
	req.DeviceCategory = deviceCategory
	req.DeviceBrand = deviceBrand

	if deviceBrand == "xiaomi" {
		// 调用米家接口
		log.Infof("begin to s.SXiaomiDeviceOperation")
		err = s.SXiaomiDeviceOperation(c, req, resp)
	} else if deviceBrand == "tuya" {
		// 调用涂鸦接口
	}

	if err != nil {
		log.Errorf("deviceOperation err, msg:%s", err.Error())
		return
	}

	// 记录打卡记录
	addClockReq := user.AddClockReq{}
	addClockReq.OpenId = req.OpenId
	addClockReq.OutTradeNo = req.OutTradeNo
	addClockReq.StaticDate = billItem.BookBeginDate
	addClockReq.DurationTime = billItem.BookDuration
	addClockReq.Status = "valid"

	err = user.UserService.SAddClock(c, addClockReq)
	// 记录失败，也只是打印日志，并不阻塞开门的过程

	log.Infof("deviceOperation ok")

	return
}
func (s *Service) STuyaDeviceOperation(c *gin.Context, req OperationReq, resp *OperationResp) (err error) {
	log.Infof("SDeviceOperation - ")

	deviceId := req.DeviceId

	// 调用设备接口，操作设备
	client := &http.Client{}

	tuyaUrl := fmt.Sprintf("%s/%s/commands", config.TuyaDeviceOption.TuyaURI, deviceId)

	log.Infof("tuyaUrl:%s", tuyaUrl)

	tuyaReq := TuyaReq{}

	tuyaCommandsList := make([]TuyaCommands, 0)
	tuyaCommands := TuyaCommands{}

	tuyaCommands.Code = "switch"
	if req.Cmd == "open" {
		tuyaCommands.Value = true
	} else {
		tuyaCommands.Value = false
	}

	tuyaCommandsList = append(tuyaCommandsList, tuyaCommands)
	tuyaReq.Commands = tuyaCommandsList

	postBody, err := json.Marshal(tuyaReq)
	if err != nil {
		log.Errorf("json.Marshal error, msg:%s", err)
		return
	}

	request, err := http.NewRequest("POST", tuyaUrl, bytes.NewBuffer(postBody))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Errorf("SDeviceOperation, request failed, err:%s", err)
		return
	}

	// 处理返回结果
	response, _ := client.Do(request)
	respBody, err := ioutil.ReadAll(response.Body)
	jsonStr := string(respBody)

	log.Infof("call tuya resp, jsonStr:%s", jsonStr)

	return
}

func (s *Service) SXiaomiDeviceOperation(c *gin.Context, req OperationReq, resp *OperationResp) (err error) {
	// 调用设备接口，操作设备
	client := &http.Client{}

	var operation string
	if req.Cmd == "open" {
		operation = "turn_on"
	} else {
		operation = "turn_off"
	}
	xiaomiUrl := fmt.Sprintf("%s/services/%s/%s", config.XiaomiDeviceOption.XiaomiURI, req.DeviceCategory, operation)

	log.Infof("xiaomiUrl:%s", xiaomiUrl)

	xiaomiReq := XiaoMiReq{}

	xiaomiReq.EntityId = req.DeviceId

	postBody, err := json.Marshal(xiaomiReq)
	if err != nil {
		log.Errorf("json.Marshal error, msg:%s", err)
		return
	}
	log.Infof("postBody:%s", postBody)

	request, err := http.NewRequest("POST", xiaomiUrl, bytes.NewBuffer(postBody))

	token := fmt.Sprintf("Bearer %s", config.XiaomiDeviceOption.XiaomiHomeAssistanceToken)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", token)

	// 处理返回结果
	response, _ := client.Do(request)
	respBody, err := ioutil.ReadAll(response.Body)
	jsonStr := string(respBody)

	log.Infof("call xiaomi resp, jsonStr:%s", jsonStr)
	return
}
