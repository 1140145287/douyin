package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	MysqlEngine *gorm.DB
	RedisEngine *redis.Client
	OSSEngine   *oss.Bucket
)
