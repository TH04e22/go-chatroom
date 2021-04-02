package chatroom

import (
	"context"
	"encoding/json"
	"time"

	"github.com/TH04e22/go-chatroom/config"
	"github.com/TH04e22/go-chatroom/internel/domain/entity"
	. "github.com/TH04e22/go-chatroom/internel/infra/redis"
	. "github.com/TH04e22/go-chatroom/internel/utils"
	"github.com/go-redis/redis/v8"
)

type Hub struct {
	CheckPeriod int
	Close       context.CancelFunc
}

// Serve Hub consume data in redis list and publish data to specific room
func (h *Hub) Serve(cancelCtx context.Context) {
	for {
		select {
		case <-cancelCtx.Done():
			DebugLog.Println("Hub close gracefully...")
			break
		default:
			message, _ := h.readMessage()
			h.publishMessage(message)
		}
	}
}

func (h *Hub) readMessage() (message *entity.Message, err error) {
	payload := RedisConn.BRPop(RedisCtx, time.Duration(h.CheckPeriod)*time.Second, config.Redis.HubPrefix)

	content, err := payload.Result()

	if err == redis.Nil {
		InfoLog.Println("The hub buffer is empty ... ")
		return nil, err
	} else if err != nil {
		ErrorLog.Printf("%v", err)
		panic(err)
	}

	message = new(entity.Message)
	json.Unmarshal([]byte(content[1]), message)

	return message, nil
}

func (h *Hub) publishMessage(message *entity.Message) error {
	channelName := config.Redis.RoomPrefix + ":" + message.RoomName

	payload, err := json.Marshal(message)
	if err != nil {
		ErrorLog.Printf("hub publish error on message marshal: %v", err)
		return err
	}

	execResult := RedisConn.Publish(RedisCtx, channelName, payload)
	if err := execResult.Err(); err != nil {
		ErrorLog.Printf("hub publish error: %v", err)
		return err
	}

	return nil
}
