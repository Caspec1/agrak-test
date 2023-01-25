package routes

import (
	"github.com/caspec1/agrak-test/controllers"
	"github.com/gin-gonic/gin"
)

// Contain the routes for products
func ProductRoutes(r *gin.Engine) {
	r.GET("/api/products", controllers.GetProducts)
	r.POST("/api/products", controllers.NewProduct)
	r.GET("/api/products/:sku", controllers.GetProductBySKU)
	r.PUT("/api/products/:sku", controllers.UpdateProduct)
	r.DELETE("/api/products/:sku", controllers.DeleteProduct)
}
