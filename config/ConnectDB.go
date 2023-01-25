package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect PostgreSQL DB
func ConnectDB() {
	var err error
	dsn := "host=" + os.Getenv("PGHOST") + " user=" + os.Getenv("PGUSER") + " password=" + os.Getenv("PGPASSWORD") + " dbname=" + os.Getenv("PGDATABASE") + " port=" + os.Getenv("PGPORT") + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect")
	}
}
