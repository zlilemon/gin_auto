package wxpay

import (
	"context"
	"gin_auto/pkg/config"
	"gin_auto/pkg/log"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var WechatPayClient *core.Client

func InitWxPayClient() {
	mchID := config.WxPayOption.WxMchId
	mchCertificateSerialNumber := config.WxPayOption.MchCertificateSerialNumber
	mchAPIv3Key := config.WxPayOption.MchAPIv3Key

	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(config.WxPayOption.PrivateKeyPath)

	if err != nil {
		log.Errorf("load merchant private key error")
	}

	ctx := context.Background()
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}

	WechatPayClient, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}
}
