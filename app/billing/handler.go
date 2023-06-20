package billing

import (
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"net/http"
)

func GetOrderStatusCheck(c *gin.Context) error {
	log.Info("start to Order")

	var req BillingStatusCheckReq
	var resp comm.Result
	//var billingStatusCheckResp []*BillingStatusCheckResp

	c.BindJSON(&req)

	billingStatusCheckResp, err := BillingService.SGetOrderStatusCheck(c, req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = billingStatusCheckResp
	}

	c.JSON(http.StatusOK, resp)
	/*
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": resp,
		})*/
	return err
}
