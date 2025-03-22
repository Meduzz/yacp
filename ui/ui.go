package ui

import (
	"net/http"

	"github.com/Meduzz/gml"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, tag gml.Tag) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, tag.Render())
}
