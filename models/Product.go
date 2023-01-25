package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	SKU         string         `gorm:"not null" json:"sku"`
	Name        string         `gorm:"not null" json:"name"`
	Brand       string         `gorm:"not null" json:"brand"`
	Size        string         `gorm:"not null" json:"size"`
	Price       float64        `json:"price"`
	ImageUrl    string         `json:"image_url"`
	OtherImages pq.StringArray `gorm:"type:text[]" json:"other_images"`
}
