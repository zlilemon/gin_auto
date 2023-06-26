package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/app/auth"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	log.Infof("start to Ping ... ")
	currentTime := time.Now()
	pingContext := fmt.Sprintf("ping Success @ %s", currentTime)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": pingContext,
	})

	return
}

func TestAuth(c *gin.Context) {
	log.Infof("start to TestAuth ... ")

	currentTime := time.Now()
	pingContext := fmt.Sprintf("ping Success @ %s", currentTime)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": pingContext,
	})

	return
}

func GetUserInfo(c *gin.Context) error {
	log.Info("start go GetUserInfo")

	openid := c.Query("openid")
	//err := auth.SFreshToken(c)
	var userInfoModel []*UserInfo
	var resp comm.Result

	userInfoModel, err := UserService.SGetUserInfo(c, openid)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = userInfoModel[0] //返回第一条用户信息记录，正常情况下也只有一条信息记录
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func SaveUserInfo(c *gin.Context) error {
	log.Info("start go SaveUserInfo")

	var req SaveUserInfoReq
	c.BindJSON(&req)

	var userInfoModel UserInfo
	userInfoModel.OpenId = req.OpenId
	userInfoModel.NickName = req.NickName
	userInfoModel.AvatarUrl = req.AvatarUrl
	userInfoModel.PhoneNo = req.PhoneNo

	log.Infof("get post argv, openid:%s, nick_name:%s", userInfoModel.OpenId, userInfoModel.NickName)

	err := UserService.SSaveUserInfo(c, userInfoModel)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "ok",
	})

	return err
}

func WxLogin(c *gin.Context) error {
	log.Infof("enter WxLogin")

	jsCode := c.Query("code")
	log.Infof("WxLogin, code:%s", jsCode)
	wxSession, err := auth.WxLogin(jsCode)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": wxSession,
	})

	return err
}

func WxSavePhoneNumber(c *gin.Context) error {
	log.Infof("enter WxSavePhoneNumber")

	openid := c.Query("openid")
	//accessToken := c.Query("access_token")
	phoneCode := c.Query("code")

	log.Infof("openid:%s, phoneCode:%s", openid, phoneCode)

	// 后台服务先从微信侧获取accessToken
	accessTokenInfo := auth.AccessTokenInfo{}
	err := auth.AuthorService.SGetAccessToken(c, &accessTokenInfo)

	tryNum := 0
	phoneNumber, err, wxRespErrCode := UserService.SWxSavePhoneNumber(c, openid, accessTokenInfo.AccessToken, phoneCode)
	tryNum += 1

	log.Infof("wxRespErrCode:%d", wxRespErrCode)

	if wxRespErrCode == 42001 && tryNum == 1 {
		// accessToken 时效，且只尝试过一次请求
		// 重新获取accessToken
		err = auth.AuthorService.SSaveAccessToken(c)

		// 重新查询accessToken
		err = auth.AuthorService.SGetAccessToken(c, &accessTokenInfo)

		phoneNumber, err, wxRespErrCode = UserService.SWxSavePhoneNumber(c, openid, accessTokenInfo.AccessToken, phoneCode)
	}

	var resp comm.Result
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = phoneNumber
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func GetClockInfo(c *gin.Context) error {
	log.Infof("enter GetClockInfo ")

	openid := c.Query("openid")
	//err := auth.SFreshToken(c)
	var getClockReq GetClockReq
	var resp comm.Result

	getClockReq.OpenId = openid

	clockInfoModel, err := UserService.SGetClockInfo(c, getClockReq)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = clockInfoModel
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func GetClockSummary(c *gin.Context) error {
	log.Infof("enter GetClockSummary ")

	openid := c.Query("openid")
	//err := auth.SFreshToken(c)
	var getClockReq GetClockReq
	var resp comm.Result

	getClockReq.OpenId = openid

	clockInfoModel, err := UserService.SGetClockInfo(c, getClockReq)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = clockInfoModel
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func GetRankingInfoList(c *gin.Context) error {
	log.Infof("enter GetRankingInfoList ")

	rankingType := c.Query("ranking_type")
	//err := auth.SFreshToken(c)
	var getRankingListReq GetRankingListReq
	var resp comm.Result

	getRankingListReq.RankingType = rankingType

	rankingInfoResp, err := UserService.SGetRankingInfoList(c, getRankingListReq)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = rankingInfoResp
	}

	c.JSON(http.StatusOK, resp)

	return err
}
