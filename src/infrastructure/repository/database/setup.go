package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func InitDB() {
	var err error
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	timezone := os.Getenv("DB_TIME_ZONE")
	dsn := "host=" + host + " port=" + port + " user=" + user + " dbname=" + database + " password=" + password + " TimeZone=" + timezone
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
	}
	db = dbConn
}

func Engine() *gorm.DB {
	return db
}
