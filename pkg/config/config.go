package config

import (
	"gorm.io/gorm/logger"
	"time"
)
import "github.com/spf13/viper"

var MysqlOption MysqlOptions
var StoreMysqlOption MysqlOptions
var AccountMysqlOption MysqlOptions

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

func InitConf() (err error) {
	//初始化日志组件
	viper.SetConfigName("main")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config fatal : %s")
	}

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

	return nil
}
