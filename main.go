package gin_auto

import (
	"gin_auto/pkg/database"
	"gin_auto/pkg/log"
	"net/http"
	"time"
)

func main() {
	log.Init(log.NewOptions())
	log.Info("log init success")

	// 读取配置文件
	config.InitConf()
	log.Infof("mysql username:%s", config.MysqlOption.Username)

	// 链接数据库实例
	database.ConnectMysql()

	// 初始化微信支付client
	wxpay.InitWxPayClient()

	router := router.InitRouter()

	s := &http.Server{
		Addr:           ":8090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
