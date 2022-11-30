package main

import (
	"demo/handler"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var passportDB *gorm.DB

func init() {
	initPassportDB()
}

func main() {
	handler.ReadFileAndQueryExtData(passportDB)
}

func initPassportDB() {
	dsn := "passport_r:XR7IhUROoZF1QaWIGZvX8H6tt@tcp(127.0.0.1:3406)/ks_sdk_server?charset=utf8&parseTime=True&loc=UTC"
	dialect := mysql.Open(dsn)
	c, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // 打印更多日志，包括所有 SQL 输出，设置为 warn 则只打印慢 SQL
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := c.DB()
	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(5)  // SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(50) // SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetConnMaxLifetime(time.Second * 180)
	} else {
		panic(err)
	}
	passportDB = c
}
