package models

import "gorm.io/gorm"

type TTable struct {
	gorm.Model
	Number int    `gorm:"unique;not null" json:"number"` // 桌号
	Uuid   string `gorm:"not null;unique" json:"uuid"`   // 添加桌号时随机生成的，后面可以根据这个生成二维码
}
