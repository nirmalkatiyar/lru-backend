package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// SetupRouter configures the routes for the API
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},  // Allow all origins
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	router.GET("/cache", GetCacheItem)
	router.POST("/cache", SetCacheItem)
	router.DELETE("/cache/:key", DeleteCacheItem)
	router.GET("/ws", WebSocketEndpoint)
	return router
}
