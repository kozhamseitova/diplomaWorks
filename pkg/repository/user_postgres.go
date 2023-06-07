package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u UserPostgres) GetStudentByUserId(userId int) (model.Student, error) {
	var student model.Student
	query := fmt.Sprintf("SELECT s.id, s.group_name, d.id, d.name, ep.id, ep.name, u.first_name, u.last_name, u.phone_number " +
		"FROM students s " +
		"INNER JOIN degrees d ON d.id = s.degree_id " +
		"INNER JOIN ep ep ON ep.id = s.ep_id " +
		"INNER JOIN users u ON u.id = s.user_id " +
		"WHERE u.id = $1")

	err := u.db.QueryRow(query, userId).Scan(
		&student.Id,
		&student.GroupName,
		&student.Degree.Id,
		&student.Degree.Name,
		&student.EP.Id,
		&student.EP.Name,
		&student.FirstName,
		&student.LastName,
		&student.PhoneNumber,
	)

	return student, err
}

func (u UserPostgres) GetInstructorByUserID(userId int) (model.Instructor, error) {
	//TODO implement me
	panic("implement me")
}
