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

func (s *WorkService) GetWorkById(id int) (model.Work, error) {
	return s.repo.GetWorkById(id)
}

func (s *WorkService) GetAllWorksForAdmin() ([]model.Work, error) {
	return s.repo.GetAllWorksForAdmin()
}

func (s *WorkService) GetWorksByInstructorId(instructorId int) ([]model.WorkInstructor, error) {
	return s.repo.GetWorksByInstructorId(instructorId)
}

func (s *WorkService) UpdateWork(userId int, id int, work model.WorkUpdate) error {
	return s.repo.UpdateWork(userId, id, work)
}

func (s *WorkService) DeleteWork(id int) error {
	return s.repo.DeleteWork(id)
}

func (s *WorkService) Create(userId int, work model.WorkInput) (int, error) {
	return s.repo.Create(userId, work)
}

func (s *WorkService) GetAll() ([]model.Work, error) {
	return s.repo.GetAll()
}
