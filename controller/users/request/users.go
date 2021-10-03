package request

import (
	"miniproject/business/users"
)

// request body for login
type UsersLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// request body for register
type Users struct {
	Username   string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	DOB        string `json:"dob"`
	Gender     string `json:"gender"`
}

// turn register body to domain
func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Username:   req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:      req.Email,
		Password:   req.Password,
		DOB:        req.DOB,
		Gender:     req.Gender,
	}
}
