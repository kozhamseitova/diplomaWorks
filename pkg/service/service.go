package service

import (
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"github.com/kozhamseitova/diplomaWorks/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Work interface {
	Create(userId int, work model.WorkInput) (int, error)
	GetAll() ([]model.Work, error)
	GetWorkById(id int) (model.Work, error)
	GetAllWorksForAdmin() ([]model.Work, error)
	GetWorksByInstructorId(instructorId int) ([]model.WorkInstructor, error)
	UpdateWork(userId int, id int, work model.WorkUpdate) error
	DeleteWork(id int) error
}

type Request interface {
	CreateRequest(request model.RequestInput) (int, error)
	GetRequestsByStudentId(studentId int) ([]model.Request, error)
	GetRequestsByWorkId(workId int) ([]model.Request, error)
	ChangeStatus(request model.RequestStatus) error
	DeleteRequest(id int) error
}

type User interface {
	GetStudentByUserId(userId int) (model.Student, error)
	GetInstructorByUserID(userId int) (model.Instructor, error)
}

type Service struct {
	Authorization
	Work
	Request
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Work:          NewWorkService(repos.Work),
		Request:       NewRequestService(repos.Request),
		User:          NewUserService(repos.User),
	}
}
