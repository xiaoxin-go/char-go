package models

type TDish struct {
	CommonModel
	Name        string  `gorm:"size:20;not null" json:"name"`
	Description string  `gorm:"size:255;not null" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Category    string  `gorm:"size:20;not null" json:"category"`
}
