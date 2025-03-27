package main

import (
	"flag"
	"log"

	"github.com/Meduzz/yacp/chat"
	"github.com/Meduzz/yacp/controller"
	"github.com/Meduzz/yacp/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	dataDir := flag.String("dir", "", "Specify the dir to store prompt history to")
	llmHost := flag.String("url", "", "Specify the url to the llm")
	flag.Parse()

	r := gin.Default()

	// Initialize storage
	err := storage.InitStorage(*dataDir)
	err = chat.InitChatService(*llmHost)

	if err != nil {
		log.Fatalf("Failed to create chat service: %v", err)
	}

	// Register routes
	controller.RegisterRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
