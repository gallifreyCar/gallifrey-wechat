package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RDB *redis.Client

func InitRedis() {
	ctx := context.Background()
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	RDB := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis connect failed")
		return
	}

	fmt.Println("redis connect success")

}
