package models

import (
	"fmt"
	"gorm.io/gorm"
	"meal-server/database"
)

type TCart struct {
	CommonModel
	TableId int     `gorm:"not null" json:"table_id"` // 桌号
	DishId  int     `gorm:"not null" json:"dish_id"`  // 商品
	Price   float64 `json:"price"`                    // 价格
	Note    string  `gorm:"size:255" json:"note"`     // 备注
}

func (t *TCart) Create(tx *gorm.DB) error {
	if tx == nil {
		tx = database.DB
	}
	if err := tx.Create(t).Error; err != nil {
		return fmt.Errorf("添加购物车失败, err: %w", err)
	}
	return nil
}

func (t *TCart) FindByTableId(tableId int) ([]*TCart, error) {
	carts := make([]*TCart, 0)
	if e := database.DB.Where("table_id = ?", tableId).Find(&carts).Error; e != nil {
		return nil, fmt.Errorf("获取购物车信息失败, err: %w", e)
	}
	return carts, nil
}

func (t *TCart) DeleteByTableId(tx *gorm.DB, tableId int) error {
	if e := tx.Where("table_id = ?", tableId).Delete(&TCart{}).Error; e != nil {
		return fmt.Errorf("清空购物车失败, err: %w", e)
	}
	return nil
}
