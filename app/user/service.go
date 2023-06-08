package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gin_auto/pkg/config"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type IService interface {
	SFreshToken(c *gin.Context) (err error)
}

type Service struct {
	repo IRepository
}

var UserService = new(Service)

func (s *Service) HelloWorld() (hello string) {
	return "hello"
}

func (s *Service) SGetUserInfo(c *gin.Context, openid string) (userInfoModel []*UserInfo, err error) {
	log.Infof("SGetUserInfo - ")

	//var userInfoModel []*UserInfo
	userInfoModel, err = UserRepository.GetUserInfo(c, openid)
	if err != nil {
		log.Errorf("failed to SGetUserInfo - %+v", err)
		return
	}

	log.Infof("userInfoMode len:%d, nickName:%s, avaUrl:%s", len(userInfoModel), userInfoModel[0].NickName, userInfoModel[0].AvatarUrl)

	log.Infof("SGetUserInfo done, openid:%s", openid)

	return
}

func (s *Service) SSaveUserInfo(c *gin.Context, userInfoModel UserInfo) (err error) {
	log.Infof("SSaveUserInfo - ")

	err = UserRepository.SaveUserInfo(c, userInfoModel)
	if err != nil {
		log.Errorf("failed to SaveUserInfo - %+v", err)
		return
	}

	log.Infof("SSaveUserInfo done")

	return
}

func (s *Service) SWxSavePhoneNumber(c *gin.Context, openid string, accessToken string, phoneCode string) (phoneNumber string, err error, wxRespErrCode int) {
	log.Infof("SWxSavePhoneNumber - ")

	client := &http.Client{}

	urlPhone := fmt.Sprintf("%s/wxa/business/getuserphonenumber?access_token=%s",
		config.WxPayOption.WxURI, accessToken)

	//jsonTmp := fmt.Sprintf("{\"code\":%s}", phoneCode)
	//var jsonData = []byte(jsonTmp)
	//var jsonData = []byte(`{
	//	"code": phoneCode
	//}"`)

	//jsonData := "{\"code\": + phoneCode }"
	//post := "{\"code\":\"" + phoneCode + "\",\"code\":\"" + phoneCode + "\"}"
	log.Infof("phoneCode:%s", phoneCode)
	postData := "{\"code\":\"" + phoneCode + "\"}"
	var postStr = []byte(postData)

	log.Infof("SWxSavePhoneNumber, url:%s", urlPhone)
	log.Infof("postData:%s, jsonStr:%s", postData, postStr)

	request, err := http.NewRequest("POST", urlPhone, bytes.NewBuffer(postStr))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Errorf("SWxSavePhoneNumber, request failed, err:%s", err)
		return
	}

	// 处理返回结果
	response, _ := client.Do(request)
	respBody, err := ioutil.ReadAll(response.Body)

	wxPhoneResp := WxPhoneResp{}
	jsonStr := string(respBody)
	log.Infof("jsonStr:", jsonStr)
	if err := json.Unmarshal(respBody, &wxPhoneResp); err != nil {
		log.Errorf("json.Unmarshal error: %s", err)
		phoneNumber = ""
		return phoneNumber, err, 0
	}
	if wxPhoneResp.ErrorCode != 0 {
		// 微信返回具体的错误码，需要根据错误码做具体处理
		if wxPhoneResp.ErrorCode == 42001 {
			wxRespErrCode = wxPhoneResp.ErrorCode
		}
		log.Errorf("errcode:%s, errmsg:%s", wxPhoneResp.ErrorCode, wxPhoneResp.ErrMsg)
		err = errors.New(wxPhoneResp.ErrMsg)

		return
	}

	// 微信返回码正确，获取具体电话号码
	phoneNumber = wxPhoneResp.PhoneInfo.PurePhoneNumber
	log.Infof("getPhoneNumber Success, code:%s, phoneNumber:%s", phoneCode, phoneNumber)

	// 保存到用户信息中
	err = UserRepository.SaveWxPhoneNumber(c, openid, phoneNumber)
	if err != nil {
		log.Errorf("SWxSavePhoneNumber error, errMsg:%v", err)
		return phoneNumber, err, 0
	}
	return

}

// 门打开后，增加用户打卡次数
func (s *Service) SAddClock(c *gin.Context, addClockReq AddClockReq) (err error) {
	log.Infof("SAddClock enter")

	_, err = UserRepository.AddClock(c, addClockReq)
	if err != nil {
		log.Errorf("failed to SAddClock - %+v", err)
		return
	}

	log.Infof("SAddClock Success")
	return
}

func (s *Service) SGetClockInfo(c *gin.Context, getClockReq GetClockReq) (clockInfoModel []*ClockInfo, err error) {
	log.Infof("SGetClockInfo enter")

	clockInfoModel, err = UserRepository.GetClock(c, getClockReq)
	if err != nil {
		log.Errorf("failed to SGetClock - %+v", err)
		return
	}

	log.Infof("SGetClockInfo Success")
	return
}

// 预定支付后，增加用户的预定排行榜
func (s *Service) SAddRanking(c *gin.Context, addRankingReq AddRankingReq) (err error) {
	log.Infof("SAddRanking enter")

	_, err = UserRepository.AddRanking(c, addRankingReq)
	if err != nil {
		log.Errorf("failed to SAddRanking - %+v", err)
		return
	}

	log.Infof("SAddRanking Success")
	return
}

// 查询排行信息
func (s *Service) SGetRankingInfoList(c *gin.Context, getRankingListReq GetRankingListReq) (getRankingListRespList []*GetRankingListRespItem, err error) {
	log.Infof("SGetRankingInfoList enter")

	var rankingInfoModel []*RankingInfo
	rankingInfoModel, err = UserRepository.GetRankingList(c, getRankingListReq)
	if err != nil {
		log.Errorf("failed to SGetRankingInfoList - %+v", err)
		return
	}

	//getRankingListRespList := make([]*GetRankingListRespItem, 0)
	// 转换前端输出的字段内容
	for _, v := range rankingInfoModel {
		oneItem := GetRankingListRespItem{}

		oneItem.OpenId = v.OpenId
		if getRankingListReq.RankingType == "weekly" {
			oneItem.Times = v.WeeklyTimes
			oneItem.DurationTime = v.WeeklyDurationTime
		} else if getRankingListReq.RankingType == "monthly" {
			oneItem.Times = v.MonthlyTimes
			oneItem.DurationTime = v.MonthlyDurationTime
		} else {
			oneItem.Times = v.TotallyTimes
			oneItem.DurationTime = v.TotallyDurationTime
		}

		getRankingListRespList = append(getRankingListRespList, &oneItem)
	}

	log.Infof("SGetRankingInfoList Success")
	return
}
