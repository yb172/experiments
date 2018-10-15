package redisq

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// NewClient creates new redis client
func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
