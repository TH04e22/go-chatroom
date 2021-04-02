package main

import (
	"context"

	"github.com/TH04e22/go-chatroom/internel/router"
	"github.com/TH04e22/go-chatroom/internel/websocket"
)

func main() {
	// run hub
	hubCtx, hubCancelFunc := context.WithCancel(context.Background())
	hub := websocket.Hub{Close: hubCancelFunc}
	defer hub.Close()
	go hub.Serve(hubCtx)

	// run http server
	app := router.App
	app.Run()
}
