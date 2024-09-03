package main

import (
	"github.com/gin-gonic/gin"
	"meal-server/config"
	"meal-server/database"
	"meal-server/routes"
)

func main() {
	r := gin.Default()

	// 加载配置
	config.LoadConfig("config.json")

	// 连接数据库
	database.ConnectDatabase()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务
	r.Run(":8080")
}
