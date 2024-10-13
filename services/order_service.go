package services

import (
	"fmt"
	"math/rand"
	"meal-server/database"
	"meal-server/models"
	"time"
)

// CreateOrder 创建订单
func CreateOrder(tableId string) error {
	// 1. 获取桌号
	table := models.TTable{}
	if e := table.FirstByUuid(tableId); e != nil {
		return e
	}
	// 2. 获取购物车商品
	carts, e := new(models.TCart).FindByTableId(table.Id)
	if e != nil {
		return e
	}
	if len(carts) == 0 {
		return fmt.Errorf("购物车为空，下单失败")
	}
	// 3. 生成订单
	sum := 0.0
	for _, cart := range carts {
		sum += cart.Price
	}
	tx := database.DB.Begin()
	order := models.TOrder{
		TableId: table.Id,
		Count:   len(carts),
		Price:   sum,
		Sn:      newSn(table.Id),
	}
	if e := order.Create(tx); e != nil {
		tx.Rollback()
		return e
	}
	orderDish := make([]*models.TOrderDish, 0)
	for _, cart := range carts {
		orderDish = append(orderDish, &models.TOrderDish{
			OrderId: order.Id,
			DishId:  cart.DishId,
			Price:   cart.Price,
			Note:    cart.Note,
		})
	}
	if e := tx.Create(&orderDish).Error; e != nil {
		tx.Rollback()
		return fmt.Errorf("生成订单失败, err: %w", e)
	}
	// 4. 清空购物车
	if e := new(models.TCart).DeleteByTableId(tx, table.Id); e != nil {
		tx.Rollback()
		return e
	}
	if e := tx.Commit().Error; e != nil {
		tx.Rollback()
		return fmt.Errorf("保存订单信息失败, err: %w", e)
	}
	return nil
}

func newSn(tableId int) string {
	now := time.Now()
	r := rand.New(rand.NewSource(now.UnixNano()))
	date := now.Format("20060102150405")
	return fmt.Sprintf("%d-%s%d", tableId, date, r.Intn(1000))
}
