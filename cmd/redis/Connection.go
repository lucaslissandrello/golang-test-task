package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(host, port, user, password string, db int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: user,
		Password: password,
		DB:       db,
	})

	return redisClient
}
