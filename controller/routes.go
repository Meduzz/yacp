package controller

import (
	"github.com/gin-gonic/gin"
)

// TODO
// - new test endpoint that loads a ui
// - new test endpoint that streams llm responses.
func RegisterRoutes(r *gin.Engine) {
	r.GET("/", Index)         // list chats
	r.GET("/new", CreateForm) // create a chat
	r.POST("/new", CreateHandler)
	r.GET("/chat/:chat", ShowChat)                     // show/open a chat
	r.POST("/chat/:chat", HandleChat)                  // post a message to a chat
	r.DELETE("/chat/:chat", RemoveChat)                // delete an entire chat
	r.DELETE("/message/:chat/:message", RemoveMessage) // delete a message
}
