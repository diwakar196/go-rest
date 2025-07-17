package service

import (
	"go-rest/internal/model"
	"go-rest/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *Service) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}
