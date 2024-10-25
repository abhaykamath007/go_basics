package database

import (
	"log"

	"github.com/abhaykamath_007/library-management-system/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.GetDBConnectionString()), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")
}
