package user

import (
	"gin_auto/pkg/database"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IRepository interface {
}

type Repository struct {
}

var UserRepository = new(Repository)

func (r *Repository) GetUserInfo(c *gin.Context, openid string) (userInfoModel []*UserInfo, err error) {
	log.Infof("GetUserInfo ")

	//var userInfoModel []UserInfo

	//num = r.dbProxy.Table("access_token").Find(&accessTokenModel).RowsAffected
	result := database.DB.Debug().Table("user_info").Where("openid=?", openid).Find(&userInfoModel)
	if result.Error != nil {
		log.Errorf("GetUserInfo error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetUserInfo Null, can not find openid : %s", openid)
	}

	//log.Infof("GetUserInfo finish")

	return
}

func (r *Repository) SaveUserInfo(c *gin.Context, userInfoModel UserInfo) (err error) {
	log.Infof("enter SaveUserInfo")

	// 保存用户信息
	insertSql := "insert into wxapp.user_info set openid=?, nick_name=?, avatar_url=?, phone_no=?"
	log.Infof("SaveUserInfo sql : %s", insertSql)

	if result := database.DB.Exec(insertSql, userInfoModel.OpenId, userInfoModel.NickName, userInfoModel.AvatarUrl,
		userInfoModel.PhoneNo); result.Error != nil {
		log.Errorf("SaveUserInfo error, to db error, err_msg:%s", result.Error)
		return result.Error
	} else {
		log.Infof("SaveUserInfo success, openid:%s, nick_name:%s, avatar_url:%s, phone_no:%s",
			userInfoModel.OpenId, userInfoModel.NickName, userInfoModel.AvatarUrl, userInfoModel.PhoneNo)
	}

	return nil
}

func (r *Repository) SaveWxPhoneNumber(c *gin.Context, openid string, phoneNumber string) (err error) {
	log.Infof("enter SaveWxPhoneNumber ")
	// 先判断用户openid 是否存在
	var userInfoModel []*UserInfo
	result := database.DB.Table("user_info").Where("openid=?", openid).Find(&userInfoModel)
	if result.Error != nil {
		log.Errorf("GetUserInfo error, err_msg:%s", result.Error)
	}

	if result.RowsAffected == 0 {
		// 该openid 不存在用户信息表中，做插入动作
		log.Infof("GetUserInfo Null, can not find openid : %s", openid)

		insertSql := "insert wxapp.user_info set openid=?, phone_no=?"
		log.Infof("Insert user and phone_no info sql :%s", insertSql)
		if result = database.DB.Debug().Exec(insertSql, openid, phoneNumber); result.Error != nil {
			log.Errorf("insert user and save phone_no error, to db error, err_msg:%s", result.Error)
			return result.Error
		} else {
			log.Infof("insert user and save phone_no success, openid:%s, phoneNumber:%s",
				openid, phoneNumber)
		}
	} else {
		// 该openid 存在用户信息表中，做更新动作
		// 保存用户电话号码
		updateSql := "update wxapp.user_info set phone_no=? where openid=? limit 1"
		log.Infof("SaveWxPhoneNumber sql : %s", updateSql)

		if result := database.DB.Debug().Exec(updateSql, phoneNumber, openid); result.Error != nil {
			log.Errorf("SaveWxPhoneNumber error, to db error, err_msg:%s", result.Error)
			return result.Error
		} else {
			log.Infof("SaveWxPhoneNumber success, openid:%s, phoneNumber:%s",
				openid, phoneNumber)
		}
	}
	return
}

func (r *Repository) AddClock(c *gin.Context, addClockInfoReq AddClockReq) (clockInfoModel []*ClockInfo, err error) {
	log.Infof("enter AddClock ")

	// 先查询该用户是否有打卡记录
	result := database.DB.Table("clock_info").Where("openid=? and out_trade_no=?",
		addClockInfoReq.OpenId, addClockInfoReq.OutTradeNo).Find(&clockInfoModel)
	if result.Error != nil {
		log.Errorf("get clock record error, err_msg:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("get clock info Null, can not find clock record, openid : %s", addClockInfoReq.OpenId)
	}

	if result.RowsAffected == 0 {
		// 该订单号不存在打卡记录，插入打卡记录
		// 若该订单号存在打卡记录，就不再重复插入
		insertSql := "insert into wxapp.clock_info set openid=?, out_trade_no=?, static_date=?, duration_time=?, status=?"
		log.Infof("add clock info sql : %s", insertSql)

		if result := database.DB.Debug().Exec(insertSql, addClockInfoReq.OpenId, addClockInfoReq.OutTradeNo, addClockInfoReq.StaticDate,
			addClockInfoReq.DurationTime, "valid"); result.Error != nil {
			log.Errorf("insert clock info error, to db error, err_msg:%s", result.Error)
			return
		} else {
			log.Infof("addClock success, openid:%s, out_trade_no:%s, static_date:%s, duration_time:%s",
				addClockInfoReq.OpenId, addClockInfoReq.OutTradeNo, addClockInfoReq.StaticDate, addClockInfoReq.DurationTime)
		}
	}

	return
}

func (r *Repository) GetClock(c *gin.Context, req GetClockReq) (clockInfoModel []*ClockInfo, err error) {
	//result := database.DB.Debug().Table("clock_info").Where("openid=?", req.OpenId).Find(&clockInfoModel)
	result := database.DB.Debug().Table("clock_info").
		Select("openid, static_date, sum(duration_time) as duration_time").Where("openid=?", req.OpenId).
		Group("openid, static_date").Find(&clockInfoModel)
	if result.Error != nil {
		log.Errorf("GetClock error, err_msg:%s", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Infof("GetUserInfo Null, can not find openid : %s", req.OpenId)
	}

	log.Infof("GetClock success")

	return
}

func (r *Repository) AddRanking(c *gin.Context, req AddRankingReq) (rankingInfoModel []*RankingInfo, err error) {
	log.Infof("enter AddRanking ")

	// 先查询该用户是否有排行记录
	result := database.DB.Table("ranking_info").Where("openid=? ",
		req.OpenId).Find(&rankingInfoModel)
	if result.Error != nil {
		log.Errorf("get ranking record error, err_msg:%s", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Infof("get ranking info Null, can not find ranking record, openid : %s", req.OpenId)
	}

	if result.RowsAffected == 0 {
		// 该用户不存在排行记录，插入排行记录
		insertSql := "insert into wxapp.ranking_info set openid=?, nick_name=?, avatar_url=?, " +
			"weekly_times=1, weekly_duration_time=?, " +
			"monthly_times=1, monthly_duration_time=?, totally_times=1, totally_duration_time=?"
		log.Infof("add ranking info sql : %s", insertSql)

		if result := database.DB.Debug().Exec(insertSql, req.OpenId, req.NickName, req.AvatarUrl,
			req.DurationTime, req.DurationTime, req.DurationTime); result.Error != nil {
			log.Errorf("insert ranking info error, to db error, err_msg:%s", result.Error)
			return
		} else {
			log.Infof("add ranking info success, openid:%s, duration_time:%s",
				req.OpenId, req.DurationTime)
		}
	} else {
		// 该用户存在排行记录，更新排行信息
		updateSql := "update wxapp.ranking_info set weekly_times=weekly_times + 1, weekly_duration_time = weekly_duration_time +?," +
			" monthly_times = monthly_times + 1, monthly_duration_time = monthly_duration_time + ?, " +
			"totally_times = totally_times +1, totally_duration_time = totally_duration_time + ? " +
			"where openid=? limit 1"
		log.Infof("update ranking info sql : %s", updateSql)

		if result := database.DB.Debug().Exec(updateSql, req.DurationTime, req.DurationTime, req.DurationTime, req.OpenId); result.Error != nil {
			log.Errorf("SaveWxPhoneNumber error, to db error, err_msg:%s", result.Error)
			return nil, result.Error
		} else {
			log.Infof("update ranking info success, openid:%s, duration_time:%s",
				req.OpenId, req.DurationTime)
		}
	}

	return
}

func (r *Repository) GetRankingList(c *gin.Context, req GetRankingListReq) (rankingInfoModel []*RankingInfo, err error) {
	//result := database.DB.Debug().Table("clock_info").Where("openid=?", req.OpenId).Find(&clockInfoModel)
	var result *gorm.DB

	if req.RankingType == "weekly" {
		result = database.DB.Debug().Table("ranking_info").
			Select("openid, nick_name, avatar_url, weekly_times, weekly_duration_time").Order("weekly_duration_time desc").Limit(20).Find(&rankingInfoModel)
		if result.Error != nil {
			log.Errorf("GetRankingList by weekly error, err_msg:%s", result.Error)
			err = result.Error
			return
		}
		if result.RowsAffected == 0 {
			log.Infof("GetRankingList by weekly Null, can not find ranking info")
		}
	} else if req.RankingType == "monthly" {
		result = database.DB.Debug().Table("ranking_info").
			Select("openid, nick_name, avatar_url, monthly_times, monthly_duration_time").Order("month_duration_time desc").Limit(20).Find(&rankingInfoModel)
		if result.Error != nil {
			log.Errorf("GetRankingList by monthly error, err_msg:%s", result.Error)
			err = result.Error
			return
		}
		if result.RowsAffected == 0 {
			log.Infof("GetRankingList by monthly Null, can not find ranking info")
		}
	} else {
		result = database.DB.Debug().Table("ranking_info").
			Select("openid, nick_name, avatar_url, totally_times, totally_duration_time").Order("totally_duration_time desc").Limit(20).Find(&rankingInfoModel)
		if result.Error != nil {
			log.Errorf("GetRankingList by totally error, err_msg:%s", result.Error)
			err = result.Error
			return
		}
		if result.RowsAffected == 0 {
			log.Infof("GetRankingList by totally Null, can not find ranking info")
		}
	}

	log.Infof("GetRankingList Success")

	return
}
