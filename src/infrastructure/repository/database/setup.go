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

/*Test other rdms*/
/*func InitDB() {
	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/gorm_playground?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Print(err)
	}
	db = dbConn
}*/
