package port

import (
	"lab5.cmo/internal/adapters/framework/right/db"
)

type DbPort interface {
	CreateUser(username, password, firstname, lastname, email, dob, avatar, address string) (db.User, error)
	UpdateUser(username, password string, updatecontent db.User) (db.User, error)
	GetUsers() ([]db.User, error)
	DeleteUser(username, password string) (db.User, error)
}
