package config

import (
	"github.com/spf13/viper"
	"github.com/zlilemon/gin_auto/pkg/log"
	"gorm.io/gorm/logger"
	"time"
)

var MysqlOption MysqlOptions
var StoreMysqlOption MysqlOptions
var AccountMysqlOption MysqlOptions

var WxPayOption WxPayOptions

var XiaomiDeviceOption XiaomiDeviceOptions
var TuyaDeviceOption TuyaDeviceOptions

var JWTSecretKey string

type MysqlOptions struct {
	Host                  string
	Port                  int64
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
	Logger                logger.Interface
}

type WxPayOptions struct {
	WxURI                      string
	WxAppID                    string
	WxSecret                   string
	WxMchId                    string
	WxPayNotifyUrl             string
	MchCertificateSerialNumber string
	MchAPIv3Key                string
	PrivateKeyPath             string
}

type XiaomiDeviceOptions struct {
	XiaomiURI                 string
	XiaomiHomeAssistanceToken string
}

type TuyaDeviceOptions struct {
	TuyaURI string
}

func InitConf() (err error) {
	//初始化日志组件
	viper.SetConfigName("main")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config fatal : %s")
	}

	// JWT secret config
	JWTSecretKey = viper.Get("jwt.secret_key").(string)

	// 读取 wxapp mysql的db配置
	MysqlOption.Host = viper.Get("mysql.ip").(string)
	MysqlOption.Port = viper.Get("mysql.port").(int64)
	MysqlOption.Username = viper.Get("mysql.user").(string)
	MysqlOption.Password = viper.Get("mysql.password").(string)
	MysqlOption.Database = viper.Get("mysql.database").(string)

	// 读取 store mysql的db配置
	StoreMysqlOption.Host = viper.Get("store_mysql.ip").(string)
	StoreMysqlOption.Port = viper.Get("store_mysql.port").(int64)
	StoreMysqlOption.Username = viper.Get("store_mysql.user").(string)
	StoreMysqlOption.Password = viper.Get("store_mysql.password").(string)
	StoreMysqlOption.Database = viper.Get("store_mysql.database").(string)

	// 读取 account mysql的db配置
	AccountMysqlOption.Host = viper.Get("account_mysql.ip").(string)
	AccountMysqlOption.Port = viper.Get("account_mysql.port").(int64)
	AccountMysqlOption.Username = viper.Get("account_mysql.user").(string)
	AccountMysqlOption.Password = viper.Get("account_mysql.password").(string)
	AccountMysqlOption.Database = viper.Get("account_mysql.database").(string)

	// 读取微信支付配置
	WxPayOption.WxURI = viper.Get("wx_config.wxUri").(string)
	WxPayOption.WxAppID = viper.Get("wx_config.wxAppID").(string)
	WxPayOption.WxSecret = viper.Get("wx_config.wxSecret").(string)
	WxPayOption.WxMchId = viper.Get("wx_config.wxMchId").(string)
	WxPayOption.WxPayNotifyUrl = viper.Get("wx_config.wxPayNotifyUrl").(string)
	WxPayOption.MchCertificateSerialNumber = viper.Get("wx_config.mchCerificateSerialNumber").(string)
	WxPayOption.MchAPIv3Key = viper.Get("wx_config.mchAPIv3Key").(string)
	WxPayOption.PrivateKeyPath = viper.Get("wx_config.privateKeyPath").(string)

	// xiaomi
	XiaomiDeviceOption.XiaomiURI = viper.Get("xiaomi_device.xiaomiUri").(string)
	XiaomiDeviceOption.XiaomiHomeAssistanceToken = viper.Get("xiaomi_device.xiaomiHomeAssistanceToken").(string)

	//tuya
	TuyaDeviceOption.TuyaURI = viper.Get("tuya_device.tuyaUri").(string)

	return nil
}
