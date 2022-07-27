package main

import (
	"github.com/luigiescalante/proyect-template/infrastructure/repository/database"
	"github.com/luigiescalante/proyect-template/infrastructure/repository/httpd"
)

func main() {
	//DATABASE PROCESS
	database.InitDB()
	database.Migrate()
	httpd.StartApiServer()
}
