package model

type User struct {
	Id          int    `json:"-" db:"id"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Student struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Degree      Degree `json:"degree"`
	EP          EP     `json:"ep"`
	GroupName   string `json:"group_name"`
}

//
//type Instructor struct {
//	Id          int    `json:"id"`
//	FirstName   string `json:"first_name"`
//	LastName    string `json:"last_name"`
//	PhoneNumber string `json:"phone_number"`
//	Role        string `json:"role"`
//	About       string `json:"about"`
//}
