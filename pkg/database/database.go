package database

import (
	"ally/pkg/config"
	"ally/pkg/logger"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize() {
	var err error
	dbConfig := mysql.Config{
		User:                 config.GetString("database.db.username"),
		Passwd:               config.GetString("database.db.password"),
		Addr:                 config.GetString("database.db.host") + ":" + config.GetString("database.db.prot"),
		Net:                  "tcp",
		DBName:               config.GetString("database.db.database"),
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", dbConfig.FormatDSN())
	logger.LogError("数据库打开失败", err)

	DB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	DB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	err = DB.Ping()
	logger.LogError("数据库连接失败！", err)
}
