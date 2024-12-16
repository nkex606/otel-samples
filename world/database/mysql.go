package database

import (
	"fmt"
	"log"
	"time"

	"otel-world/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMysql() {
	dbConf := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DBName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect to mysql failed: %+v\n", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dbConf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbConf.ConnMaxLifeTime))
}

func GetDB() *gorm.DB {
	return db
}
