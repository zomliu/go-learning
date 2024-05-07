package main

import (
	"demo/handler"
	"demo/handler/passport"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlDB *gorm.DB

func main() {
	//processOmniServerReader()
	//processPassportReader()
	//handler.QueryIPLocal()
	processPassportDataV2()
}

func processOmniServerReader() {
	initDB("xg_readonly", "NIk8NnUjiLJkfA", "127.0.0.1", "3336", "xgsdk_db")
	//handler.WriteOrderToFile(mysqlDB)
	handler.QuerySpecificOrder(mysqlDB)
}

func processPassportReader() {
	initDB("passport_r", "XR7IhUROoZF1QaWIGZvX8H6tt", "127.0.0.1", "3406", "ks_sdk_server")
	handler.ReadFileAndQueryExtData(mysqlDB)
}

func processPassportDataV2() {
	initDB("passport_r", "XR7IhUROoZF1QaWIGZvX8H6tt", "127.0.0.1", "3406", "ks_sdk_server")
	p := &passport.PassportHandler{
		DB: mysqlDB,
		SourceFile: "/Users/leon/Downloads/moyu_data/2024_03.csv",
		TargetFile: "/Users/leon/Downloads/moyu_data/2024_03_result.csv",
	}
	p.ReadFileAndQuery()
	/*
	for i := 1; i <= 7; i++ {
		p := &passport.PassportHandler{
			DB: mysqlDB,
			SourceFile: fmt.Sprintf("/Users/leon/Downloads/moyu_data/part%d.csv", i),
			TargetFile: fmt.Sprintf("/Users/leon/Downloads/moyu_data/part%d_result.csv", i),
		}
		p.ReadFileAndQuery()
	}
	*/
}

func initDB(userName, password, host, port, dbName string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, host, port, dbName)
	//dsn := "passport_r:XR7IhUROoZF1QaWIGZvX8H6tt@tcp(127.0.0.1:3406)/ks_sdk_server?charset=utf8&parseTime=True&loc=UTC"
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
	mysqlDB = c
}
