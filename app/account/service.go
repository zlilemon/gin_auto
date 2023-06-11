package account

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
)

type IService interface {
	SCharge(c *gin.Context, req ChargeReq, resp *comm.Result) (err error)
}

type Service struct {
	repo IRepository
}

var AccountService = new(Service)

func (s *Service) SAccountGet(c *gin.Context, req BalanceReq, resp *BalanceResp) (err error) {
	log.Infof("enter SAccountGet ")

	//根据条件，查询账户内容
	err = AccountRepository.GetAccountBalance(c, req, resp)
	if err != nil {
		log.Errorf("GetAccountBalance error, err_msg:%s", err)
		return
	}

	// 判断当前账户是否有效
	openid := resp.OpenId
	status := resp.Status

	log.Infof("openid:%s, status:%s", openid, status)

	if status != "Valid" {
		log.Errorf("openid:%s, account Invalid, account save error", openid)
		err = errors.New("account Invalid, account save error")
		return
	}

	log.Infof("enter SAccountGet Success")

	return
}

func (s *Service) SAccountCharge(c *gin.Context, req ChargeReq, resp *ChargeResp) (err error) {

	log.Infof("enter SAccountCharge")
	// 判断充值金额，是否在配置的充值档位中
	chargeGearReq := GetChargeGearReq{}
	chargeGearResp := GetChargeGearResp{}
	chargeGearReq.ChargeAmount = req.Amount
	err = AccountService.SChargeGearMapping(c, chargeGearReq, &chargeGearResp)
	if err != nil {
		log.Errorf("SChargeGearMapping error, err_msg:%s", err)
		return
	}

	// 从账户中，先把该用户的当前版本号获取出来，防止并发带来的重复充值
	accountGetReq := BalanceReq{}
	accountGetReq.OpenId = req.OpenId

	accountGetResp := BalanceResp{}

	err = AccountService.SAccountGet(c, accountGetReq, &accountGetResp)

	if err != nil {
		log.Errorf("SCharge error, err:%v", err)
		return
	}

	// 进行充值
	req.VersionId = accountGetResp.VersionId

	// 根据充值档位，替换充值到账户中到最终金额
	log.Infof("charge_amount:%d, account_point:%d", req.Amount, chargeGearResp.ChargeGearList[0].AccountPoint)

	req.Amount = chargeGearResp.ChargeGearList[0].AccountPoint

	err = AccountRepository.AccountCharge(c, req, resp)
	if err != nil {
		log.Errorf("SCharge error, err:%v", err)
		return
	}
	return
}

func (s *Service) SAccountConsume(c *gin.Context, req ConsumeReq, resp *ConsumeResp) (err error) {
	log.Infof("enter SAccountConsume")

	// 从账户中，先把该用户的当前版本号获取出来，防止并发带来的重复扣费
	accountGetReq := BalanceReq{}
	accountGetReq.OpenId = req.OpenId

	accountGetResp := BalanceResp{}

	err = AccountService.SAccountGet(c, accountGetReq, &accountGetResp)

	if err != nil {
		log.Errorf("SAccountGet error, err:%v", err)
		return
	}

	// 消耗动作
	req.VersionId = accountGetResp.VersionId
	err = AccountRepository.AccountConsume(c, req, resp)
	if err != nil {
		log.Errorf("AccountConsume error, err:%v", err)
		return
	}

	log.Infof("SAccountConsume Success")
	return
}

func (s *Service) SAccountWaterList(c *gin.Context, req WaterListReq, resp *WaterListResp) (err error) {
	log.Infof("enter SAccountWaterList ")

	// 根据条件，查询账户下的充值记录
	accountWaterListResp := make([]*WaterInfo, 0)

	waterModel, err := AccountRepository.GetAccountWaterList(c, req)

	if err != nil {
		log.Errorf("SAccountWaterList error, err:%v", err)
		return
	}

	for _, v := range waterModel {
		oneWaterInfo := WaterInfo{}

		oneWaterInfo.OutTradeNo = v.OutTradeNo
		oneWaterInfo.OpenId = v.OpenId
		oneWaterInfo.Amount = v.Amount
		oneWaterInfo.Currency = v.Currency
		oneWaterInfo.TranType = v.TranType
		oneWaterInfo.ChannelOrderNo = v.ChannelOrderNo
		//oneWaterInfo.PayMethod =
		oneWaterInfo.TranTime = v.TranTime

		accountWaterListResp = append(accountWaterListResp, &oneWaterInfo)

	}

	resp.WaterList = accountWaterListResp

	return
}

func (s *Service) SAccountWaterDetail(c *gin.Context, req WaterDetailReq, resp *comm.Result) (err error) {
	log.Infof("enter SAccountWaterDetail ")

	// 根据条件，查询具体交易订单记录
	accountWaterDetailResp := WaterInfo{}

	err = AccountRepository.GetAccountWaterDetail(c, req, &accountWaterDetailResp)

	if err != nil {
		log.Errorf("SAccountWaterList error, err:%v", err)
		return
	}

	resp.Data = accountWaterDetailResp

	return
}

