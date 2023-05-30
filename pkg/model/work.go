package model

type Work struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	TypeId         int    `json:"type_id"`
	DegreeId       int    `json:"degree_id"`
	FacultyId      int    `json:"faculty_id"`
	InstructorId   int    `json:"instructor_id"`
	IsApproved     bool   `json:"is_approved"`
	RequestNumbers int    `json:"request_numbers"`
}

type Request struct {
	Id           int    `json:"id"`
	WorkId       int    `json:"work_id"`
	InstructorId int    `json:"instructor_id"`
	StudentId    int    `json:"student_id"`
	StatusId     int    `json:"status_id"`
	Description  string `json:"description"`
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

type EP struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	FacultyId int    `json:"faculty_id"`
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
