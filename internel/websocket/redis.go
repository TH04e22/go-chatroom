package websocket

import (
	"context"

	"github.com/TH04e22/go-chatroom/config"
	. "github.com/TH04e22/go-chatroom/internel/utils"
	"github.com/go-redis/redis/v8"
)

var redisConn *redis.Client
var redisCtx = context.Background()

func init() {
	redisConn = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       0,
	})

	status := redisConn.Ping(redisCtx)
	_, err := status.Result()

	if err != nil {
		ErrorLog.Printf("error on redis initializing: %s", err.Error())
		panic(err)
	}
}
