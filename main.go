package main

import (
	"flag"
	"log"

	"github.com/Meduzz/yacp/controller"
	"github.com/Meduzz/yacp/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	dataFile := flag.String("file", "", "Specify the file to store prompt history to")
	flag.Parse()

	r := gin.Default()

	// Initialize storage
	var badgerStorage *storage.BadgerStorage
	if *dataFile == "" {
		badgerStorage = storage.NewInMemoryBadgerStorage()
	} else {
		badgerStorage = storage.NewFileBasedBadgerStorage(*dataFile)
	}

	// Register routes
	controller.RegisterRoutes(r, badgerStorage)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
