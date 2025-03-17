package models

import "gorm.io/gorm"

type Order struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"` // Use uint for ID
	CustomerID string `json:"customer_id"`
	ProductID  uint   `json:"product_id"` // Use uint since product ID is now uint
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}

func MigrateOrders(db *gorm.DB) {
	db.AutoMigrate(&Order{})
}
