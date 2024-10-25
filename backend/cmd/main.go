package main

import (
	"log"

	"github.com/abhaykamath_007/library-management-system/backend/config"
	"github.com/abhaykamath_007/library-management-system/backend/database"
	"github.com/abhaykamath_007/library-management-system/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	database.InitDB()

	router := gin.Default()

	routes.SetUpRoutes(router)

	log.Printf("Starting server at %s", config.GetServerAddress())

	if err := router.Run(config.GetServerAddress()); err != nil {
		log.Fatalf("failed to start server : %v", err)
	}

	router.Run(":8080")
}
