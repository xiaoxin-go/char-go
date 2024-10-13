package models

import "time"

type CommonModel struct {
	Id        int       `gorm:"primary key" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *CommonModel) GetId() int {
	return c.Id
}
