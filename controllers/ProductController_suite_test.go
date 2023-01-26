package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caspec1/agrak-test/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test response code of the endpoint
func TestNewProduct(t *testing.T) {
	mockResponse := `{"msg":"Created successfully"}`

	r := gin.Default()

	r.POST("/api/products", NewProductMocked)
	product := models.Product{
		SKU:         "99999999",
		Name:        "Zapatillas Mujer",
		Brand:       "Adidas",
		Size:        "38",
		Price:       54500.00,
		ImageUrl:    "http://localhost:3000",
		OtherImages: []string{"http://localhost:3000", "http://localhost:8080"},
	}

	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusCreated, w.Code)
}

// Test response code of the endpoint
func TestUpdateProduct(t *testing.T) {
	mockResponse := `{"msg":"Updated Successfully"}`

	r := gin.Default()

	r.PUT("/api/products/:sku", UpdateProductMocked)

	product := models.Product{
		SKU:         "99999999",
		Name:        "Zapatillas Mujer",
		Brand:       "Adidas",
		Size:        "38",
		Price:       54500.00,
		ImageUrl:    "http://localhost:3000",
		OtherImages: []string{"http://localhost:3000", "http://localhost:8080"},
	}

	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("PUT", "/api/products/FAL-"+product.SKU, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test response code of the endpoint
func TestGetProducts(t *testing.T) {

	r := gin.Default()
	r.GET("/api/products")
	req, _ := http.NewRequest("GET", "/api/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Test response code of the endpoint
func TestGetProductBySKU(t *testing.T) {

	r := gin.Default()

	r.GET("/api/products/:sku", GetProductBySKUMocked)

	product := models.Product{
		SKU:         "99999999",
		Name:        "Zapatillas Mujer",
		Brand:       "Adidas",
		Size:        "38",
		Price:       54500.00,
		ImageUrl:    "http://localhost:3000",
		OtherImages: []string{"http://localhost:3000", "http://localhost:8080"},
	}

	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("GET", "/api/products/FAL-"+product.SKU, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Test response code of the endpoint
func TestDeleteProduct(t *testing.T) {

	r := gin.Default()

	r.DELETE("/api/products/:sku", GetProductBySKUMocked)

	product := models.Product{
		SKU:         "99999999",
		Name:        "Zapatillas Mujer",
		Brand:       "Adidas",
		Size:        "38",
		Price:       54500.00,
		ImageUrl:    "http://localhost:3000",
		OtherImages: []string{"http://localhost:3000", "http://localhost:8080"},
	}

	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("DELETE", "/api/products/FAL-"+product.SKU, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
