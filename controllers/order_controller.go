package controllers

import (
	"net/http"
	"retailer-service/config"
	"retailer-service/models"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if product exists
	var product models.Product
	if err := config.DB.First(&product, "id = ?", order.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Check if stock is available
	if product.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Deduct stock
	product.Quantity -= order.Quantity
	config.DB.Save(&product)

	order.Status = "order placed"
	config.DB.Create(&order)

	c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := config.DB.First(&order, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
