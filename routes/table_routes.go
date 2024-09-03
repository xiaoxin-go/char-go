package routes

import (
	"github.com/gin-gonic/gin"
	"meal-server/controllers"
)

func TableRoutes(r *gin.RouterGroup) {
	table := r.Group("table")
	RegisterRestRoutes(table, controllers.NewTableController())
}
