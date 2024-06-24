package database

import (
	"ally/pkg/logger"
	"ally/pkg/viperConf"
	"database/sql"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize() {
	var err error
	var conf viperConf.AppConfig
	err = viperConf.Data.Unmarshal(&conf)
	if err != nil {
		logger.LogError("配置文件解析失败", err)
		return
	}
	config := mysql.Config{
		User:                 conf.Mysql.User,
		Passwd:               conf.Mysql.Password,
		Addr:                 conf.Mysql.Host + strconv.Itoa(conf.Mysql.Port),
		Net:                  "tcp",
		DBName:               conf.Mysql.Db,
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())

	logger.LogError("数据库打开失败", err)

	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError("数据库连接失败！", err)
}
