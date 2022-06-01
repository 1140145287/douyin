package database

import (
	"douyin/global"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var rdb *redis.Client

// GetDBConnect get the mysql connection poll
func GetDBConnect() (conn *gorm.DB) {
	if db != nil {
		return db
	}
	dataSourceName := config.AppConfig.Get("mysql.dataSourceName").(string)
	var err error
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "dy_", // table name prefix, table for `User` would be `dy_user`
			SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		log.Panic(err)
	}
	//database connection pool config
	sqlDB, _ := db.DB()
	log.Printf("数据库为mysql , 数据库链接为%s", dataSourceName)
	sqlDB.SetConnMaxLifetime(time.Hour)
	//空闲连接池
	sqlDB.SetMaxIdleConns(10)
	//最大连接数
	sqlDB.SetMaxOpenConns(100)
	return db
}

// GetRedisConnect get the redis connection pool
func GetRedisConnect() *redis.Client {
	if rdb != nil {
		return rdb
	}
	log.Printf("Redis连接地址%s", global.RedisSetting.Url)
	rdb = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Url,
		Password: global.RedisSetting.Password,
		DB:       global.RedisSetting.Database,
	})
	return rdb
}
