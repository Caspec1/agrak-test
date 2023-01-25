package main

import (
	"github.com/caspec1/agrak-test/config"
	"github.com/caspec1/agrak-test/routes"
	"github.com/gin-gonic/gin"
)

// Initialize env variables and connect DB
func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	// Config Router
	r := gin.Default()

	// Routes of Products
	routes.ProductRoutes(r)

	// Run server
	r.Run()
}
