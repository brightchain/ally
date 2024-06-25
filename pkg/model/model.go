package model

import (
	"ally/pkg/viperConf"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 连接管理器
type RDBManager struct {
	OpenTx bool     // 是否开启事务
	DsName string   // 数据源名称
	Db     *gorm.DB // 非事务实例
	Tx     *gorm.Tx // 事务实例
	Errors []error  // 操作过程中记录的错误
}

// db连接
var (
	MASTER = "db1"                    // 默认主数据源
	RDB    = map[string]*RDBManager{} // 初始化时加载数据源到集合
)

func InitDb() {
	dbConf := viperConf.Data.Sub("database")
	confMap := dbConf.AllSettings()

	for k, v := range confMap {
		var conf viperConf.Mysql
		//map[string]interface{}转结构体
		mapstructure.Decode(v, &conf)
		slog.Info("数据库信息", conf)
		connByConf(k, conf)
	}

}

func connByConf(key string, input viperConf.Mysql) {
	db, err := connectDB(input)
	if err != nil {
		slog.Error("数据库链接失败!", err)
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	rdb := &RDBManager{
		Db:     db,
		DsName: key,
	}
	RDB[key] = rdb
}

func connectDB(conf viperConf.Mysql) (*gorm.DB, error) {
	logConf := viperConf.Data.Sub("logger")
	filename := logConf.GetString("gormFile")
	level := logConf.GetString("filename")
	logOps := logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,       // Disable color
	}
	switch level {
	case "debug":
		logOps.LogLevel = logger.Info
	case "info":
		logOps.LogLevel = logger.Info
	case "warn":
		logOps.LogLevel = logger.Warn
	case "error":
		logOps.LogLevel = logger.Error
	default:
		logOps.LogLevel = logger.Warn
	}
	// 初始化GORM日志配置
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("logger.Setup err: %v", err)
	}
	newLogger := logger.New(log.New(f, "\r\n", log.LstdFlags), logOps)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DbName)
	conn, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		slog.Error("数据库连接失败", err)
		return nil, errors.New("数据库连接失败")
	}

	return conn, nil
}
