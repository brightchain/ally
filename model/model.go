package model

import (
	"ally/config"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb() {
	var conf config.AppConfig
	err := config.GlobalConfig.Unmarshal(&conf)
	if err != nil {
		slog.Error("配置文件解析失败", err)
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Db)
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
	//设置连接池
	sqlDB.SetConnMaxLifetime(5 * time.Second)
	//空闲
	sqlDB.SetMaxIdleConns(2)
	//打开
	sqlDB.SetMaxOpenConns(3)
	DB = db
}
