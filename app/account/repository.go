package account

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zlilemon/gin_auto/pkg/database"
	"github.com/zlilemon/gin_auto/pkg/log"
	"strings"
	"time"
)

type IRepository interface {
}

type Repository struct {
}

var AccountRepository = new(Repository)

func (r *Repository) GetAccountBalance(c *gin.Context, req BalanceReq, resp *BalanceResp) (err error) {
	log.Infof("GetAccountBalance ")

	var accountBalanceRespList []*BalanceModel
	result := database.AccountDB.Table("account_balance").Where("openid=?", req.OpenId).Find(&accountBalanceRespList)
	if result.Error != nil {
		log.Errorf("GetAccountBalance error, err_msg:%s", result.Error)
		err = result.Error
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetAccountBalance Null, can not find account balance info")
		log.Infof("begin to init account balance to zero, openid:%s", req.OpenId)

		insertSql := "insert into account.account_balance set openid=?, currency='CNY', amount=0, version_id=0, status='Valid'"
		log.Infof("insert sql:%s", insertSql)
		if dbErr := database.AccountDB.Exec(insertSql, req.OpenId); err != nil {
			log.Errorf("account_charge insert to db error, err_smg:%s, error", dbErr.Error)
			return dbErr.Error
		} else {
			log.Infof("init account_balance success, openid:%s", req.OpenId)
		}
		//err = errors.New("GetAccountBalance Null, can not find account balance info")
		resp.OpenId = req.OpenId
		resp.Currency = "CNY"
		resp.Amount = 0
		resp.VersionId = 0
		resp.Status = "Valid"

		log.Infof("GetAccountBalance ok")
		return
	}

	log.Infof("result.RowsAffected num:%d", result.RowsAffected)
	//获取到数据，取第一表记录返回
	resp.OpenId = accountBalanceRespList[0].OpenId
	resp.Currency = accountBalanceRespList[0].Currency
	resp.Amount = accountBalanceRespList[0].Amount

	log.Infof("resp : +%v", accountBalanceRespList[0])
	log.Infof("openid:%s, currency:%s, amount:%d", resp.OpenId, resp.Currency, resp.Amount)

	resp.VersionId = accountBalanceRespList[0].VersionId
	resp.Status = accountBalanceRespList[0].Status

	log.Infof("GetAccountBalance ok")
	return
}

/*
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
*/

