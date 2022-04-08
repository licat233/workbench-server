package gormx

import (
	"log"
	"os"
	"time"
	"workbench-server/common/global"
	"workbench-server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateGorm() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.DataSource(), // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		SkipInitializeWithVersion: false,          // 根据版本自动配置
	}
	gormOption := GormOption()
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormOption); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func GormOption() *gorm.Config {
	config := &gorm.Config{
		SkipDefaultTransaction:                   true, // 禁用默认事务
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true, // 创建并缓存预编译语句
	}
	newLogger := logger.New(
		NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)),
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, //慢日志
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	switch global.CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = newLogger.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = newLogger.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = newLogger.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = newLogger.LogMode(logger.Info)
	default:
		config.Logger = newLogger.LogMode(logger.Info)
	}
	return config
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.DataSource(), // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		SkipInitializeWithVersion: false,          // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), GormOption()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
