package models

import "gorm.io/gorm"

type Order struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"` 
	CustomerID string `json:"customer_id"`
	ProductID  uint   `json:"product_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}

func MigrateOrders(db *gorm.DB) {
	db.AutoMigrate(&Order{})
}
