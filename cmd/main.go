package main

import (
	"context"
	"net/http"

	"github.com/TH04e22/go-chatroom/internel/websocket"
	"github.com/gin-gonic/gin"
)

func main() {
	// run hub
	hubCtx, hubCancelFunc := context.WithCancel(context.Background())
	hub := websocket.Hub{Close: hubCancelFunc}
	defer hub.Close()
	go hub.Serve(hubCtx)

	// run http server
	router := gin.Default()
	router.Static("/assets", "../web/assets")
	router.LoadHTMLGlob("../web/views/*")

	router.GET("/chatroom", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chatroom.htm", nil)
	})

	router.GET("/ws/client", func(c *gin.Context) {
		websocket.ServeWs(c.Writer, c.Request)
	})

	router.Run()
}
