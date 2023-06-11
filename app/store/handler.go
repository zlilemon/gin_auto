package store

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"net/http"
	"strconv"
)

func ShopList(c *gin.Context) error {
	log.Info("start go ShopList")

	city := c.Query("city")

	var req GetStoreInfoListReq
	var resp comm.Result

	req.City = city
	log.Infof("city:%s", req.City)

	//var storeInfoModel []*StoreInfo
	err := StoreService.SGetStoreInfoList(c, req, &resp)

	log.Infof("resp_data:%v", resp.Data)

	c.JSON(http.StatusOK, resp)

	/*
		c.JSON(http.StatusOK, gin.H{
			"errcode": 0,
			"errmsg": "ok",
			"data": storeInfoModel,
		})
	*/
	return err
}

func ShopDetail(c *gin.Context) error {
	log.Info("start to ShopDetail")
	shopId := c.Query("shop_id")

	var req StoreDetailReq
	var resp comm.Result

	req.StoreId = shopId
	log.Infof("shop_id:%s", req.StoreId)

	err := StoreService.SGetStoreInfo(c, req, &resp)
	log.Infof("resp_data:%v", resp.Data)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
	}

	c.JSON(http.StatusOK, resp)

	return err

}

// 公示接口，eg：wifi、公告
func ShopNotice(c *gin.Context) error {
	log.Info("start to ShopNotice")
	//shopId := c.Query("shop_id")

	var resp comm.Result

	//req.StoreId = shopId
	//log.Infof("shop_id:%s", req.StoreId)

	err := StoreService.SGetStoreNotice(c, &resp)
	log.Infof("resp_data:%v", resp.Data)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
	}

	c.JSON(http.StatusOK, resp)

	return err

}

// 公告接口, 获取在当前时间以后的有效的公告列表
func BulletinList(c *gin.Context) error {
	log.Info("start to BulletinList")
	//shopId := c.Query("shop_id")

	var resp comm.Result

	//req.StoreId = shopId
	//log.Infof("shop_id:%s", req.StoreId)

	err := StoreService.SGetBulletinList(c, &resp)
	log.Infof("resp_data:%v", resp.Data)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
	}

	c.JSON(http.StatusOK, resp)

	return err
}

// 公告详情
func BulletinDetail(c *gin.Context) error {
	log.Info("start to BulletinDetail")
	bulletinId := c.Query("bulletin_id")

	var req BulletinDetailReq
	var resp comm.Result

	err := errors.New("")
	req.Id, err = strconv.ParseInt(bulletinId, 10, 64)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	}
	log.Infof("bulletin_id:%d", req.Id)

	//req.StoreId = shopId
	//log.Infof("shop_id:%s", req.StoreId)

	err = StoreService.SGetBulletinDetail(c, req, &resp)
	log.Infof("resp_data:%v", resp.Data)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
	}

	c.JSON(http.StatusOK, resp)

	return err
}
