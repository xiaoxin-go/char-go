package migrate

import (
	"fmt"
	"meal-server/database"
	"meal-server/models"
)

func AutoMigrate() {
	if e := database.DB.AutoMigrate(
		models.TTable{},
		models.TCart{},
		models.TDish{},
		models.TCartDish{},
		models.TOrder{},
	); e != nil {
		fmt.Println(e.Error())
	}
}
