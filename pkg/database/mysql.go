package database

import (
	"fmt"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zlilemon/gin_auto/pkg/config"

	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB
var StoreDB *gorm.DB
var AccountDB *gorm.DB

// Options defines optsions for mysql database.
type Options struct {
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

func ConnectMysql() {

	/*
		dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s`,
			mysqlOption.Username,
			mysqlOption.Password,
			mysqlOption.Host,
			mysqlOption.Port,
			mysqlOption.Database,
			true,
			"Local")
	*/
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s`,
		config.MysqlOption.Username,
		config.MysqlOption.Password,
		config.MysqlOption.Host,
		config.MysqlOption.Port,
		config.MysqlOption.Database,
		true,
		"Local")

	dbProxy, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if dbProxy.Error != nil {
		fmt.Printf("database error %v", dbProxy.Error)
	}

	DB = dbProxy

	dsnStore := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s`,
		config.StoreMysqlOption.Username,
		config.StoreMysqlOption.Password,
		config.StoreMysqlOption.Host,
		config.StoreMysqlOption.Port,
		config.StoreMysqlOption.Database,
		true,
		"Local")

	dbStoreProxy, err := gorm.Open("mysql", dsnStore)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if dbStoreProxy.Error != nil {
		fmt.Printf("database error %v", dbProxy.Error)
	}

	dsnAccount := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s`,
		config.AccountMysqlOption.Username,
		config.AccountMysqlOption.Password,
		config.AccountMysqlOption.Host,
		config.AccountMysqlOption.Port,
		config.AccountMysqlOption.Database,
		true,
		"Local")

	dbAccountProxy, err := gorm.Open("mysql", dsnAccount)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if dbAccountProxy.Error != nil {
		fmt.Printf("database error %v", dbProxy.Error)
	}

	DB = dbProxy
	StoreDB = dbStoreProxy
	AccountDB = dbAccountProxy
}

// New create a new gorm db instance with the given options.
/*
func New(opts *Options) (*gorm.DB,  error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
		true,
		"Local")

	var err error
	Eloquent, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: opts.Logger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := Eloquent.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return Eloquent, err
}
*/
