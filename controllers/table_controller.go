package controllers

import (
	"meal-server/libs"
	"meal-server/models"
)

var TableController libs.Restfuller

func init() {
	controller := &libs.Controller{}
	controller.ModelFunc = func() libs.Instance {
		return new(models.TTable)
	}
	controller.ListFunc
	TableController = controller
}
