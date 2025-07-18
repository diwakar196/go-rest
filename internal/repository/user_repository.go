package repository

import (
	"go-rest/internal/model"
	"go-rest/pkg/database"
)

func CreateUser(user *model.User) error {
	db := database.GetDB()
	return db.Create(user).Error
}

func GetAllUsers() ([]model.User, error) {
	db := database.GetDB()
	var users []model.User
	result := db.Find(&users)
	return users, result.Error
}
