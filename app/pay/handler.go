package pay

import (
	"gin_auto/pkg/comm"
	"gin_auto/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WxPay(c *gin.Context) error {
	log.Info("start go wxPay")

	var req WxPayReq
	var wxpayResp WxPayResp
	var resp comm.Result

	c.BindJSON(&req)

	err := PayService.SWxPayV3(c, req, &wxpayResp)

	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = "ok"
		resp.Data = wxpayResp
	}

	c.JSON(http.StatusOK, resp)
	/*
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": resp,
		})*/
	return err
}

func WxQueryOrder(c *gin.Context) error {
	log.Infof("start to WxQueryOrder")

	var req QueryReq
	var resp comm.Result

	c.BindJSON(&req)
	//transaction_id := c.Query("transaction_id")
	//mchid := c.Query("mchid")

	err := PayService.SWxQueryOrder(c, req, &resp)

	c.JSON(http.StatusOK, resp)
	return err
}

func WxPayNotify(c *gin.Context) error {
	log.Infof("start to WxPayNotify")
	var req WxpayNotifyReq
	var resp comm.Result

	//log.Infof("wxpayNotify: %+v", req)

	//body,_ := ioutil.ReadAll(c.Request.Body)

	//log.Infof("request body:%s", body)
	//fmt.Println(json.Valid(body))

	//if err := json.Unmarshal(body, &req); err != nil {
	//	log.Errorf("notify request conver to json error, errMsg:%s", err)
	//}
	//c.BindJSON(&req)

	//log.Infof("-- wxpayNotify: %+v", req)

	err := PayService.SWxPayNotify(c, req, &resp)
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

/*
func downloadPlatformCA(c *gin.Context) error {
	log.Info("start go downloadPlatformCA")

	//var req PayReq
	var resp comm.Result

	//c.BindJSON(&req)

	orderId, err := PayService.SDownloadPlatformCA(c, &resp)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": orderId,
	})

	return err
}

*/
