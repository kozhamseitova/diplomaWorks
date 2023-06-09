package service

import (
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"github.com/kozhamseitova/diplomaWorks/pkg/repository"
)

type RequestService struct {
	repo repository.Request
}

func NewRequestService(repo repository.Request) *RequestService {
	return &RequestService{repo: repo}
}

func (s *RequestService) DeleteRequest(id int) error {
	return s.repo.DeleteRequest(id)
}

func (s *RequestService) CreateRequest(request model.RequestInput) (int, error) {
	return s.repo.CreateRequest(request)
}

func (s *RequestService) GetRequestsByStudentId(studentId int) ([]model.Request, error) {
	return s.repo.GetRequestsByStudentId(studentId)
}

func (s *RequestService) GetRequestsByWorkId(workId int) ([]model.Request, error) {
	return s.repo.GetRequestsByWorkId(workId)
}

func (s *RequestService) ChangeStatus(request model.RequestStatus) error {
	return s.repo.ChangeStatus(request)
}
