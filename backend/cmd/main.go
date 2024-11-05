package main

import (
	"log"

	"github.com/abhaykamath_007/library-management-system/backend/config"
	"github.com/abhaykamath_007/library-management-system/backend/database"
	"github.com/abhaykamath_007/library-management-system/backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	database.InitDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // Maximum age for preflight requests
	}))

	routes.SetUpRoutes(router)

	log.Printf("Starting server at %s", config.GetServerAddress())

	if err := router.Run(config.GetServerAddress()); err != nil {
		log.Fatalf("failed to start server : %v", err)
	}

	router.Run(":8080")
}
