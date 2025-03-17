package models

import "gorm.io/gorm"

type Product struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`  
	Name     string  `gorm:"not null" json:"product_name"`
	Price    float64 `gorm:"not null;default:0" json:"price"`
	Quantity int     `gorm:"not null;default:0" json:"quantity"`
}

func MigrateProducts(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}
