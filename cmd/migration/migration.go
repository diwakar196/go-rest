package migration

import (
	"go-rest/internal/model"
	"go-rest/pkg/database"
)

func Migrate() {
	db := database.GetDB()
	db.AutoMigrate(&model.User{})
}
