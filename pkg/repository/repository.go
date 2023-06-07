package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
}

type Work interface {
	Create(userId int, work model.WorkInput) (int, error)
	GetAll() ([]model.Work, error)
}

type Request interface {
	CreateRequest(request model.RequestInput) (int, error)
	GetRequestsByStudentId(studentId int) ([]model.Request, error)
	IsRequestExists(workId int, studentId int) bool
	GetRequestsByWorkId(workId int) ([]model.Request, error)
	ChangeStatus(id int, statusId int) (string, error)
}

type User interface {
	GetStudentByUserId(userId int) (model.Student, error)
	GetInstructorByUserID(userId int) (model.Instructor, error)
}

type Repository struct {
	Authorization
	Work
	Request
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Work:          NewWorkPostgres(db),
		Request:       NewRequestPostgres(db),
		User:          NewUserPostgres(db),
	}
}
