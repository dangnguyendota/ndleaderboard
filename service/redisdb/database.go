package redisdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Error struct {
	f string
	err error
}

func NewError(f string, err error) *Error {
	return &Error{
		f:   f,
		err: err,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("redis db error on function %s: %s", e.f, e.err)
}

type RedisDB struct {
	redis.Client
}

func NewRedisDB() *RedisDB {
	return &RedisDB{}
}

func (r *RedisDB) HealthCheck(ctx context.Context) error {
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return NewError("HealthCheck", err)
	}
	return nil
}

func (r *RedisDB) RemoveLeaderboard(ctx context.Context, leaderboardID string) error {
	err := r.Client.Del(ctx, leaderboardID).Err()
	if err != nil {
		return NewError("RemoveLeaderboard", err)
	}
	return nil
}



