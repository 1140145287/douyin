package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	MysqlEngine *gorm.DB
	RedisEngine *redis.Client
)
