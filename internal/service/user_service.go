package service

import (
	"go-rest/internal/model"
	"go-rest/internal/repository"
)

func CreateUser(user *model.User) error {
	return repository.CreateUser(user)
}

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}
