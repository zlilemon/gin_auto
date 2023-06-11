package account

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"net/http"
	"strconv"
)

/*
func AccountCharge(c *gin.Context) error {
	log.Info("start go account charge")

	var req AccountChargeReq
	var resp comm.Result

	c.BindJSON(&req)

	err := AccountService.SCharge(c, req, &resp)
	c.JSON(http.StatusOK, resp)

	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"data": resp,
	//	})
	return err
}

*/

func AccountGetBalance(c *gin.Context) error {
	log.Infof("start to AccountGetBalance ")
	var req BalanceReq

	var balanceResp BalanceResp
	var resp comm.Result

	openId := c.Query("openid")
	req.OpenId = openId

	err := AccountService.SAccountGet(c, req, &balanceResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = balanceResp
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func AccountWaterList(c *gin.Context) error {
	log.Infof("start to AccountWaterList")
	var req WaterListReq
	var waterListResp WaterListResp
	var resp comm.Result

	openId := c.Query("openid")
	tranType := c.Query("tran_type")
	strPage := c.DefaultQuery("page", "0")
	//pageSize := c.GetInt("page_size")
	beginDate := c.DefaultQuery("begin_date", "")
	endDate := c.DefaultQuery("end_date", "")

	strPageSize := c.DefaultQuery("page_size", "20")

	err := errors.New("")
	req.Page, err = strconv.Atoi(strPage)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	}

	req.PageSize, err = strconv.Atoi(strPageSize)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	}

	req.OpenId = openId
	req.TranType = tranType
	//req.PageSize = pageSize  // 默认查询20条记录
	req.BeginDate = beginDate
	req.EndDate = endDate

	err = AccountService.SAccountWaterList(c, req, &waterListResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = waterListResp
	}

	c.JSON(http.StatusOK, resp)
	return err
}

func AccountWaterDetail(c *gin.Context) error {
	log.Infof("start to AccountWaterDetail")
	var req WaterDetailReq
	var resp comm.Result

	outTradeNo := c.Query("out_trade_no")

	req.OutTradeNo = outTradeNo

	err := AccountService.SAccountWaterDetail(c, req, &resp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		//resp.Data = balanceResp
	}

	c.JSON(http.StatusOK, resp)
	return err
}

func AccountCharge(c *gin.Context) error {
	log.Infof("start to AccountCharge ")
	var chargeReq ChargeReq

	var chargeResp ChargeResp
	var resp comm.Result

	c.BindJSON(&chargeReq)

	err := AccountService.SAccountCharge(c, chargeReq, &chargeResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = chargeResp
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func AccountConsume(c *gin.Context) error {
	log.Infof("start to AccountConsume ")
	var consumeReq ConsumeReq

	var consumeResp ConsumeResp
	var resp comm.Result

	c.BindJSON(&consumeReq)

	log.Infof("consumeReq;%+v", consumeReq)
	err := AccountService.SAccountConsume(c, consumeReq, &consumeResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = consumeResp
	}

	c.JSON(http.StatusOK, resp)

	return err
}

// 积分操作
func JFAccountGetBalance(c *gin.Context) error {
	log.Infof("start to JFAccountGetBalance ")
	var req BalanceReq

	var balanceResp BalanceResp
	var resp comm.Result

	openId := c.Query("openid")
	req.OpenId = openId

	err := AccountService.SJFAccountGet(c, req, &balanceResp)
	resp.Data = balanceResp
	c.JSON(http.StatusOK, resp)

	return err
}

func JFAccountCharge(c *gin.Context) error {
	log.Infof("start to JFAccountCharge ")
	var chargeReq ChargeReq

	var chargeResp ChargeResp
	var resp comm.Result

	c.BindJSON(&chargeReq)

	err := AccountService.SJFAccountCharge(c, chargeReq, &chargeResp)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		//resp.Data = balanceResp
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func JFAccountConsume(c *gin.Context) error {
	log.Infof("start to JFAccountConsume ")
	var consumeReq ConsumeReq

	var consumeResp ConsumeResp
	var resp comm.Result

	c.BindJSON(&consumeReq)

	err := AccountService.SJFAccountConsume(c, consumeReq, &consumeResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = consumeResp
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func JFAccountWaterDetail(c *gin.Context) error {
	log.Infof("start to JFAccountWaterDetail")
	var req WaterDetailReq
	var resp comm.Result

	outTradeNo := c.Query("out_trade_no")

	req.OutTradeNo = outTradeNo

	err := AccountService.SJFAccountWaterDetail(c, req, &resp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		//resp.Data = balanceResp
	}

	c.JSON(http.StatusOK, resp)
	return err
}

func JFAccountWaterList(c *gin.Context) error {
	log.Infof("start to JFAccountWaterList")
	var req WaterListReq
	var waterListResp WaterListResp
	var resp comm.Result

	openId := c.Query("openid")
	tranType := c.Query("tran_type")
	strPage := c.DefaultQuery("page", "0")
	//pageSize := c.GetInt("page_size")
	beginDate := c.DefaultQuery("begin_date", "")
	endDate := c.DefaultQuery("end_date", "")

	strPageSize := c.DefaultQuery("page_size", "20")

	err := errors.New("")
	req.Page, err = strconv.Atoi(strPage)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	}

	req.PageSize, err = strconv.Atoi(strPageSize)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	}

	req.OpenId = openId
	req.TranType = tranType
	//req.PageSize = pageSize  // 默认查询20条记录
	req.BeginDate = beginDate
	req.EndDate = endDate

	err = AccountService.SJFAccountWaterList(c, req, &waterListResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = waterListResp
	}

	c.JSON(http.StatusOK, resp)
	return err
}

// 充值映射档位查询
func ChargeGearMappingList(c *gin.Context) error {
	log.Infof("start to ChargeGearMappingList ")

	var chargeGearReq GetChargeGearReq
	var chargeGearResp GetChargeGearResp
	var resp comm.Result

	chargeGearReq.Type = "all"
	err := AccountService.SChargeGearMapping(c, chargeGearReq, &chargeGearResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = chargeGearResp
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func ChargeGearMappingInfo(c *gin.Context) error {
	log.Infof("start to ChargeGearMappingInfo ")

	var chargeGearReq GetChargeGearReq
	var chargeGearResp GetChargeGearResp
	var resp comm.Result

	chargeGearReq.Type = ""
	strChargeAmount := c.Query("charge_amount")
	iChargeAmount, err := strconv.ParseInt(strChargeAmount, 10, 64)
	chargeGearReq.ChargeAmount = iChargeAmount * 100

	err = AccountService.SChargeGearMapping(c, chargeGearReq, &chargeGearResp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = chargeGearResp
	}
	c.JSON(http.StatusOK, resp)

	return err
}
