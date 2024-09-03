package models

type TCart struct {
	CommonModel
	TableId int `gorm:"not null" json:"table_id"`
}

type TCartDish struct {
	Id     int    `gorm:"primary_key;AUTO_INCREMENT"`
	CartId int    `gorm:"not null" json:"cart_id"`
	DishId int    `gorm:"not null" json:"dish_id"`
	Count  int    `gorm:"default:1" json:"count"`
	Desc   string `gorm:"type:text" json:"desc"`
}
