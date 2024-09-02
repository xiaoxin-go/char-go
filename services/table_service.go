package services

import (
	"fmt"
	"meal-server/database"
	"meal-server/models"
)

func AddTable(table *models.TTable) error {
	if e := database.DB.Create(table).Error; e != nil {
		return fmt.Errorf("添加桌号失败, err: %w", e)
	}
	return nil
}
