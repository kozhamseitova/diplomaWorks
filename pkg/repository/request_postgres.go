package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
)

type RequestPostgres struct {
	db *sqlx.DB
}

func NewRequestPostgres(db *sqlx.DB) *RequestPostgres {
	return &RequestPostgres{db: db}
}

func (r *RequestPostgres) IsRequestExists(workId int, studentId int) bool {
	var id int
	query := fmt.Sprintf("select count(id) from requests where work_id=$1 and student_id=$2")

	row := r.db.QueryRow(query, workId, studentId)
	if err := row.Scan(&id); err != nil {
		return false
	}

	if id != 0 {
		return false
	}

	return true
}

func (r *RequestPostgres) CreateRequest(request model.RequestInput) (int, error) {
	if r.IsRequestExists(request.WorkId, request.StudentId) {
		var id int
		query := fmt.Sprintf("INSERT INTO %s (created_at, work_id, student_id, status_id, description) values (now(), $1, $2, 1, $3) RETURNING id", requestTable)

		row := r.db.QueryRow(query, request.WorkId, request.StudentId, request.Description)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
		return id, nil
	}

	var err error
	err = errors.New("request is already exists")

	return 0, err
}

func (r *RequestPostgres) GetRequestsByStudentId(studentId int) ([]model.Request, error) {
	return nil, nil
}

func (r *RequestPostgres) GetRequestsByWorkId(workId int) ([]model.Request, error) {
	var requests []model.Request

	return requests, nil
}

func (r *RequestPostgres) ChangeStatus(id int, statusID int) (string, error) {
	query := fmt.Sprintf("UPDATE %s SET status_id = $1 WHERE id = $2", requestTable)

	_, err := r.db.Exec(query, statusID, id)
	if err != nil {
		return "", err
	}
	return "Success", nil
}
