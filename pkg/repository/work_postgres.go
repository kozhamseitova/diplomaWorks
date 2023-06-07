package repository

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/jmoiron/sqlx"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"time"
)

type WorkPostgres struct {
	db *sqlx.DB
}

func NewWorkPostgres(db *sqlx.DB) *WorkPostgres {
	return &WorkPostgres{db: db}
}

func (r *WorkPostgres) Create(userId int, work model.WorkInput) (int, error) {
	//tx, err := r.db.Begin()
	//if err != nil {
	//	return 0, err
	//}
	//
	//var id int
	//createWorkQuery := fmt.Sprintf("INSERT INTO %s (created_at, title, description, type_id, degree_id, instructor_id, student_id, is_approved, request_numbers, progress_id) values ($1, $2)", worksTable)
	//_, err = tx.Exec(createWorkQuery, userId, id)
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}
	//
	//createWorkQuery2 := fmt.Sprintf("INSERT INTO %s (created_at, title, description, type_id, degree_id, instructor_id, student_id, is_approved, request_numbers, progress_id) values ($1, $2)", worksTable)
	//_, err = tx.Exec(createWorkQuery2, userId, id)
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}
	//
	//return id, tx.Commit()

	var id int
	query := fmt.Sprintf("INSERT INTO %s (created_at, title, description, type_id, degree_id, instructor_id, student_id, is_approved, progress_id) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", worksTable)

	row := r.db.QueryRow(query, time.Now(), work.Title, work.Description, work.TypeId, work.DegreeId, work.InstructorId, work.StudentId, work.IsApproved, work.ProgressId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil

}

//func (r *WorkPostgres) GetAllById(userId int) ([]model.Work, error) {
//	var works []model.Work
//	query := fmt.Sprintf("Select * from %s ")
//}

func (r *WorkPostgres) GetAll() ([]model.Work, error) {
	//var works []model.Work
	//query := fmt.Sprintf("Select * from %s ")
	//err := r.db.Select(&works, query)
	//
	//return works, err

	var works []model.Work
	query := fmt.Sprintf("SELECT w.id, w.title, w.description, w.is_approved,\n       t.id AS \"type.id\", t.name AS \"type.name\",\n       d.id AS \"degree.id\", d.name AS \"degree.name\",\n       u1.first_name AS \"instructor.first_name\", u1.last_name AS \"instructor.last_name\", i.about AS \"instructor.about\",\n       jsonb_agg(jsonb_build_object('id', ep.id, 'name', ep.name)) AS \"ep\",\n       jsonb_agg(DISTINCT jsonb_build_object('id', f.id, 'name', f.name)) AS \"faculty\"\nFROM works w\n         INNER JOIN types t ON t.id = w.type_id\n         INNER JOIN degrees d ON d.id = w.degree_id\n         INNER JOIN instructors i ON i.id = w.instructor_id\n         INNER JOIN users u1 ON u1.id = i.user_id\n         INNER JOIN works_eps we ON we.work_id = w.id\n         INNER JOIN ep ON ep.id = we.ep_id\n         INNER JOIN faculties f ON f.id = ep.faculty_id\nWHERE w.student_id IS NULL\nGROUP BY w.id, w.created_at, t.id, t.name, d.id, d.name, u1.first_name, u1.last_name, i.about;\n")
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var work model.Work
		var epData []byte
		var facultyData []byte
		err = rows.Scan(
			&work.Id,
			&work.Title,
			&work.Description,
			&work.IsApproved,
			&work.Type.Id,
			&work.Type.Name,
			&work.Degree.Id,
			&work.Degree.Name,
			&work.Instructor.FirstName,
			&work.Instructor.LastName,
			&work.Instructor.About,
			&epData,
			&facultyData,
		)
		if err != nil {
			return nil, err
		}
		var eps []*model.EP
		err = json.Unmarshal(epData, &eps)
		if err != nil {
			return nil, err
		}
		work.EP = make([]model.EP, len(eps))
		for i, ep := range eps {
			work.EP[i] = *ep
		}

		var faculties []*model.Faculty
		err = json.Unmarshal(facultyData, &faculties)
		if err != nil {
			return nil, err
		}
		work.Faculty = make([]model.Faculty, len(faculties))
		for i, faculty := range faculties {
			work.Faculty[i] = *faculty
		}

		works = append(works, work)
	}
	return works, nil
}
