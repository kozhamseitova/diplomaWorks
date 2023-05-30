package service

import (
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"github.com/kozhamseitova/diplomaWorks/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
