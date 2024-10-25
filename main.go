package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"restapi/controllers"
	"restapi/routes"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "root:mysql123#@tcp(127.0.0.1:3306)/album_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Set the global db variable in controllers
	controllers.SetDB(db)

	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run("localhost:8080")
}
