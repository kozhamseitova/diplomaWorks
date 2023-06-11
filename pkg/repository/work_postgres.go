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
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (created_at, title, description, type_id, degree_id, instructor_id, student_id, is_approved, progress_id) values ($1, $2, $3, $4, $5, $6, null, $7, $8) RETURNING id", worksTable)

	row := tx.QueryRow(query, time.Now(), work.Title, work.Description, work.TypeId, work.DegreeId, work.InstructorId, work.IsApproved, work.ProgressId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, value := range work.WorkEp {
		worksEpsQuery := fmt.Sprintf("INSERT Into works_eps (work_id, ep_id) values ($1, $2)")
		_, err := tx.Exec(worksEpsQuery, id, value)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *WorkPostgres) GetAll() ([]model.Work, error) {

	var works []model.Work
	query := fmt.Sprintf("SELECT w.id, w.title, w.description, w.is_approved,\n       t.id AS \"type.id\", t.name AS \"type.name\",\n       d.id AS \"degree.id\", d.name AS \"degree.name\",\n       u1.first_name AS \"instructor.first_name\", u1.last_name AS \"instructor.last_name\", i.about AS \"instructor.about\",\n       jsonb_agg(jsonb_build_object('id', ep.id, 'name', ep.name)) AS \"ep\",\n       jsonb_agg(DISTINCT jsonb_build_object('id', f.id, 'name', f.name)) AS \"faculty\"\nFROM works w\n         INNER JOIN types t ON t.id = w.type_id\n         INNER JOIN degrees d ON d.id = w.degree_id\n         INNER JOIN instructors i ON i.id = w.instructor_id\n         INNER JOIN users u1 ON u1.id = i.user_id\n         INNER JOIN works_eps we ON we.work_id = w.id\n         INNER JOIN ep ON ep.id = we.ep_id\n         INNER JOIN faculties f ON f.id = ep.faculty_id\nWHERE w.student_id IS NULL and is_approved = true\nGROUP BY w.id, w.created_at, t.id, t.name, d.id, d.name, u1.first_name, u1.last_name, i.about;")
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

func (r *WorkPostgres) GetWorkById(id int) (model.Work, error) {
	//TODO implement me
	panic("implement me")
}

func (r *WorkPostgres) GetAllWorksForAdmin() ([]model.Work, error) {
	var works []model.Work
	query := fmt.Sprintf("SELECT w.id, w.title, w.description, w.is_approved,\n       t.id AS \"type.id\", t.name AS \"type.name\",\n       d.id AS \"degree.id\", d.name AS \"degree.name\",\n       u1.first_name AS \"instructor.first_name\", u1.last_name AS \"instructor.last_name\", i.about AS \"instructor.about\",\n       jsonb_agg(jsonb_build_object('id', ep.id, 'name', ep.name)) AS \"ep\",\n       jsonb_agg(DISTINCT jsonb_build_object('id', f.id, 'name', f.name)) AS \"faculty\"\nFROM works w\n         INNER JOIN types t ON t.id = w.type_id\n         INNER JOIN degrees d ON d.id = w.degree_id\n         INNER JOIN instructors i ON i.id = w.instructor_id\n         INNER JOIN users u1 ON u1.id = i.user_id\n         INNER JOIN works_eps we ON we.work_id = w.id\n         INNER JOIN ep ON ep.id = we.ep_id\n         INNER JOIN faculties f ON f.id = ep.faculty_id\nWHERE is_approved = false\nGROUP BY w.id, w.created_at, t.id, t.name, d.id, d.name, u1.first_name, u1.last_name, i.about;")
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

//func (r *WorkPostgres) ApproveWork(id int) (int, error) {
//	query := fmt.Sprintf("UPDATE works\nSET is_approved = false where id = $1 ;\n")
//	_, err := r.db.Exec(query, id)
//	if err != nil {
//		return 0, err
//	}
//	return -1, nil
//}

func (r *WorkPostgres) GetWorksByInstructorId(instructorId int) ([]model.WorkInstructor, error) {
	var works []model.WorkInstructor
	query := fmt.Sprintf("SELECT w.id, w.title, w.description, w.is_approved,\n       t.id AS \"type.id\", t.name AS \"type.name\",\n       d.id AS \"degree.id\", d.name AS \"degree.name\",\n       p.id AS \"progress.id\", p.name AS \"progress.name\",\n       jsonb_agg(jsonb_build_object('id', ep.id, 'name', ep.name)) AS \"ep\",\n       jsonb_agg(DISTINCT jsonb_build_object('id', f.id, 'name', f.name)) AS \"faculty\"\nFROM works w\n         INNER JOIN types t ON t.id = w.type_id\n         INNER JOIN degrees d ON d.id = w.degree_id\n         INNER JOIN instructors i ON i.id = w.instructor_id\n         INNER JOIN users u1 ON u1.id = i.user_id\n         INNER JOIN works_eps we ON we.work_id = w.id\n         INNER JOIN ep ON ep.id = we.ep_id\n         INNER JOIN faculties f ON f.id = ep.faculty_id\n         INNER JOIN progress p on p.id = w.progress_id\nWHERE i.id = $1\nGROUP BY w.id, w.created_at, t.id, t.name, d.id, d.name, u1.first_name, u1.last_name, i.about, p.id;\n")
	rows, err := r.db.Query(query, instructorId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var work model.WorkInstructor
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
			&work.Progress.Id,
			&work.Progress.Name,
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

func (r *WorkPostgres) UpdateWork(userId int, id int, workInput model.WorkUpdate) error {
	query := fmt.Sprintf("UPDATE works SET")
	if workInput.StudentId > 0 {
		query += fmt.Sprintf(" student_id = $1 WHERE id = $2")
		_, err := r.db.Exec(query, workInput.StudentId, id)
		if err != nil {
			return err
		}
	}
	if workInput.IsApproved {
		query += fmt.Sprintf(" is_approved = $1 WHERE id = $2")
		_, err := r.db.Exec(query, workInput.IsApproved, id)
		if err != nil {
			return err
		}
	}
	if workInput.ProgressId > 0 {
		query += fmt.Sprintf(" progress_id = $1 WHERE id = $2")
		_, err := r.db.Exec(query, workInput.ProgressId, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *WorkPostgres) DeleteWork(id int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	queryWorks := "DELETE FROM works WHERE id = $1;"
	queryWorksEP := "DELETE FROM works_eps WHERE work_id = $1;"

	_, err = tx.Exec(queryWorksEP, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(queryWorks, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
