package models

import "gorm.io/gorm"

type TUser struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);unique;not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
}
