package controllers

import (
	"meal-server/libs"
	"meal-server/models"
)

type DishController struct {
	libs.Controller
}

func NewDishController() libs.Restfuller {
	controller := &DishController{}
	controller.ModelFunc = func() libs.Instance {
		return new(models.TDish)
	}
	controller.ListFunc = func() any {
		return new([]*models.TDish)
	}
	return controller
}
