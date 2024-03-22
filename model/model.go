package model

import (
	"ally/config"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

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
	DsName   string // 数据源名称
	Host     string // 地址IP
	Port     int    // 数据库端口
	Database string // 数据库名称
	Username string // 账号
	Password string // 密码
}

// db连接
var (
	MASTER = "db1"                    // 默认主数据源
	RDBs   = map[string]*RDBManager{} // 初始化时加载数据源到集合
)

func InitDb() {

	db1Con := config.GlobalConfig.Sub("database")
	confMap := db1Con.AllSettings()

	for _, v := range confMap {
		// 这里假设你有一个 DBConfig 结构体，需要根据实际结构体字段进行解码或转换
		//var conf DBConfig
		// 进行类型断言或转换，注意这里取决于v的实际内容和DBConfig结构体的定义
		// 通常情况下，这可能需要通过json.Unmarshal或其他方式解析
		fmt.Print("配置%T", v)
		fmt.Print("配置%s", v["username"])

		// 现在你可以安全地使用conf变量了
		//connByConf(conf)
	}

	// var db1Conf DBConfig
	// db1Con.Unmarshal(&db1Conf)
	// db1, err := connDB(db1Conf)
	// if err != nil {
	// 	slog.Error("数据库链接失败 %s ", err.Error())
	// 	return
	// }
	// if len(db1Conf.DsName) == 0 {
	// 	db1Conf.DsName = MASTER
	// }
	// rdb1 := &RDBManager{
	// 	Db:     db1,
	// 	DsName: db1Conf.DsName,
	// }
	// RDBs[db1Conf.DsName] = rdb1
	// db2Con := config.GlobalConfig.Sub("database.db2")
	// var db2Conf DBConfig
	// db2Con.Unmarshal(&db2Conf)
	// db2, err := connDB(db2Conf)
	// if err != nil {
	// 	slog.Error("数据库链接失败 %s ", err.Error())
	// 	return
	// }
	// if len(db2Conf.DsName) == 0 {
	// 	db2Conf.DsName = MASTER
	// }
	// slog.Info("数据库连接器", db2Conf.DsName)
	// rdb2 := &RDBManager{
	// 	Db:     db2,
	// 	DsName: db2Conf.DsName,
	// }
	// RDBs[db2Conf.DsName] = rdb2
}

func connByConf(input DBConfig) {
	db, err := connDB(input)
	if err != nil {
		slog.Error("数据库链接失败 %s ", err.Error())
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
		logOps.LogLevel = logger.Warn
	default:
		logOps.LogLevel = logger.Warn
	}
	// 初始化GORM日志配置
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("logger.Setup err: %v", err)
	}
	newLogger := logger.New(log.New(f, "\r\n", log.LstdFlags), logOps)
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	dialector := mysql.New(mysql.Config{
		DSN:                       dbURI, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})
	conn, err := gorm.Open(dialector, &gorm.Config{
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
