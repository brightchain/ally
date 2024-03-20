package model

import (
	"ally/config"
	"log"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

func InitDb() {
	// 初始化GORM日志配置
	f, err := os.OpenFile(`./log/gorm.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("logger.Setup err: %v", err)
	}
	newLogger := logger.New(
		log.New(f, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dsn := config.GlobalConfig.GetString("mysqlList.dsn.1")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		slog.Error("数据库连接失败", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("数据库连接失败", err)
		return
	}
	dsnList := config.GlobalConfig.Get("mysqlList.dsn")
	if len(dsnList.([]interface{})) > 1 {
		var m = make([]gorm.Dialector, len(dsnList.([]interface{})))
		for i, v := range dsnList.([]interface{}) {
			dsnStr := v.(string)
			m[i] = mysql.Open(dsnStr)
		}
		db.Use(dbresolver.Register(dbresolver.Config{
			Sources: m,
		}))
	}

	//设置连接池
	sqlDB.SetConnMaxLifetime(5 * time.Second)
	//空闲
	sqlDB.SetMaxIdleConns(3)
	//打开
	sqlDB.SetMaxOpenConns(5)
	DB = db
}
