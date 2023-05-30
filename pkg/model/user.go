package model

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

type Student struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	DegreeId   int    `json:"degree_id"`
	FacultyId  int    `json:"faculty_id"`
	Group      string `json:"group"`
	WorkTitle  string `json:"workTitle"`
	ProgressId int    `json:"progress_id"`
}

type Instructor struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	About  string `json:"about"`
}
