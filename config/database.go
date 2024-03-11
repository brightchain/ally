package config

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb() (*gorm.DB, error) {

	db, err := GetDbDatabase(GlobalConfig.GetString("mysql.db"))

	return db, err
}

func GetDbDatabase(database string) (*gorm.DB, error) {
	var conf AppConfig
	err := GlobalConfig.Unmarshal(&conf)
	if err != nil {
		slog.Error("数据库配置文件解析失败")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("数据库链接失败：%s", err)
	}

	return db, err
}
