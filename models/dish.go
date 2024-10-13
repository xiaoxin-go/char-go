package models

import (
	"fmt"
	"meal-server/database"
)

type TDish struct {
	CommonModel
	Name        string  `gorm:"size:20;not null" json:"name"`
	Description string  `gorm:"size:255;not null" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Category    string  `gorm:"size:20;not null" json:"category"`
}

func (t *TDish) FirstById(id int) error {
	if e := database.DB.First(t, id).Error; e != nil {
		return fmt.Errorf("获取商品信息失败, err: %w", e)
	}
	return nil
}