// 积分操作
func (s *Service) SJFAccountGet(c *gin.Context, req BalanceReq, resp *BalanceResp) (err error) {
	log.Infof("enter SJFAccountGet ")

	//根据条件，查询账户内容
	err = AccountRepository.GetJFAccountBalance(c, req, resp)
	if err != nil {
		log.Errorf("GetJFAccountBalance error, err_msg:%s", err)
		return
	}

	// 判断当前账户是否有效
	openid := resp.OpenId
	status := resp.Status

	if status != "Valid" {
		log.Errorf("openid:%s, account Invalid, account get error", openid)
		err = errors.New("account Invalid, account get error")
		return
	}

	log.Infof("enter SJFAccountGet Success")

	return
}

func (s *Service) SJFAccountCharge(c *gin.Context, req ChargeReq, resp *ChargeResp) (err error) {
	log.Infof("enter SJFAccountCharge")

	// 从账户中，先把该用户的当前版本号获取出来，防止并发带来的重复充值
	accountGetReq := BalanceReq{}
	accountGetReq.OpenId = req.OpenId

	accountGetResp := BalanceResp{}

	err = AccountService.SJFAccountGet(c, accountGetReq, &accountGetResp)

	if err != nil {
		log.Errorf("SJFAccountCharge error, err:%v", err)
		return
	}

	// 进行充值
	req.VersionId = accountGetResp.VersionId
	err = AccountRepository.JFAccountCharge(c, req, resp)
	if err != nil {
		log.Errorf("SCharge error, err:%v", err)
		return
	}
	return
}

func (s *Service) SJFAccountConsume(c *gin.Context, req ConsumeReq, resp *ConsumeResp) (err error) {
	log.Infof("enter SJFAccountConsume")

	// 从账户中，先把该用户的当前版本号获取出来，防止并发带来的重复扣费
	accountGetReq := BalanceReq{}
	accountGetReq.OpenId = req.OpenId

	accountGetResp := BalanceResp{}

	err = AccountService.SJFAccountGet(c, accountGetReq, &accountGetResp)

	if err != nil {
		log.Errorf("SJFAccountGet error, err:%v", err)
		return
	}

	// 消耗动作
	req.VersionId = accountGetResp.VersionId
	err = AccountRepository.JFAccountConsume(c, req, resp)
	if err != nil {
		log.Errorf("SJFAccountConsume error, err:%v", err)
		return
	}

	log.Infof("SJFAccountConsume Success")
	return
}

func (s *Service) SJFAccountWaterDetail(c *gin.Context, req WaterDetailReq, resp *comm.Result) (err error) {
	log.Infof("enter SJFAccountWaterDetail ")

	// 根据条件，查询具体交易订单记录
	accountWaterDetailResp := WaterInfo{}

	err = AccountRepository.GetJFAccountWaterDetail(c, req, &accountWaterDetailResp)

	if err != nil {
		log.Errorf("SJFAccountWaterDetail error, err:%v", err)
		return
	}

	resp.Data = accountWaterDetailResp

	return
}

func (s *Service) SJFAccountWaterList(c *gin.Context, req WaterListReq, resp *WaterListResp) (err error) {
	log.Infof("enter SJFAccountWaterList ")

	// 根据条件，查询账户下的充值记录
	accountWaterListResp := make([]*WaterInfo, 0)

	waterModel, err := AccountRepository.GetJFAccountWaterList(c, req)

	if err != nil {
		log.Errorf("SJFAccountWaterList error, err:%v", err)
		return
	}

	for _, v := range waterModel {
		oneWaterInfo := WaterInfo{}

		oneWaterInfo.OutTradeNo = v.OutTradeNo
		oneWaterInfo.OpenId = v.OpenId
		oneWaterInfo.Amount = v.Amount
		oneWaterInfo.Currency = v.Currency
		oneWaterInfo.TranType = v.TranType
		oneWaterInfo.ChannelOrderNo = v.ChannelOrderNo
		//oneWaterInfo.PayMethod =
		oneWaterInfo.TranTime = v.TranTime

		accountWaterListResp = append(accountWaterListResp, &oneWaterInfo)

	}

	resp.WaterList = accountWaterListResp

	return
}

// 充值档位映射表
func (s *Service) SChargeGearMapping(c *gin.Context, req GetChargeGearReq, resp *GetChargeGearResp) (err error) {

	log.Infof("enter SChargeGearMapping")

	chargeGearModel := make([]*ChargeGearInfo, 0)
	chargeGearModel, err = AccountRepository.GetChargeGear(c, req)
	if err != nil {
		log.Errorf("SChargeGearMapping error, err:%v", err)
		return
	}

	chargeGearList := make([]ChargeGearItem, 0)
	for _, v := range chargeGearModel {
		newItem := ChargeGearItem{}
		newItem.RelationId = v.RelationId
		newItem.RelationName = v.RelationName
		newItem.ChargeAmount = v.ChargeAmount
		newItem.AccountPoint = v.AccountPoint

		chargeGearList = append(chargeGearList, newItem)
	}

	resp.ChargeGearList = chargeGearList

	log.Infof("SChargeGearMapping Success ")
	return
}
