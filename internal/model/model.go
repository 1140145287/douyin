package model

import (
	"douyin/global"
	"douyin/pkg/setting"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// GetDBConnect get the mysql connection poll
func NewMysqlEngine(MysqlSetting *setting.MysqlSettingS) (*gorm.DB, error) {
	dataSourceName := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(dataSourceName,
		MysqlSetting.UserName,
		MysqlSetting.Password,
		MysqlSetting.Host,
		MysqlSetting.DBName,
		MysqlSetting.Charset,
		MysqlSetting.ParseTime,
	)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "dy_", // table name prefix, table for `User` would be `dy_user`
			SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB() //配置连接池
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(MysqlSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(MysqlSetting.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	//空闲连接池
	return db, nil
}

// GetRedisConnect get the redis connection pool
func NewRedisEngine(RedisSetting *setting.RedisSettingS) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Url,
		Password: global.RedisSetting.Password,
		DB:       global.RedisSetting.Database,
	})
	return rdb, nil
}
