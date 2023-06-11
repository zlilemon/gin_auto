package book

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"net/http"
	"strconv"
)

func Order(c *gin.Context) error {
	log.Info("start to Order")

	var req BookReq
	var resp comm.Result

	c.BindJSON(&req)

	err := BookService.SOrder(c, req, &resp)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
	}

	c.JSON(http.StatusOK, resp)
	/*
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": resp,
		})*/
	return err
}

func GetOrder(c *gin.Context) error {
	log.Infof("start to GetOrder")
	var req BookOrderReq
	var resp comm.Result

	req.OpenId = c.Query("openid")
	req.OrderType = c.Query("order_type")
	req.OutTradeNo = c.Query("out_trade_no")
	req.BeginDate = c.Query("begin_date")
	req.EndDate = c.Query("end_date")

	strPage := c.DefaultQuery("page", "0")
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

	log.Infof("order_type:%s", req.OrderType)

	err = BookService.SGetOrder(c, req, &resp)
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

func GetOrderDivideByTime(c *gin.Context) error {
	log.Infof("start to GetOrderDivideByTime")
	var req BookOrderDivideByTimeReq
	var resp comm.Result

	req.OpenId = c.Query("openid")
	req.FutureFlag = c.Query("future_flag")
	//req.Page = c.Query("page")
	strPage := c.DefaultQuery("page", "0")
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

	err = BookService.SGetOrderDivideByTime(c, req, &resp)
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
