package model

type WorkInput struct {
	Id           int    `json:"id" db:"id"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	Title        string `json:"title" db:"title" binding:"required"`
	Description  string `json:"description" db:"description"`
	TypeId       int    `json:"type_id" db:"type_id"`
	DegreeId     int    `json:"degree_id" db:"degree_id"`
	FacultyId    int    `json:"faculty_id" db:"faculty_id"`
	InstructorId int    `json:"instructor_id" db:"instructor_id"`
	StudentId    int    `json:"student_id" db:"student_id"`
	IsApproved   bool   `json:"is_approved" db:"is_approved"`
	ProgressId   int    `json:"progress_id" db:"progress_id"`
}

type Work struct {
	Id          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title" binding:"required"`
	Description string     `json:"description" db:"description"`
	Type        Type       `json:"type" db:"type"`
	Degree      Degree     `json:"degree" db:"degree"`
	EP          []EP       `json:"ep" db:"ep"`
	Faculty     []Faculty  `json:"faculty" db:"faculty"`
	Instructor  Instructor `json:"instructor" db:"instructor"`
	IsApproved  bool       `json:"is_approved" db:"is_approved"`
}

type Instructor struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//PhoneNumber string `json:"phone_number"`
	About string `json:"about"`
}

type Request struct {
	Id          int         `json:"id"`
	CreatedAt   string      `json:"created_at" db:"created_at"`
	WorkRequest WorkRequest `json:"work"`
	Instructor  Instructor  `json:"instructor"`
	Student     Student     `json:"student"`
	Status      Status      `json:"status"`
	Description string      `json:"description"`
}

type RequestInput struct {
	WorkId      int    `json:"work_id"`
	StudentId   int    `json:"student_id"`
	Description string `json:"description"`
}
type WorkRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Progress struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Degree struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Faculty struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Type struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type EP struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DiplomaFaculties struct {
	WorkId    int `json:"work_id"`
	FacultyId int `json:"faculty_id"`
}

type DiplomaFacultiesEP struct {
	WorkId    int `json:"work_id"`
	FacultyId int `json:"faculty_id"`
	EPId      int `json:"ep_id"`
}

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
