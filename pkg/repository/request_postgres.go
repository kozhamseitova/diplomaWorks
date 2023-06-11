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

func (r *RequestPostgres) DeleteRequest(id int) error {
	query := fmt.Sprintf("DELETE FROM requests where id = $1")

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
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
	var requests []model.Request
	query := fmt.Sprintf("select r.id, r.created_at, w.id, w.title, w.description, s.id, s.name, r.description from requests r inner join works w on w.id = r.work_id inner join statuses s on s.id = r.status_id where r.student_id = $1")
	rows, err := r.db.Query(query, studentId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var request model.Request
		err = rows.Scan(
			&request.Id,
			&request.CreatedAt,
			&request.WorkRequest.Id,
			&request.WorkRequest.Title,
			&request.WorkRequest.Description,
			&request.Status.Id,
			&request.Status.Name,
			&request.Description,
		)

		requests = append(requests, request)
	}
	return requests, nil
}

func (r *RequestPostgres) GetRequestsByWorkId(workId int) ([]model.Request, error) {
	var requests []model.Request
	query := fmt.Sprintf("select r.id, r.created_at, s.id, s.name, r.description, u.id, u.first_name, u.last_name from requests r inner join statuses s on s.id = r.status_id inner join students s2 on r.student_id = s2.id inner join users u on s2.user_id = u.id where r.work_id = $1")
	rows, err := r.db.Query(query, workId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var request model.Request
		err = rows.Scan(
			&request.Id,
			&request.CreatedAt,
			&request.Student.Id,
			&request.Student.FirstName,
			&request.Student.LastName,
			&request.Student.GroupName,
			&request.Student.Degree.Id,
			&request.Student.Degree.Name,
			&request.Student.EP.Id,
			&request.Student.EP.Name,
			&request.Status.Id,
			&request.Status.Name,
			&request.Description,
		)

		requests = append(requests, request)
	}
	return requests, nil
}

func (r *RequestPostgres) ChangeStatus(request model.RequestStatus) error {
	query := fmt.Sprintf("UPDATE %s SET status_id = $1 WHERE id = $2", requestTable)

	_, err := r.db.Exec(query, request.StatusId, request.Id)
	if err != nil {
		return err
	}
	return nil
}
