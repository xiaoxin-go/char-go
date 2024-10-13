package models

import (
	"fmt"
	"github.com/google/uuid"
	"meal-server/database"
	"time"
)

type TTable struct {
	CommonModel
	Number         int       `gorm:"unique;not null" json:"number" binding:"required"` // 桌号
	Uuid           string    `gorm:"not null;unique" json:"uuid"`                      // 添加桌号时随机生成的，后面可以根据这个生成二维码
	LastActiveTime time.Time `json:"lastActiveTime"`
}

func (t *TTable) FirstByUuid(uuid string) error {
	if e := database.DB.Where("uuid = ?", uuid).First(t).Error; e != nil {
		return fmt.Errorf("获取桌子信息失败, uuid: %s, err: %w", uuid, e)
	}
	return nil
}

func (t *TTable) BeforeCreate() {
	t.Uuid = uuid.New().String()
}

func (t *TTable) TableName() string {
	return "t_table"
}
