package infra

import (
	"context"

	"github.com/TH04e22/go-chatroom/config"
	. "github.com/TH04e22/go-chatroom/internel/utils"
	"github.com/go-redis/redis/v8"
)

var RedisConn *redis.Client
var RedisCtx = context.Background()

func init() {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       0,
	})

	status := RedisConn.Ping(RedisCtx)
	_, err := status.Result()

	if err != nil {
		ErrorLog.Printf("error on redis initializing: %s", err.Error())
		panic(err)
	}
}
