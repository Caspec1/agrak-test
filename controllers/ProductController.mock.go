package controllers

import (
	"net/http"

	"github.com/caspec1/agrak-test/models"
	"github.com/caspec1/agrak-test/validations"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Mock newProduct function without db insertion
func NewProductMocked(c *gin.Context) {
	var product models.Product

	// Get info from json
	err := c.ShouldBindJSON(&product)

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

	c.JSON(http.StatusCreated, gin.H{
		"msg": "Created successfully",
	})
}

// Mock updateProduct function without db insertion
func UpdateProductMocked(c *gin.Context) {
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
	err := c.ShouldBindJSON(&body)

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

	// Returns a message to the client
	c.JSON(http.StatusOK, gin.H{
		"msg": "Updated Successfully",
	})
}

// Mock getProducts
func GetProductsMocked(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "Products",
	})
}

// Mock getProductsBySku
func GetProductBySKUMocked(c *gin.Context) {

	// Return product
	c.JSON(http.StatusOK, gin.H{
		"msg": "Product",
	})
}

// Mock deleteProducts
func DeleteProductMocked(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Delete Successfully",
	})
}