func (r *Repository) AccountCharge(c *gin.Context, req ChargeReq, resp *ChargeResp) (err error) {
	log.Infof("GetAccountBalance ")

	// 更新账户余额
	updateSql := "update account.account_balance set amount=amount + ?, version_id = version_id + 1 where openid=? and version_id=?"
	log.Infof("update account_balance sql : %s", updateSql)

	if error := database.AccountDB.Debug().Exec(updateSql, req.Amount, req.OpenId, req.VersionId); error.Error != nil {
		log.Errorf("account_balance update to db error, err_smg:%s, error")
		err = error.Error
		return error.Error
	} else {
		log.Infof("Update account_balance success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	resp.OpenId = req.OpenId
	resp.Currency = req.Currency
	//resp.AfterSaveAmount = req.Amount

	//TODO : 增加写账户充值流水记录，跟账户充值放到一个事务中执行
	//insertSql := "insert into account.account_water set static_date=?, openid=?, out_trade_no=?, channel_order_no=?, " +
	//	"tran_type=?, sub_tran_type=?, currency=?, amount=?, tran_time=?, remark=?"
	insertSql := "insert into account.account_water set static_date=?, openid=?, out_trade_no=?, channel_order_no=?, " +
		"tran_type=?, sub_tran_type=?, currency=?, amount=?, remark=?, tran_time=?, tran_time_unix_time=?"
	log.Infof("insert account_water sql : %s", insertSql)

	//static_date := strings.Index(req.ChargeTime, " ")
	//timeLayout := "2006-01-02 15:04:05"
	//loc, _ := time.LoadLocation("Local")
	//theTime, _ := time.ParseInLocation(timeLayout, req.ChargeTime, loc) //使用模板在对应时区转化为time.time类型
	//sr := theTime.Unix()                                                //转化为时间戳 类型是int64
	//staticTime := req.ChargeTime.Format("2006-01-02 15:04:05")
	//staticDate := strings.Index(staticTime, " ")

	staticTime := time.Unix(req.ChargeTime, 0)
	tt := staticTime.Format("2006-01-02 15:04:05")

	//staticTime := req.ChargeTime.Format("2006-01-02 15:04:05")
	staticDate := strings.Index(tt, " ")

	if error := database.AccountDB.Debug().Exec(insertSql, staticDate, req.OpenId, req.OutTradeNO, req.ChannelOrderNo,
		"save", "", req.Currency, req.Amount, "", tt, req.ChargeTime); error.Error != nil {
		log.Errorf("account_water insert to db error, err_smg:%s", error.Error)
		err = error.Error
		return error.Error
	} else {
		log.Infof("insert account_water success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	return
}

func (r *Repository) AccountConsume(c *gin.Context, req ConsumeReq, resp *ConsumeResp) (err error) {
	log.Infof("AccountConsume ")

	// 更新账户余额
	updateSql := "update account.account_balance set amount=amount - ?, version_id = version_id + 1 where openid=? and version_id=?"
	log.Infof("update account_balance sql : %s", updateSql)

	if error := database.DB.Debug().Exec(updateSql, req.Amount, req.OpenId, req.VersionId); error.Error != nil {
		log.Errorf("account_balance update to db error, err_smg:%s, error")
		err = error.Error
		return error.Error
	} else {
		log.Infof("Update account_balance success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	//TODO : 增加写账户充值流水记录，跟账户充值放到一个事务中执行
	//insertSql := "insert into account.account_water set static_date=?, openid=?, out_trade_no=?,  " +
	//	"tran_type=?, sub_tran_type=?, currency=?, amount=?, tran_time=?, remark=?"
	// 时间戳问题没解决，先插入没有tran_time的版本
	insertSql := "insert into account.account_water set static_date=?, openid=?, out_trade_no=?,  " +
		"tran_type=?, sub_tran_type=?, currency=?, amount=?, remark=?, tran_time=?, tran_time_unix_time=?"
	log.Infof("insert account_water sql : %s", insertSql)

	staticTime := time.Unix(req.ConsumeTime, 0)
	tt := staticTime.Format("2006-01-02 15:04:05")
	//staticTime = staticTime.Format("2006-01-02 15:04:05")
	staticDate := strings.Index(tt, " ")

	log.Infof("req.ConsumeTime:%d, consumeTime:%s, staticDate:%s", req.ConsumeTime, staticTime, staticDate)
	//timestamp := time.Now().Unix()
	//staticDate := strings.Index(req.ConsumeTime, " ")
	//timeLayout := "2006-01-02 15:04:05"
	//loc, _ := time.LoadLocation("Local")
	//theTime, _ := time.ParseInLocation(timeLayout, req.ConsumeTime, loc) //使用模板在对应时区转化为time.time类型
	//sr := theTime.Unix()                                                //转化为时间戳 类型是int64

	//if error := database.AccountDB.Debug().Exec(insertSql, staticDate, req.OpenId, req.OutTradeNO,
	//	"consume", "", req.Currency, req.Amount, req.ConsumeTime, ""); error.Error != nil {
	if error := database.AccountDB.Debug().Exec(insertSql, staticDate, req.OpenId, req.OutTradeNO,
		"consume", "", req.Currency, req.Amount, "", tt, req.ConsumeTime); error.Error != nil {
		log.Errorf("account_water insert to db error, err_smg:%s", error.Error)
		err = error.Error
		return error.Error
	} else {
		log.Infof("insert account_water success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	resp.OpenId = req.OpenId
	resp.Currency = req.Currency
	//resp.AfterSaveAmount = req.Amount
	log.Infof("AccountConsume Success, openid:%s, amount:%d", req.OpenId, req.Amount)

	return
}

func (r *Repository) GetAccountWaterList(c *gin.Context, req WaterListReq) (waterModel []*WaterModel, err error) {
	log.Infof("GetAccountWaterList ")

	//waterModel := make([]*WaterModel, 0)
	// 查询账户交易流水
	// 根据交易时间字段，倒序获取最近20条记录
	//result := database.AccountDB.Table("account_water").Order("tran_time desc").

	// 设置查询分页
	/*
		if req.Page >= 0 && req.PageSize > 0 {
			database.AccountDB = database.AccountDB.Limit(req.PageSize).Offset((req.Page) * req.PageSize)

			log.Infof("enter set select limit, pageSize:%d, page:%d", req.PageSize, req.Page)
		}

	*/

	/*
		result := database.AccountDB.Table("account_water").Order("created_at desc").
			Where("openid=?", req.OpenId).Find(waterModel)

		if result.Error != nil {
			log.Errorf("GetStoreInfoList error, err_msg:%s", result.Error)
		}
		if result.RowsAffected == 0 {
			log.Infof("GetStoreInfoList Null, can not find store info")
		}

	*/

	var result *gorm.DB

	if req.OpenId != "" && req.TranType != "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.AccountDB.Debug().Table("account_water").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Order("created_at desc").
			Where("openid=? and tran_type=? and "+
				"created_at >=? and created_at<?", req.OpenId, req.TranType, req.BeginDate, req.EndDate).
			Find(&waterModel)
	} else if req.OpenId != "" && req.TranType != "" {
		result = database.AccountDB.Debug().Table("account_water").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Order("created_at desc").
			Where("openid=? and tran_type=? ", req.OpenId, req.TranType).
			Find(&waterModel)
	} else if req.OpenId != "" {
		result = database.AccountDB.Debug().Table("account_water").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Order("created_at desc").
			Where("openid=? ", req.OpenId).Find(&waterModel).Limit(req.PageSize).Offset((req.Page) * req.PageSize)
	} else {
		log.Errorf("openid is null")
		return
	}

	if result.Error != nil {
		err = result.Error
		log.Errorf("GetAccountWaterList error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetAccountWaterList Null, can not find openid consume water")
	}

	log.Infof("GetAccountWaterList Success")

	return
}

func (r *Repository) GetAccountWaterDetail(c *gin.Context, req WaterDetailReq, resp *WaterInfo) (err error) {
	log.Infof("GetAccountWaterDetail ")

	waterModelList := make([]*WaterModel, 0)
	// 查询账户交易流水
	result := database.AccountDB.Table("account_water").
		Where("out_trade_no=?", req.OutTradeNo).Find(waterModelList)

	if result.Error != nil {
		err = result.Error
		log.Errorf("GetAccountWaterDetail error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetAccountWaterDetail Null, can not find water info")
	}

	resp.OpenId = waterModelList[0].OpenId
	resp.OutTradeNo = waterModelList[0].OutTradeNo

	return
}

// 积分处理
func (r *Repository) GetJFAccountBalance(c *gin.Context, req BalanceReq, resp *BalanceResp) (err error) {
	log.Infof("GetJFAccountBalance ")

	var accountBalanceRespList []*BalanceModel
	result := database.AccountDB.Debug().Table("jf_account_balance").Where("openid=?", req.OpenId).Find(&accountBalanceRespList)
	if result.Error != nil {
		err = result.Error
		log.Errorf("GetJFAccountBalance error, err_msg:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetJFAccountBalance Null, can not find account balance")
		log.Infof("begin to init jf_account balance to zero, openid:%s", req.OpenId)

		insertSql := "insert into account.jf_account_balance set openid=?, currency='CNY', amount=0, version_id=0, status='Valid'"
		log.Infof("insert sql:%s", insertSql)
		if dbErr := database.AccountDB.Exec(insertSql, req.OpenId); err != nil {
			log.Errorf("jf_account_charge insert to db error, err_smg:%s, error", dbErr.Error)
			return dbErr.Error
		} else {
			log.Infof("init jf_account_balance success, openid:%s", req.OpenId)
		}
		//err = errors.New("GetAccountBalance Null, can not find account balance info")
		resp.OpenId = req.OpenId
		resp.Currency = "CNY"
		resp.Amount = 0
		resp.VersionId = 0
		resp.Status = "Valid"

		log.Infof("GetJFAccountBalance ok")
		return
	}

	//获取到数据，取第一表记录返回
	resp.OpenId = accountBalanceRespList[0].OpenId
	resp.Currency = accountBalanceRespList[0].Currency
	resp.Amount = accountBalanceRespList[0].Amount
	resp.VersionId = accountBalanceRespList[0].VersionId
	resp.Status = accountBalanceRespList[0].Status

	log.Infof("GetJFAccountBalance ok")
	return
}

func (r *Repository) JFAccountCharge(c *gin.Context, req ChargeReq, resp *ChargeResp) (err error) {
	log.Infof("JFAccountCharge ")

	// 更新账户余额
	updateSql := "update account.jf_account_balance set amount=amount + ?, version_id = version_id + 1 " +
		"where openid=? and version_id=?"
	log.Infof("update jf_account_balance sql : %s", updateSql)

	if error := database.AccountDB.Debug().Exec(updateSql, req.Amount, req.OpenId, req.VersionId); error.Error != nil {
		log.Errorf("jf_account_balance update to db error, err_smg:%s, error", error.Error)
		err = error.Error
		return error.Error
	} else {
		log.Infof("Update jf_account_balance success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	resp.OpenId = req.OpenId
	resp.Currency = req.Currency
	//resp.AfterSaveAmount = req.Amount

	//TODO : 增加写账户充值流水记录，跟账户充值放到一个事务中执行
	insertSql := "insert into account.jf_account_water set static_date=?, openid=?, out_trade_no=?, channel_order_no=?, " +
		"tran_type=?, sub_tran_type=?, currency=?, amount=?, tran_time=?, tran_time_unix_time=?, remark=?"
	log.Infof("insert jf_account_water sql : %s", insertSql)

	//static_date := strings.Index(req.ChargeTime, " ")
	//timeLayout := "2006-01-02 15:04:05"
	//loc, _ := time.LoadLocation("Local")
	//theTime, _ := time.ParseInLocation(timeLayout, req.ChargeTime, loc) //使用模板在对应时区转化为time.time类型
	//sr := theTime.Unix()                                                //转化为时间戳 类型是int64
	staticTime := time.Unix(req.ChargeTime, 0)
	tt := staticTime.Format("2006-01-02 15:04:05")

	//staticTime := req.ChargeTime.Format("2006-01-02 15:04:05")
	staticDate := strings.Index(tt, " ")
	//tm := time.Unix(req.ChargeTime, 0)
	if error := database.AccountDB.Debug().Exec(insertSql, staticDate, req.OpenId, req.OutTradeNO, req.ChannelOrderNo,
		"save", "", req.Currency, req.Amount, tt, req.ChargeTime, req.PayInfo); error.Error != nil {
		log.Errorf("jf_account_water insert to db error, err_smg:%s, error", error.Error)
		err = error.Error
		return error.Error
	} else {
		log.Infof("insert jf_account_water success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	return
}

func (r *Repository) JFAccountConsume(c *gin.Context, req ConsumeReq, resp *ConsumeResp) (err error) {
	log.Infof("JFAccountConsume ")

	// 更新账户余额
	updateSql := "update account.jf_account_balance set amount=amount - ?, version_id = version_id + 1 where openid=? and version_id=?"
	log.Infof("update jf_account_balance sql : %s", updateSql)

	if error := database.DB.Exec(updateSql, req.Amount, req.OpenId, req.VersionId); error.Error != nil {
		log.Errorf("jf_account_balance update to db error, err_smg:%s, error", error.Error)
		err = error.Error
		return error.Error
	} else {
		log.Infof("Update account_balance success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	//TODO : 增加写账户充值流水记录，跟账户充值放到一个事务中执行
	//insertSql := "insert into account.jf_account_water set static_date=?, openid=?, out_trade_no=?,  " +
	//	"tran_type=?, sub_tran_type=?, currency=?, amount=?, tran_time=?, remark=?"
	insertSql := "insert into account.jf_account_water set static_date=?, openid=?, out_trade_no=?,  " +
		"tran_type=?, sub_tran_type=?, currency=?, amount=?, remark=?, tran_time=?, tran_time_unix_time=?"
	log.Infof("insert jf_account_water sql : %s", insertSql)

	//staticTime := req.ConsumeTime.Format("2006-01-02 15:04:05")
	//staticDate := strings.Index(staticTime, " ")
	staticTime := time.Unix(req.ConsumeTime, 0)
	tt := staticTime.Format("2006-01-02 15:04:05")
	//staticTime = staticTime.Format("2006-01-02 15:04:05")
	staticDate := strings.Index(tt, " ")

	//staticDate := strings.Index(req.ConsumeTime, " ")
	//timeLayout := "2006-01-02 15:04:05"
	//loc, _ := time.LoadLocation("Local")
	//theTime, _ := time.ParseInLocation(timeLayout, req.ConsumeTime, loc) //使用模板在对应时区转化为time.time类型
	//sr := theTime.Unix()                                                //转化为时间戳 类型是int64

	if error := database.AccountDB.Debug().Exec(insertSql, staticDate, req.OpenId, req.OutTradeNO,
		"consume", "", req.Currency, req.Amount, "", tt, req.ConsumeTime); error.Error != nil {
		log.Errorf("jf_account_water insert to db error, err_smg:%s", error.Error)
		err = error.Error
		return error.Error
	} else {
		log.Infof("insert jf_account_water success, openid:%s, version_id:%d", req.OpenId, req.VersionId)
	}

	resp.OpenId = req.OpenId
	resp.Currency = req.Currency
	resp.OutTradeNo = req.OutTradeNO
	resp.Amount = req.Amount
	//resp.AfterSaveAmount = req.Amount

	log.Infof("JFAccountConsume Success, openid:%s, amount:%d", req.OpenId, req.Amount)

	return
}

func (r *Repository) GetJFAccountWaterDetail(c *gin.Context, req WaterDetailReq, resp *WaterInfo) (err error) {
	log.Infof("GetJFAccountWaterDetail ")

	waterModelList := make([]*WaterModel, 0)
	// 查询账户交易流水
	result := database.AccountDB.Table("jf_account_water").
		Where("out_trade_no=?", req.OutTradeNo).Find(waterModelList)

	if result.Error != nil {
		err = result.Error
		log.Errorf("GetJFAccountWaterDetail error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetJFAccountWaterDetail Null, can not find water info")
	}

	resp.OpenId = waterModelList[0].OpenId
	resp.OutTradeNo = waterModelList[0].OutTradeNo

	return
}

func (r *Repository) GetJFAccountWaterList(c *gin.Context, req WaterListReq) (waterModel []*WaterModel, err error) {
	log.Infof("GetJFAccountWaterList ")

	//waterModel := make([]*WaterModel, 0)
	// 查询账户交易流水
	// 根据交易时间字段，倒序获取最近20条记录
	//result := database.AccountDB.Table("account_water").Order("tran_time desc").

	// 设置查询分页
	/*
		if req.Page >= 0 && req.PageSize > 0 {
			database.AccountDB = database.AccountDB.Limit(req.PageSize).Offset((req.Page) * req.PageSize)

			log.Infof("enter set select limit, pageSize:%d, page:%d", req.PageSize, req.Page)
		}

	*/

	/*
		result := database.AccountDB.Table("account_water").Order("created_at desc").
			Where("openid=?", req.OpenId).Find(waterModel)

		if result.Error != nil {
			log.Errorf("GetStoreInfoList error, err_msg:%s", result.Error)
		}
		if result.RowsAffected == 0 {
			log.Infof("GetStoreInfoList Null, can not find store info")
		}

	*/

	var result *gorm.DB

	if req.OpenId != "" && req.TranType != "" && req.BeginDate != "" && req.EndDate != "" {
		result = database.AccountDB.Debug().Table("jf_account_water").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Order("created_at desc").
			Where("openid=? and tran_type=? and "+
				"created_at >=? and created_at<?", req.OpenId, req.TranType, req.BeginDate, req.EndDate).
			Find(&waterModel)
	} else if req.OpenId != "" && req.TranType != "" {
		result = database.AccountDB.Debug().Table("jf_account_water").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Order("created_at desc").
			Where("openid=? and tran_type=? ", req.OpenId, req.TranType).
			Find(&waterModel)
	} else if req.OpenId != "" {
		result = database.AccountDB.Debug().Table("jf_account_water").Limit(req.PageSize).Offset((req.Page)*req.PageSize).
			Order("created_at desc").
			Where("openid=? ", req.OpenId).Find(&waterModel).Limit(req.PageSize).Offset((req.Page) * req.PageSize)
	} else {
		log.Errorf("openid is null")
		return
	}

	if result.Error != nil {
		err = result.Error
		log.Errorf("GetJFAccountWaterList error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetJFAccountWaterList Null, can not find openid consume water")
	}

	log.Infof("GetJFAccountWaterList Success")

	return
}

// 获取账户充值档位配置
func (r *Repository) GetChargeGear(c *gin.Context, req GetChargeGearReq) (chargeGearMappingList []*ChargeGearInfo, err error) {
	log.Infof("GetChargeGear ")

	//var chargeGearMappingList []*ChargeGearInfo
	result := database.DB
	if req.Type == "all" {
		result = database.AccountDB.Debug().Table("charge_point_mapping").Find(&chargeGearMappingList)
	} else {
		result = database.AccountDB.Debug().Table("charge_point_mapping").
			Where("charge_amount=?", req.ChargeAmount).Find(&chargeGearMappingList)
	}

	if result.Error != nil {
		log.Errorf("GetChargeGear error, err_msg:%s", result.Error)
		err = result.Error
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("GetChargeGear Null, can not find charge gear mapping, charge_amount:%d", req.ChargeAmount*100)
		err = errors.New("GetChargeGear Null, can not find charge gear mapping")
		return
	}

	log.Infof("GetChargeGear Success")
	return
}
