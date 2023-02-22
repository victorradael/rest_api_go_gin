package database

import (
	"log"

	"github.com/victorradael/rest_api_go_gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDatabase() {
	connectionConfig := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionConfig))
	if err != nil {
		log.Panic("Database connection erorr", err)
	}
	DB.AutoMigrate(&models.Student{})
}
