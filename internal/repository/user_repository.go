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

func GetUser(userId uint) (model.User, error) {
	db := database.GetDB()
	var user model.User
	result := db.First(&user, userId)
	return user, result.Error
}

func UpdateUser(user *model.User) error {
	db := database.GetDB()
	return db.Save(user).Error
}

func DeleteUser(userId uint) error {
	db := database.GetDB()
	return db.Delete(&model.User{}, userId).Error
}
