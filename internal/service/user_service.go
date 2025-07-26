package service

import (
	"go-rest/internal/model"
	"go-rest/internal/repository"
	"go-rest/pkg/config"
	"go-rest/pkg/util"
	"log"
	"time"
)

func CreateUser(user *model.User) error {
	encryptedPassword, err := encryptPassword(user.Password)
	if err != nil {
		log.Printf("Failed to encrypt password: %v", err)
		return err
	}
	user.Password = encryptedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return repository.CreateUser(user)
}

func GetUser(userId uint) (model.User, error) {
	return repository.GetUser(userId)
}

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}

func encryptPassword(password string) (string, error) {
	config := config.GetConfig()
	encryptionKey := config.EncryptionKey
	return util.Encrypt(encryptionKey, password)
}

func UpdateUser(user *model.User) error {
	user.UpdatedAt = time.Now()
	return repository.UpdateUser(user)
}

func DeleteUser(userId uint) error {
	return repository.DeleteUser(userId)
}
