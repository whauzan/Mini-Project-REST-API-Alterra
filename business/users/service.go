package users

import (
	"miniproject/app/middleware"
	"miniproject/business"
	"miniproject/helper/encrypt"
	"time"
)

type UserService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserService(repo Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) Service {
	return &UserService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

// Business logic for regist and login User
func (service *UserService) Register(domain *Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, business.ErrUsrnameNotFound
	}
	if domain.Email == "" {
		return Domain{}, business.ErrEmailNotFound
	}
	if domain.Password == "" {
		return Domain{}, business.ErrPasswdNotFound
	}
	encryptPass, _ := encrypt.HashAndPassword(domain.Password)
	domain.Password = encryptPass
	usr, err := service.repository.Register(domain)

	if err != nil {
		return Domain{}, err
	}
	return usr, nil
}

func (service *UserService) Login(usrname, passwd string) (Domain, error) {
	if usrname == "" {
		return Domain{}, business.ErrUsrnameNotFound
	}
	if passwd == "" {
		return Domain{}, business.ErrPasswdNotFound
	}

	usr, err := service.repository.Login(usrname, passwd)
	if err != nil {
		return Domain{}, err
	}
	valid := encrypt.CheckPasswordHash(passwd, usr.Password)
	if !valid {
		return Domain{}, business.ErrwrongPasswd
	}
	if err != nil {
		return Domain{}, err
	}
	usr.Token = service.jwtAuth.GenerateToken(usr.ID, "user")
	if err != nil {
		return Domain{}, err
	}
	return usr, nil
}
