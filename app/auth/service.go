package auth

import (
	"encoding/json"
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

/*
func ProvideService() (IService, error) {
	var s = &Service{}
	return s, nil
}
*/

var AuthorService = new(Service)

func (s *Service) SSaveAccessToken(c *gin.Context) (err error) {
	log.Infof("enter SGetAccessToken ")

	client := &http.Client{}

	url := fmt.Sprintf("%s/cgi-bin/token?"+
		"appid=%s&secret=%s&grant_type=client_credential",
		config.WxPayOption.WxURI, config.WxPayOption.WxAppID, config.WxPayOption.WxSecret)

	log.Infof("SGetAccessToken, url:%s", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("SGetAccessToken, request failed, err:%s", err)
		return err
	}

	// 处理返回结果
	response, _ := client.Do(request)
	body, err := ioutil.ReadAll(response.Body)

	jsonStr := string(body)
	log.Infof("jsonStr:", jsonStr)
	session := AccessTokenResponse{}
	if err := json.Unmarshal(body, &session); err != nil {
		log.Errorf("json.Unmarshal error: %s", err)

		return err
	}
	if session.ErrorCode != 0 {
		// 返回具体的错误码，需要根据错误码做具体处理
		log.Errorf("errcode:%s, errmsg:%s", session.ErrorCode, session.ErrMsg)
		return
	} else {
		// 获取accessToken成功，写入或更新db中accessToken内容

		existNum, dbErr := AuthorRepository.GetNumOfAccessToken(c)
		if dbErr != nil {
			log.Errorf("GetNumOfAccessToken error, errMsg:%s", dbErr)
			err = dbErr
			return
		}

		if existNum == 0 {
			err = AuthorRepository.InsertAccessToken(c, session.AccessToken, session.ExpiresIn)
		} else {
			err = AuthorRepository.UpdateAccessToken(c, session.AccessToken, session.ExpiresIn)
		}

	}
	return
}

func (s *Service) SGetAccessToken(c *gin.Context, accessTokenInfo *AccessTokenInfo) (err error) {
	log.Infof("SGetAccessToken - ")

	accessToken, expiresIn, err := AuthorRepository.GetAccessToken(c)
	if err != nil {
		log.Errorf("failed to SGetAccessToken - %+v", err)
		return
	}
	accessTokenInfo.AccessToken = accessToken
	accessTokenInfo.ExpiresIn = expiresIn

	log.Infof("SGetAccessToken Success")

	return
}

func (s *Service) SFreshToken(c *gin.Context) (err error) {
	log.Infof("SFreshToken - ")

	num, err := AuthorRepository.GetNumOfAccessToken(c)
	if err != nil {
		log.Errorf("failed to refress token - %+v", err)
		return err
	}

	log.Infof("GetNumOfAccessToken, num : %d", num)
	if num == 0 {
		// 不存在，需要插入动作
		return nil
	} else {
		// 存在 access_token，做更新操作
		err = AuthorRepository.UpdateAccessToken(c, "1243", 4565)
	}

	return nil
}
