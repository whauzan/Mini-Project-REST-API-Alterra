package users

import (
	"miniproject/business/users"
	"time"
)

type Users struct {
	ID             int    `gorm:"primaryKey"`
	Username       string `gorm:"unique"`
	FirstName      string
	LastName       string
	Email          string `gorm:"unique"`
	Password       string
	DOB            string
	Gender         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func toDomain(user Users) users.Domain {
	return users.Domain{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		DOB:       user.DOB,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		Username:  domain.Username,
		FirstName: domain.FirstName,
		LastName:  domain.LastName,
		Email:     domain.Email,
		Password:  domain.Password,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
