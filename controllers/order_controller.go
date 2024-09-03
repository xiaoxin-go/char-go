package controllers

import (
	"meal-server/libs"
	"meal-server/models"
)

type OrderController struct {
	libs.Controller
}

func NewOrderController() libs.Restfuller {
	controller := &OrderController{}
	controller.ModelFunc = func() libs.Instance {
		return new(models.TOrder)
	}
	controller.ListFunc = func() any {
		return new([]*models.TOrder)
	}
	return controller
}
