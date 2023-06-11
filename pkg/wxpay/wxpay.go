package wxpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"github.com/zlilemon/gin_auto/pkg/config"
	"github.com/zlilemon/gin_auto/pkg/log"
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
