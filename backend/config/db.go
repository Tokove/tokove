package config

import (
	"backend/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	// 打开数据库
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to init database: " + err.Error())
	}

	// 获取sqldb底层
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to configure database: " + err.Error())
	}

	// 配置数据库
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.Db = db
	log.Println("Init database successfully")
}
