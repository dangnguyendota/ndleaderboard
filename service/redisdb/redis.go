package redisdb

import (
	"github.com/go-redis/redis/v8"
)

func init() {
}

type RedisDB struct {
	redis.Client
}

func NewRedisDB() *RedisDB {
	return &RedisDB{}
}

func (r *RedisDB) HealthCheck() {

}

