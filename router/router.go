package router

import (
	"gin_auto/app/account"
	"gin_auto/app/auth"
	"gin_auto/app/book"
	"gin_auto/app/device"
	"gin_auto/app/pay"
	"gin_auto/app/store"
	"gin_auto/app/user"
	"gin_auto/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		user.Ping(context)
	})

	// user 相关
	router.GET("/user/login", func(context *gin.Context) {
		user.WxLogin(context)
	})
	router.GET("/user/phone", func(context *gin.Context) {
		user.WxSavePhoneNumber(context)
	})
	router.GET("/users", func(context *gin.Context) {
		auth.Users(context)
	})
	router.GET("/user/get", func(context *gin.Context) {
		user.GetUserInfo(context)
	})
	router.POST("/user/saveUserInfo", func(context *gin.Context) {
		user.SaveUserInfo(context)
	})
	router.GET("/user/clockInfo", func(context *gin.Context) {
		user.GetClockInfo(context)
	})
	router.GET("/user/rankingInfoList", func(context *gin.Context) {
		user.GetRankingInfoList(context)
	})

	// 店铺 shop 相关
	router.GET("/shop/list", middleware.AuthCheck(true), func(context *gin.Context) {
		store.ShopList(context)
	})
	router.GET("/shop/detail", func(context *gin.Context) {
		store.ShopDetail(context)
	})
	router.GET("/shop/shopNotice", func(context *gin.Context) {
		store.ShopNotice(context)
	})
	router.GET("/shop/bulletinList", func(context *gin.Context) {
		store.BulletinList(context)
	})
	router.GET("/shop/bulletinDetail", func(context *gin.Context) {
		store.BulletinDetail(context)
	})

	// 支付相关
	router.POST("/pay/order", func(context *gin.Context) {
		book.Order(context)
	})
	router.POST("/pay/notify", func(context *gin.Context) {
		pay.WxPayNotify(context)
	})

	// 账户相关
	// 账户充值 -- 有service层实现，还没封装handler
	//router.POST("/account/charge", func(context *gin.Context) {
	//
	//})
	// 查询余额
	router.GET("/account/balance", func(context *gin.Context) {
		account.AccountGetBalance(context)
	})
	// 详细订单查询
	router.GET("/account/orderdetail", func(context *gin.Context) {
		account.AccountWaterDetail(context)
	})
	// 订单批量查询
	router.GET("/account/waterlist", func(context *gin.Context) {
		account.AccountWaterList(context)
	})
	// 账户充值
	router.POST("/account/charge", func(context *gin.Context) {
		account.AccountCharge(context)
	})
	// 账户消耗
	router.POST("/account/consume", func(context *gin.Context) {
		account.AccountConsume(context)
	})
	// 获取账户充值档位所有列表
	router.GET("/account/chargeGearMappingList", func(context *gin.Context) {
		account.ChargeGearMappingList(context)
	})
	// 获取账户充值具体某个档位信息
	router.GET("/account/chargeGearMappingInfo", func(context *gin.Context) {
		account.ChargeGearMappingInfo(context)
	})

	router.GET("/getAccessToken", func(context *gin.Context) {
		auth.GetAccessToken(context)
	})
	router.GET("/freshAccessToken", func(context *gin.Context) {
		auth.FreshToken(context)
	})
	router.GET("/saveAccessToken", func(context *gin.Context) {
		auth.SaveAccessToken(context)
	})

	// 积分余额查询
	router.GET("/jfAccount/balance", func(context *gin.Context) {
		account.JFAccountGetBalance(context)
	})
	// 积分账户充值
	router.POST("/jfAccount/charge", func(context *gin.Context) {
		account.JFAccountCharge(context)
	})
	// 积分账户消耗
	router.POST("/jfAccount/consume", func(context *gin.Context) {
		account.JFAccountConsume(context)
	})
	// 积分账户详细订单查询
	router.GET("/jfAccount/orderdetail", func(context *gin.Context) {
		account.JFAccountWaterDetail(context)
	})
	// 积分订单批量查询
	router.GET("/jfAccount/waterlist", func(context *gin.Context) {
		account.JFAccountWaterList(context)
	})

	// 设备相关
	router.POST("/device/operation", func(context *gin.Context) {
		device.DeviceOperation(context)
	})
	router.POST("/xiaomiDevice/operation", func(context *gin.Context) {
		device.XiaoMiDeviceOperation(context)
	})

	//支付相关
	router.POST("/pay/wxpay", func(context *gin.Context) {
		pay.WxPay(context)
	})
	router.GET("/pay/queryOrder", func(context *gin.Context) {
		pay.WxQueryOrder(context)
	})
	router.GET("/order/list", func(context *gin.Context) {
		book.GetOrder(context)
	})

	router.GET("/order/listDivideByTime", func(context *gin.Context) {
		book.GetOrderDivideByTime(context)
	})

	return router
}
