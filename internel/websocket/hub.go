package websocket

import (
	"context"
	"time"

	"github.com/TH04e22/go-chatroom/config"
	. "github.com/TH04e22/go-chatroom/internel/utils"
	"github.com/go-redis/redis/v8"
)

const (
	hubCheckPeriod = 60
)

type Hub struct {
	Close context.CancelFunc
}

// Serve Hub consume data in redis list and publish data to specific room
func (h *Hub) Serve(cancelCtx context.Context) {
	for {
		select {
		case <-cancelCtx.Done():
			DebugLog.Println("Hub close gracefully...")
			break
		default:
			messages, err := redisConn.BRPop(redisCtx, hubCheckPeriod*time.Second, config.Redis.HubPrefix).Result()
			if err == redis.Nil {
				InfoLog.Println("The hub buffer is empty ... ")
			} else if err != nil {
				ErrorLog.Printf("%v", err)
				panic(err)
			}

			InfoLog.Printf("Get message: %s", messages[1])
		}
	}
}
