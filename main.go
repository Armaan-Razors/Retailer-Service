package main

import (
	"log"
	"retailer-service/config"
	"retailer-service/models"
	"retailer-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	models.MigrateProducts(config.DB)
	models.MigrateOrders(config.DB)

	r := gin.Default()

	routes.ProductRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
