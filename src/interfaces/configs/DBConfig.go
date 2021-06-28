package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 配置都需要放到k8s里面
type DBConfig struct{}

func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

func (c *DBConfig) GormDB() *gorm.DB {
	dsn := "root:123456@tcp(49.232.252.187:3306)/jstock?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	return db
}
