package repository

import (
	"gorm.io/gorm"
	userDomain "miniproject/business/users"
	userDB "miniproject/repository/database/users"
	foodDomain "miniproject/business/food"
	foodDB "miniproject/repository/database/food"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewRepositoryFood(conn *gorm.DB) foodDomain.Repository {
	return foodDB.NewRepositoryMySQL(conn)
}
