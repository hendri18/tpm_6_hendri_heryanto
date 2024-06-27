package main

import (
	"tpm_6_HendriHeryanto/models"
	"tpm_6_HendriHeryanto/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var PORT = ":8080"
	// passwordnya saya sensor
	dsn := "host=localhost user=postgres password=****** dbname=tpm_6 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// melakukan migrasi / DDL
	// membuat table user
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err)
	}

	routers.SetupRouter(db).Run(PORT)
}
