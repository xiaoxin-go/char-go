package controllers

import (
	"github.com/gin-gonic/gin"
	"meal-server/libs"
	"meal-server/models"
)

type TableController struct {
	libs.Controller
}

func NewTableController() libs.Restfuller {
	controller := &TableController{}
	controller.ModelFunc = func() libs.Instance {
		return new(models.TTable)
	}
	controller.ListFunc = func() any {
		return new([]*models.TTable)
	}
	return controller
}

func (c *TableController) Update(ctx *gin.Context) {
	libs.HttpServerError(ctx, "不支持更新")
	return
}
