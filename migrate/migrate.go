package main

import (
	"github.com/caspec1/agrak-test/config"
	"github.com/caspec1/agrak-test/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

// Migrate the models to DB
func main() {
	config.DB.AutoMigrate(&models.Product{})
}
