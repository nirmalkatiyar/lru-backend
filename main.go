package main

import (
	"log"

	"github.com/nirmalkatiyar/api"
)

var PORT = ":8080"

// main: entry point of the code
func main() {

	// Start the cleanup LRU goroutine to remove expired nodes
	go api.LruCache.CleanupExpiredItems()

	// Start the WebSocket broadcaster
	go api.BroadcastCacheState()
	// API end-point router
	app := api.SetupRouter()

	// Start the application here
	err := app.Run(PORT)

	if err != nil {
		log.Println("Server is not running at PORT ", PORT)
	}
}
