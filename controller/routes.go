package controller

import (
	"github.com/Meduzz/yacp/storage"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, storage *storage.BadgerStorage) {
	r.GET("/", func(c *gin.Context) {
		// Index page logic here
	})

	r.POST("/chat", func(c *gin.Context) {
		// Chat endpoint logic here
	})
}
