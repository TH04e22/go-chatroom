package router

import (
	"net/http"

	"github.com/TH04e22/go-chatroom/internel/websocket"
	"github.com/gin-gonic/gin"
)

var App *gin.Engine

func init() {
	App = gin.Default()

	App.Static("/assets", "../web/assets")
	App.LoadHTMLGlob("../web/views/*")

	root := App.Group("/")
	{
		root.GET("/chatroom", func(c *gin.Context) {
			c.HTML(http.StatusOK, "chatroom.html", nil)
		})

		root.GET("/ws/client", func(c *gin.Context) {
			websocket.ServeWs(c.Writer, c.Request)
		})
	}

	user := App.Group("/user")
	{
		user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

		user.GET("/sign-up", func(c *gin.Context) {
			c.HTML(http.StatusOK, "signup.html", nil)
		})
	}

}
