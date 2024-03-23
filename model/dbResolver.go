package model

import (
	"ally/config"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var DbRes *gorm.DB

func InitDbResolver() {
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
	dbconf1 := config.GlobalConfig.Get("database.db1")
	var conf DBConfig
	mapstructure.Decode(dbconf1, &conf)
	dsn := getDsn(conf)
	slog.Info(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		slog.Error("数据库连接失败", err)
		return
	}

	dbConf := config.GlobalConfig.Sub("database")
	confMap := dbConf.AllSettings()
	if len(confMap) > 1 {
		var m = make([]gorm.Dialector, len(confMap)-1)
		var conf1 DBConfig
		i := 0
		for _, v := range confMap {
			mapstructure.Decode(v, &conf1)
			if conf1.DsName == "db1" {
				continue
			}
			dsn := getDsn(conf1)

			m[i] = mysql.Open(dsn)
			i++
		}

		db.Use(dbresolver.Register(dbresolver.Config{
			Sources: m,
		}).SetMaxIdleConns(10).
			SetConnMaxLifetime(time.Hour).
			SetMaxOpenConns(200),
		)
	}

	DbRes = db

	slog.Info("数据库初始化完成")
}

func getDsn(conf DBConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	return dsn
}
