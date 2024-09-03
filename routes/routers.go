package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meal-server/libs"
	"strings"
)

func SetupRoutes(r *gin.Engine) {
	group := r.Group("/api/v1")

	// 可以在这里添加其他路由
	TableRoutes(group)

}

// RegisterRestRoutes 注意restful路由
func RegisterRestRoutes(r *gin.RouterGroup, rest libs.Restfuller) {
	fmt.Println("----->", r.BasePath())
	parts := strings.Split(r.BasePath(), "/")
	path := parts[len(parts)-1]
	r.GET(fmt.Sprintf("%ss", path), rest.List)
	r.GET(path, rest.Get)
	r.POST(path, rest.Create)
	r.PUT(path, rest.Update)
	r.DELETE(path, rest.Delete)
}
