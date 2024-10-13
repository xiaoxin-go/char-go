package models

import (
	"fmt"
	"gorm.io/gorm"
	"meal-server/database"
)

type TOrder struct {
	CommonModel
	UserId  int     `gorm:"type:int(11);not null" json:"user_id"`    // 下单用户
	Sn      string  `gorm:"size:50;not null;unique" json:"order_id"` // 订单号
	TableId int     `gorm:"not null" json:"table_id"`                // 桌子
	Count   int     `json:"count"`                                   // 商品数量
	Price   float64 `json:"price"`                                   // 价格
}

func (t *TOrder) Create(tx *gorm.DB) error {
	if tx == nil {
		tx = database.DB
	}
	if e := tx.Create(t).Error; e != nil {
		return fmt.Errorf("订单生成失败, err: %w", e)
	}
	return nil
}

// TOrderDish 订单商品表
type TOrderDish struct {
	CommonModel
	OrderId int     `gorm:"not null" json:"order_id"`
	DishId  int     `gorm:"not null" json:"dish_id"`
	Price   float64 `json:"price"`
	Note    string  `gorm:"size:500" json:"note"`
}
