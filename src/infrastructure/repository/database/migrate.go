package database

import (
	"github.com/luigiescalante/proyect-template/core/domain"
	"log"
)

func Migrate() {
	err := db.AutoMigrate(
		&domain.Users{},
		&domain.Roles{},
	)
	if err != nil {
		log.Panic(err.Error())
		return
	}
}
