package redis

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"hero/configs"
	"hero/enums"
	"hero/pkg/logger"
)

var (
	defaultClient *redis.Client
)

func init() {
	ctx := context.Background()
	addr := configs.Get("database.redis_url")
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 5,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Print("redis init", err.Error())
		panic(enums.ErrorRedisPingFail)
	}
	logger.Print("redis init", "success")
	defaultClient = client
}
func Client() *redis.Client {
	return defaultClient
}
