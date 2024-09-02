package models

import "gorm.io/gorm"

type TOrder struct {
	gorm.Model
	UserId int `gorm:"type:int(11);not null" json:"user_id"`
	DishId int `gorm:"type:int(11);not null" json:"dish_id"`
}
