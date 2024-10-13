package routes

import (
	"github.com/gin-gonic/gin"
	"meal-server/controllers"
)

func TableRoutes(r *gin.RouterGroup) {
	RegisterRestRoutes(r, "table", controllers.NewTableController())
}
