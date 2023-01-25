package controllers

import (
	"net/http"

	"github.com/caspec1/agrak-test/config"
	"github.com/caspec1/agrak-test/models"
	"github.com/caspec1/agrak-test/validations"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Create a new product
func NewProduct(c *gin.Context) {
	var product models.Product

	// Get info from json
	err := c.ShouldBind(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Can't get information"})
		return
	}

	// Check that the information don't be blank
	if product.SKU == "" || product.Name == "" || product.Brand == "" || product.Size == "" || product.ImageUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "All fields required"})
		return
	}

	// Validate sku have 7 digits and join FAL at start
	sku, errValidate := validations.SkuValidation(product.SKU)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": errValidate.Error()})
		return
	}

	// Assignment a new sku that starts with FAL
	product.SKU = sku

	// Validate the name's text size
	errName := validations.SizeValidation(product.Name)
	if errName != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": errName.Error()})
		return
	}

	// Validate the brand's text size
	errBrand := validations.SizeValidation(product.Brand)
	if errBrand != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": errBrand.Error()})
		return
	}

	// Validate price
	errPrice := validations.PriceValidation(product.Price)
	if errPrice != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": errPrice.Error()})
		return
	}

	// Validate url
	errURL := validations.UrlValidation(product.ImageUrl)
	if errURL != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": errURL.Error()})
		return
	}

	// Validate url's array
	if len(product.OtherImages) > 0 {
		errArray := validations.OtherUrlValidation(product.OtherImages)
		if errArray != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errArray.Error()})
			return
		}
	}

	// Save in DB
	config.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{
		"product": product,
		"msg":     "Created successfully",
	})
}

// Update a product
func UpdateProduct(c *gin.Context) {
	sku := c.Param("sku")
	var body struct {
		SKU         string         `json:"sku"`
		Name        string         `json:"name"`
		Brand       string         `json:"brand"`
		Size        string         `json:"size"`
		Price       float64        `json:"price"`
		ImageUrl    string         `json:"image_url"`
		OtherImages pq.StringArray `json:"other_images"`
	}

	// Get info from json
	err := c.ShouldBind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Can't get information"})
		return
	}

	if body.SKU != "" {
		// Validate sku have 7 digits and join FAL at start
		sku, errValidate := validations.SkuValidation(body.SKU)

		if errValidate != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errValidate.Error()})
			return
		}
		body.SKU = sku
	}

	if body.Name != "" {
		// Validate the name's text size
		errName := validations.SizeValidation(body.Name)
		if errName != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errName.Error()})
			return
		}
	}

	if body.Brand != "" {
		// Validate the name's text size
		errBrand := validations.SizeValidation(body.Brand)
		if errBrand != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errBrand.Error()})
			return
		}
	}

	if body.Price != 0 {
		// Validate price
		errPrice := validations.PriceValidation(body.Price)
		if errPrice != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errPrice.Error()})
			return
		}
	}

	if body.ImageUrl != "" {
		// Validate url
		errURL := validations.UrlValidation(body.ImageUrl)
		if errURL != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errURL.Error()})
			return
		}
	}

	if len(body.OtherImages) > 0 {
		errArray := validations.OtherUrlValidation(body.OtherImages)
		if errArray != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errArray.Error()})
			return
		}
	}

	var product models.Product

	// Get product from DB
	config.DB.Take(&product, "sku = ?", sku)

	// Check if the product exists
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Producto no encontrado"})
		return
	}

	// Update product
	config.DB.Model(&product).Updates(models.Product{
		SKU:         body.SKU,
		Name:        body.Name,
		Brand:       body.Brand,
		Size:        body.Size,
		Price:       body.Price,
		ImageUrl:    body.ImageUrl,
		OtherImages: body.OtherImages,
	})

	// Returns a message to the client
	c.JSON(http.StatusOK, gin.H{
		"msg":     "Updated Successfully",
		"product": product,
	})

}

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product

	config.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

// Get product by sku
func GetProductBySKU(c *gin.Context) {
	sku := c.Param("sku")
	var product models.Product

	// Get product from DB
	config.DB.Take(&product, "sku = ?", sku)

	// Check if the product exists
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Producto no encontrado"})
		return
	}

	// Return product
	c.JSON(http.StatusOK, product)
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	sku := c.Param("sku")
	var product models.Product

	// Get product from DB
	config.DB.Take(&product, "sku = ?", sku)

	// Check if the product exists
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Producto no encontrado"})
		return
	}

	// Delete from DB
	config.DB.Delete(&product)

	// NOTE: if you want delete the complete data from DB run:
	// config.DB.Unscoped().Delete(&product)

	// Returns a message to the client
	c.JSON(http.StatusOK, gin.H{
		"msg":     "Delete Successfully",
		"product": product,
	})
}
