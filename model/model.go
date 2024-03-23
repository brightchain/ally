package model

import (
	"ally/config"
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

// 数据库配置
type DBConfig struct {
	DsName   string `json:"dsName"`   // 数据源名称
	Host     string `json:"host"`     // 地址IP
	Port     int    `json:"port"`     // 数据库端口
	Database string `json:"database"` // 数据库名称
	Username string `json:"username"` // 账号
	Password string `json:"password"` // 密码
}

// db连接
var (
	MASTER = "db1"                    // 默认主数据源
	RDBs   = map[string]*RDBManager{} // 初始化时加载数据源到集合
)

func InitDb() {

	dbConf := config.GlobalConfig.Sub("database")
	confMap := dbConf.AllSettings()

	for _, v := range confMap {
		var conf DBConfig
		//map[string]interface{}转结构体
		mapstructure.Decode(v, &conf)
		slog.Info("数据库信息", conf.Database)
		connByConf(conf)
	}

}

func connByConf(input DBConfig) {
	db, err := connDB(input)
	if err != nil {
		slog.Error("数据库链接失败!", err)
		return
	}
	if len(input.DsName) == 0 {
		input.DsName = MASTER
	}
	rdb := &RDBManager{
		Db:     db,
		DsName: input.DsName,
	}
	RDBs[input.DsName] = rdb
}

func connDB(conf DBConfig) (*gorm.DB, error) {
	logConf := config.GlobalConfig.Sub("logger")
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
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	conn, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		slog.Error("数据库连接失败", err)
		return nil, errors.New("数据库连接失败")
	}
	sqlDB, err := conn.DB()
	if err != nil {
		slog.Error("数据库连接失败", err)
		return nil, errors.New("数据库连接失败")
	}
	//设置连接池
	sqlDB.SetConnMaxLifetime(5 * time.Second)
	//空闲
	sqlDB.SetMaxIdleConns(3)
	//打开
	sqlDB.SetMaxOpenConns(5)

	return conn, nil
}
