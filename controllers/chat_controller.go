package controllers

import (
	"meal-server/libs"
	"meal-server/models"
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
