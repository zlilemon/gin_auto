package pay

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"github.com/zlilemon/gin_auto/app/account"
	"github.com/zlilemon/gin_auto/app/billing"
	"github.com/zlilemon/gin_auto/app/user"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/config"
	"github.com/zlilemon/gin_auto/pkg/log"
	"github.com/zlilemon/gin_auto/pkg/wxpay"
	"math"
	"time"
)

type IService interface {
	SCreateOrderId() (err error)
	SPrepay(c *gin.Context, openid string, tranAmt int64) (err error)
}

type Service struct {
	//repo  		IRepository
}

type contentType struct {
	Mchid           *string    `json:"mchid"`
	Appid           *string    `json:"appid"`
	CreateTime      *time.Time `json:"create_time"`
	OutContractCode *string    `json:"out_contract_code"`
}

var PayService = new(Service)

func (s *Service) SWxPayV3(c *gin.Context, req WxPayReq, resp *WxPayResp) (err error) {

	//currentTime := time.Now()
	//outTradeNo := strconv.FormatInt(currentTime.Unix(), 10)
	//根据当前时间戳，生成 业务订单号
	//outTradeNo := currentTime.Format("2006-01-02 15:04:05")

	//log.Infof("currentTime:%s, outTradeNo:%s", currentTime, outTradeNo)
	svc := jsapi.JsapiApiService{Client: wxpay.WechatPayClient}

	log.Infof("out_trade_no :%s, core_out_trade_no:%s", req.OutTradeNo, core.String(req.OutTradeNo))
	//获得prepay_id
	wxpayResp, result, err := svc.PrepayWithRequestPayment(c,
		jsapi.PrepayRequest{
			Appid:       core.String(config.WxPayOption.WxAppID),
			Mchid:       core.String(config.WxPayOption.WxMchId),
			Description: core.String("abcde"),
			OutTradeNo:  core.String(req.OutTradeNo),
			Attach:      core.String(req.PayInfo),
			NotifyUrl:   core.String(config.WxPayOption.WxPayNotifyUrl),
			Amount: &jsapi.Amount{
				Total: core.Int64(req.Amount),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(req.OpenId),
			},
		})

	if err == nil {
		log.Infof("PrepayWithRequestPayment Success, resp:%s, result:%s", wxpayResp, result)
		log.Infof("prepay_id:%s", *wxpayResp.PrepayId)
	} else {
		log.Errorf("PrepayWithRequestPayment error, resp:%s", wxpayResp)
		log.Errorf("PrepayWithRequestPayment error, result:%s", result)
		log.Errorf("PrepayWithRequestPayment error, err:%s", err)

		return
	}

	resp.PrepayId = *wxpayResp.PrepayId
	resp.TimeStamp = *wxpayResp.TimeStamp
	resp.NonceStr = *wxpayResp.NonceStr
	resp.Package = *wxpayResp.Package
	resp.PaySign = *wxpayResp.PaySign

	return
}

