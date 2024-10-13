package routes

import (
	"github.com/gin-gonic/gin"
	"meal-server/controllers"
)

func CartRoutes(r *gin.RouterGroup) {
	RegisterRestRoutes(r, "cart", controllers.NewCartController())
}
