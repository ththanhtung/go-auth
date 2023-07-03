package port

import "lab5.cmo/internal/application/core"

type DbPort interface {
	CreateUser(username, password, firstname, lastname, email, dob, avatar, address string) (core.User, error)
	UpdateUser(username, password string, updatecontent core.User) (core.User, error)
	GetUsers() ([]core.User, error)
	DeleteUser(username, password string) (core.User, error)
}
