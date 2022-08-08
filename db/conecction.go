package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=erickype password=erickype97. dbname=gorn port=5432"
var DB *gorm.DB

func DBConecction() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB coneccted")
	}
}
