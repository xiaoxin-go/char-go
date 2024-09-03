package models

import "time"

type CommonModel struct {
	ID        int       `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *CommonModel) GetId() int {
	return c.ID
}