func (s *Service) SWxPayNotify(c *gin.Context, req WxpayNotifyReq, resp *comm.Result) (err error) {
	log.Infof("enter SWxPayNotify ")

	log.Infof("wxpayNotify, id:%s", req.ID)
	// 1. 初始化商户API v3 Key及微信支付平台证书

	mchAPIv3Key := config.WxPayOption.MchAPIv3Key

	ctx := context.Background()

	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(config.WxPayOption.PrivateKeyPath)

	log.Infof("comment -- 2")

	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey,
		config.WxPayOption.MchCertificateSerialNumber, config.WxPayOption.WxMchId, mchAPIv3Key)

	log.Infof("comment -- 3")

	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(config.WxPayOption.WxMchId)
	handler := notify.NewNotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	log.Infof("comment -- 4")

	//content := new(contentType)
	//content := make(map[string]interface{})
	transaction := new(payments.Transaction)

	//notifyReq, err := handler.ParseNotifyRequest(c, c.Request, content)
	notifyReq, err := handler.ParseNotifyRequest(context.Background(), c.Request, transaction)
	//notifyReq, err := handler.ParseNotifyRequest(c, c.Request , content)
	if err != nil {
		log.Errorf("ParseNotifyRequest error, err_msg:%v", err)
		return
	}

	log.Infof("transaction :%+v", transaction)
	log.Infof("transaction :%s", transaction)
	log.Infof("notifyReq:%+v", notifyReq)
	outTradeNo := *transaction.OutTradeNo
	openId := *transaction.Payer.Openid
	channelOrderNo := *transaction.TransactionId
	tradeState := *transaction.TradeState
	successTime := *transaction.SuccessTime
	payerTotal := *transaction.Amount.PayerTotal

	log.Infof("outTradeNo:%s, openId:%s, channelOrderNo:%s, tradeState:%s, successTime:%s, payerTotal:%d",
		outTradeNo, openId, channelOrderNo, tradeState, successTime, payerTotal)
	// 通过回调函数，把渠道订单号更新到原来到交易订单中
	// 先查询原来交易的订单信息
	//bookOrderReq := book.BookOrderReq{}
	billingReq := billing.BillingInfoReq{}
	billingReq.OpenId = openId
	billingReq.OutTradeNo = outTradeNo
	billingReq.NotifyStatus = tradeState
	billingReq.Page = 0      // init default page = 0
	billingReq.PageSize = 20 // init default pageSize = 20

	billingRespList := make([]*billing.BillingInfoResp, 0)
	billingRespList, err = billing.BillingService.SGetBillingInfo(c, billingReq)
	if err != nil {
		log.Errorf("SGetOrder error, err_msg:%v", err)
		return
	}
	if len(billingRespList) == 0 {
		// 查找不到具体的订单记录
		err = errors.New("SGetOrder resp is NULL")
		log.Errorf("SGetOrder error, err_msg:%v", err)
		return
	} else if len(billingRespList) > 1 {
		// 查找到多于一条的订单记录
		err = errors.New("SGetOrder resp record more than 1")
		log.Errorf("SGetOrder error, err_msg:%v", err)
		return
	}

	// 更新对应订单号为支付成功，并更新渠道回调订单号
	bookOrderUpdateReq := billing.UpdateBillingReq{}
	bookOrderUpdateReq.OutTradeNo = outTradeNo
	bookOrderUpdateReq.OpenId = openId
	bookOrderUpdateReq.ChannelOrderNo = channelOrderNo
	bookOrderUpdateReq.PayStatus = billingRespList[0].PayStatus
	bookOrderUpdateReq.NotifyStatus = tradeState

	//err = book.BookService.SUpdateOrder(c, bookOrderUpdateReq)
	err = billing.BillingService.SUpdateBilling(c, bookOrderUpdateReq)
	if err != nil {
		log.Errorf("SUpdateOrder error, err_msg:%v", err)
		return
	}

	log.Infof("SUpdateBilling Success - ")

	// 针对账户充值accountSave类型的支付
	// 支付通知接口成功后，直接调用账户充值接口进行账户充值
	bookOrderResp := billingRespList[0]
	if bookOrderResp.OrderType == "accountSave" {
		// 账户充值，调用账户充值模块进行充值
		accountSaveReq := account.ChargeReq{}
		accountSaveResp := account.ChargeResp{}

		accountSaveReq.OpenId = openId
		accountSaveReq.Currency = bookOrderResp.Currency
		accountSaveReq.Amount = bookOrderResp.Amount
		accountSaveReq.PayMethod = bookOrderResp.PayMethod
		accountSaveReq.OutTradeNO = bookOrderResp.OutTradeNo
		accountSaveReq.ChannelOrderNo = bookOrderResp.ChannelOrderNo

		err = account.AccountService.SAccountCharge(c, accountSaveReq, &accountSaveResp)
		if err != nil {
			log.Errorf("SAccountCharge error, err_msg:%v", err)
			return
		}

		log.Infof("SAccountCharge Success, openid:%s, out_trade_no:%s, channel_order_no:%s, currency:%s, amount:%d",
			openId, accountSaveReq.OutTradeNO, accountSaveReq.ChannelOrderNo, accountSaveReq.Currency, accountSaveReq.Amount)
	} else {
		// 记录排行信息
		addRankingReq := user.AddRankingReq{}
		addRankingReq.OpenId = openId
		durationUnixTime := bookOrderResp.EndUnixTime - bookOrderResp.BeginUnixTime
		durationTime := math.Ceil(float64(durationUnixTime) / 3600) //往上取整
		addRankingReq.DurationTime = int(durationTime)
		addRankingReq.BookTimes = 1

		// 获取用户的nick_name 和 avatar_url
		userInfoModel, err := user.UserService.SGetUserInfo(c, openId)
		if err != nil {
			log.Errorf("SGetUserInfo error, errMsg:%s", err)
		} else {
			// 获取用户信息成功，赋值具体的字段内容
			log.Infof("userInfoModel len:%d, nickName:%s, avaUrl:%s", len(userInfoModel), userInfoModel[0].NickName, userInfoModel[0].AvatarUrl)
			addRankingReq.NickName = userInfoModel[0].NickName
			addRankingReq.AvatarUrl = userInfoModel[0].AvatarUrl
		}

		err = user.UserService.SAddRanking(c, addRankingReq)
		// 记录失败，也只是打印日志，并不阻塞支付的过程
	}

	return
}

func (s *Service) SWxQueryOrder(c *gin.Context, req QueryReq, resp *comm.Result) (err error) {
	log.Infof("transaction_id:%s, mcdid:%s", req.TransactionId, req.MchId)

	svc := jsapi.JsapiApiService{Client: wxpay.WechatPayClient}
	wxpayResp, result, err := svc.QueryOrderById(c,
		jsapi.QueryOrderByIdRequest{
			TransactionId: core.String(req.TransactionId),
			Mchid:         core.String(req.MchId),
		})
	if err == nil {
		log.Infof("QueryOrderByIdRequest Success:%s, resp, result:%s", wxpayResp, result)
	} else {
		log.Errorf("QueryOrderByIdRequest error, resp:%s", wxpayResp)
		log.Errorf("QueryOrderByIdRequest error, result:%s", result)
		log.Errorf("QueryOrderByIdRequest error, err:%s", err)
	}

	resp.Data = *wxpayResp.OutTradeNo

	return
}

func (s *Service) SDownloadPaltformCA(c *gin.Context, resp *comm.Result) {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775B6A45ACD588826D15E583A95F5DD********" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(config.WxPayOption.PrivateKeyPath)
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	//ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(c, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: client}
	downloadCAresp, result, err := svc.DownloadCertificates(c)
	log.Infof("status=%d resp=%s", result.Response.StatusCode, downloadCAresp)
}
