package users

import (
	"miniproject/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(username string, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "username = ?", username)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return toDomain(user), nil
}

func (rep *MysqlUserRepository) Register(domain *users.Domain) (users.Domain, error) {
	user := fromDomain(*domain)
	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return toDomain(user), nil
}
