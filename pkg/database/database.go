package database

import (
	"ally/pkg/config"
	"ally/pkg/logger"
	"database/sql"
	"log/slog"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
)

var DB *sql.DB

func Initialize() {
	var err error
	var conf config.Mysql
	mysqlConf := config.Data.Sub("database.db1")
	mysqlMap := mysqlConf.AllSettings()
	mapstructure.Decode(mysqlMap, &conf)

	config := mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Addr:                 conf.Host + ":" + strconv.Itoa(conf.Port),
		Net:                  "tcp",
		DBName:               conf.DbName,
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	slog.Info("数据库信息", conf)
	logger.LogError("数据库打开失败", err)

	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError("数据库连接失败！", err)
}
