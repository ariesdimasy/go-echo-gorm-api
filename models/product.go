package models

import "gorm.io/gorm"

type Product struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	gorm.Model
}
