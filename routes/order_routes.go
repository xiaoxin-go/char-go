package routes

import (
	"github.com/gin-gonic/gin"
	"meal-server/controllers"
)

func OrderRoutes(r *gin.RouterGroup) {
	RegisterRestRoutes(r, "order", controllers.NewOrderController())
}
