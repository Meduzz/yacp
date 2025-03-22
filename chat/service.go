package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OllamaService struct {
	host string
}

func NewOllamaService(host string) *OllamaService {
	return &OllamaService{host}
}

func (os *OllamaService) HandleChat(c *gin.Context) {
	// Ollama-specific logic here
	c.JSON(http.StatusOK, gin.H{"message": "Chat handled by Ollama"})
}
