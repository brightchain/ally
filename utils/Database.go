package utils

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDb(host string, port int, username string, password string, database string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&timeout=8s", username, password, host, port, database)
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		slog.Error("mysql lost", err)
		return nil, err
	}
	// 设置连接池中空闲连接的最大数量。
	Db.SetMaxIdleConns(1)
	// 设置打开数据库连接的最大数量。
	Db.SetMaxOpenConns(1)
	// 设置连接可复用的最大时间。
	Db.SetConnMaxLifetime(time.Second * 30)
	//设置连接最大空闲时间
	Db.SetConnMaxIdleTime(time.Second * 30)

	err = Db.Ping()
	if err != nil {
		slog.Error("mysql lost", err)
		return nil, err
	}

	return Db, err

}

func QueryOne(Db *sql.DB, query string) map[string]string {
	rows, err := Db.Query(query)
	if err != nil {
		slog.Error("mysql lost: %v", err)
		panic(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	if len(cols) > 1 {
		buffer := make([]interface{}, len(cols))
		data := make([][]byte, len(cols))
		dataKv := make(map[string]string, len(cols))
		for i, _ := range buffer {
			buffer[i] = &data[i]
		}

		for rows.Next() {
			rows.Scan(buffer...)
		}

		for k, col := range data {
			dataKv[cols[k]] = string(col)
			//fmt.Printf("%30s:\t%s\n", cols[k], col)
		}
		return dataKv

	}

	return nil

}

func QueryAll(Db *sql.DB, query string) []map[string]string {
	rows, err := Db.Query(query)
	if err != nil {
		slog.Error("mysql lost: %v", err)
		panic(err)
	}
	defer rows.Close()
	cols, _ := rows.Columns()
	if len(cols) > 1 {
		var ret []map[string]string
		for rows.Next() {
			buff := make([]interface{}, len(cols))
			data := make([][]byte, len(cols))
			for i, _ := range cols {
				buff[i] = &data[i]
			}
			rows.Scan(buff...)
			dataKv := make(map[string]string, len(cols))
			for k, col := range data {
				dataKv[cols[k]] = string(col)
			}
			ret = append(ret, dataKv)
		}
		return ret
	}

	return nil
}

func GetRows(Db *sql.DB, query string) *sql.Rows {
	rows, err := Db.Query(query)
	if err != nil {
		slog.Error("mysql lost: %v", err)
		panic(err)
	}
	defer rows.Close()
	return rows
}
