package service

import (
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"github.com/kozhamseitova/diplomaWorks/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u UserService) GetStudentByUserId(userId int) (model.Student, error) {
	return u.repo.GetStudentByUserId(userId)
}

func (u UserService) GetInstructorByUserID(userId int) (model.Instructor, error) {
	return u.repo.GetInstructorByUserID(userId)
}
