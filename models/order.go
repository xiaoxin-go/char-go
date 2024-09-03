package models

type TOrder struct {
	CommonModel
	UserId int `gorm:"type:int(11);not null" json:"user_id"`
	DishId int `gorm:"type:int(11);not null" json:"dish_id"`
}
