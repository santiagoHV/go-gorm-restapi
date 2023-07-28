package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=user password=password dbname=gofaztapi-db port=5433"
var DB *gorm.DB

func DBconnection() {
	var error error

	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}
}
