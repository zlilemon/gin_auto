package auth

import (
	"encoding/json"
	"fmt"
	"gin_auto/pkg/comm"
	"gin_auto/pkg/config"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func Users(c *gin.Context) {
	log.Info("enter Users ")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "books",
	})
}

func GetAccessToken(c *gin.Context) error {

	var resp comm.Result
	var accessTokenInfo AccessTokenInfo
	err := AuthorService.SGetAccessToken(c, &accessTokenInfo)

	currentTime := time.Now()
	log.Infof("currentTime:%v", currentTime)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = accessTokenInfo
	}

	c.JSON(http.StatusOK, resp)

	return err
}

func FreshToken(c *gin.Context) error {
	log.Info("start go FreshToken")

	//err := auth.SFreshToken(c)
	err := AuthorService.SFreshToken(c)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "books",
	})

	return err
}

func SaveAccessToken(c *gin.Context) error {
	var resp comm.Result
	//var accessTokenInfo AccessTokenInfo
	err := AuthorService.SSaveAccessToken(c)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = ""
	}

	c.JSON(http.StatusOK, resp)

	return err
}
func WxLogin(jsCode string) (session Code2SessionResponse, err error) {
	client := &http.Client{}

	url := fmt.Sprintf("%s/sns/jscode2session?"+
		"appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.WxPayOption.WxURI, config.WxPayOption.WxAppID, config.WxPayOption.WxSecret, jsCode)

	log.Infof("WxLogin, url:%s", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("WxLogin, request failed, err:%s", err)
		panic(err)
	}

	// 处理返回结果
	response, _ := client.Do(request)
	body, err := ioutil.ReadAll(response.Body)

	jsonStr := string(body)
	log.Infof("jsonStr:", jsonStr)
	if err := json.Unmarshal(body, &session); err != nil {
		log.Errorf("json.Unmarshal error: %s", err)
		session.SessionKey = jsonStr
		return session, err
	}
	if session.ErrorCode != 0 {
		// 微信登录返回具体的错误码，需要根据错误码做具体处理
		log.Errorf("errcode:%s, errmsg:%s", session.ErrorCode, session.ErrMsg)
		return
	} else {
		// 微信登录返回码正确，开始处理其他业务流程
		return
	}

	return
}
