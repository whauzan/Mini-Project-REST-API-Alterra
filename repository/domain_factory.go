package repository

import (
	"gorm.io/gorm"
	userDomain "miniproject/business/users"
	userDB "miniproject/repository/database/users"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}
