package business

import "errors"

var (
	ErrInternalServer  = errors.New("Our service can't reach you :( Please contact our administrator")
	ErrNotFound        = errors.New("Data not found :(")
	ErrIDNotFound      = errors.New("ID not found :(")
	ErrDuplicateData   = errors.New("There is something duplicated")
	ErrUsrnameNotFound = errors.New("Ara, username not found")
	ErrEmailNotFound   = errors.New("Email not found")
	ErrPasswdNotFound  = errors.New("Ara ara, password not found")
	ErrwrongPasswd     = errors.New("Ara ara ara, password incorrect")
)
