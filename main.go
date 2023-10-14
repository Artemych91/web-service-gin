package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Artemych91/web-service-gin/database"
	"github.com/Artemych91/web-service-gin/handlers"
)

func main() {
	db, error := database.InitDB()
	if error != nil {
		return
	}
	defer db.Close()

	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Routers
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbum)
	router.POST("/albums", handlers.PostAlbum)

	router.Run("localhost:8080")
}
