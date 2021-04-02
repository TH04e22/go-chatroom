package chatroom

import (
	"context"
	"net/http"
	"time"

	"github.com/TH04e22/go-chatroom/config"
	. "github.com/TH04e22/go-chatroom/internel/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 8) / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client used for get user websocket data and push data
// to redis list which websocket hub consuming...
type Client struct {
	conn   *websocket.Conn
	pubSub *redis.PubSub
	close  context.CancelFunc
}

// ReadPump  Read websocket from user websocket
func (client *Client) ReadPump(ctx context.Context) {
	defer func() {
		client.conn.Close()
	}()

	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := client.conn.ReadMessage()

			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					ErrorLog.Printf("websocket client read error: %v", err)
				} else {
					DebugLog.Printf("websocket client read error: %v", err)
				}
				client.close()
			}

			redisConn.LPush(redisCtx, config.Redis.HubPrefix, message)
		}
	}
}

// WritePump write data to user websocket
func (client *Client) WritePump(ctx context.Context) {
	defer func() {
		client.conn.Close()
	}()

	redisChannel := client.pubSub.ChannelSize(10)
	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case <-ctx.Done():
			return
		case channelMsg := <-redisChannel:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.TextMessage, []byte(channelMsg.Payload)); err != nil {
				ErrorLog.Printf("websocket client write payload error: %v", err)
				ticker.Stop()
				client.close()
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				ErrorLog.Printf("websocket client write ping message error: %v", err)
				ticker.Stop()
				client.close()
			}
		}
	}
}

// ServeWs upgrade http request to websocket connection and
// create a client to manage this web connection
func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ErrorLog.Println(err.Error())
		return
	}

	InfoLog.Println("Websocket client connect established")
	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{conn: conn,
		pubSub: redisConn.Subscribe(redisCtx, "room:PHP", "room:GO", "room:Javascript"),
		close:  cancel}

	go client.ReadPump(ctx)
	go client.WritePump(ctx)
}
