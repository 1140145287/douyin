package dao

import (
	"douyin/pkg/setting"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)


// NewMysqlEngine return the mysql connection connection
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

// NewRedisEngine return the redis connection
func NewRedisEngine(RedisSetting *setting.RedisSettingS) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisSetting.Url,
		Password: RedisSetting.Password,
		DB:       RedisSetting.Database,
	})
	return rdb, nil
}

// NewOSSEngine return the oss connection
func NewOSSEngine(OSSettings *setting.OSSettingS) (*oss.Bucket, error) {
	// 创建OSSClient实例
	client, err := oss.New(OSSettings.Endpoint, OSSettings.AccessKeyId, OSSettings.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	// 使用 特定的 bucket
	bucket, err := client.Bucket(OSSettings.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
