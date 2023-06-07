package service

import (
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"github.com/kozhamseitova/diplomaWorks/pkg/repository"
)

type WorkService struct {
	repo repository.Work
}

func NewWorkService(repo repository.Work) *WorkService {
	return &WorkService{repo: repo}
}

func (s *WorkService) Create(userId int, work model.WorkInput) (int, error) {
	return s.repo.Create(userId, work)
}

func (s *WorkService) GetAll() ([]model.Work, error) {
	return s.repo.GetAll()
}
