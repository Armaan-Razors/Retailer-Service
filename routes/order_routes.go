package routes

import (
	"retailer-service/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.POST("/order", controllers.PlaceOrder)
	r.GET("/order/:id", controllers.GetOrder)
}
