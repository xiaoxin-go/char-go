package controllers

import (
	"github.com/gin-gonic/gin"
	"meal-server/libs"
	"meal-server/models"
	"meal-server/services"
)

type CartController struct {
	libs.Controller
}

func NewCartController() libs.Restfuller {
	controller := &CartController{}
	controller.ModelFunc = func() libs.Instance {
		return new(models.TCart)
	}
	controller.ListFunc = func() any {
		return new([]*models.TCart)
	}
	return controller
}

// Create 添加购物车
func (c *CartController) Create(ctx *gin.Context) {
	req := services.AddCartReq{}
	if e := ctx.ShouldBindJSON(&req); e != nil {
		libs.HttpParamsError(ctx, "校验参数失败, err: %s", e.Error())
		return
	}
	if e := services.AddCart(&req); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "添加成功")
}

func (c *CartController) Update(ctx *gin.Context) {
	libs.HttpParamsError(ctx, "不支持更新")
}
