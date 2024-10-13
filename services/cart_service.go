package services

import "meal-server/models"

type AddCartReq struct {
	TableId string `json:"table_id" binding:"required"`
	DishId  int    `json:"dish_id" binding:"required"`
	Note    string `json:"note"`
}

func AddCart(req *AddCartReq) error {
	// 1. 获取桌号，判断桌号是否存在
	table := models.TTable{}
	if e := table.FirstByUuid(req.TableId); e != nil {
		return e
	}
	// 2. 获取商品信息
	dish := models.TDish{}
	if e := dish.FirstById(req.DishId); e != nil {
		return e
	}
	// 2. 将商品加入购物车中
	cart := models.TCart{
		TableId: table.Id,
		DishId:  req.DishId,
		Note:    req.Note,
		Price:   dish.Price,
	}
	if e := cart.Create(nil); e != nil {
		return e
	}
	return nil
}
