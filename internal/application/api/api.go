package api

import (
	"lab5.cmo/internal/application/core"
	"lab5.cmo/internal/port"
)

type Application struct {
	db port.DbPort
}

func NewApplication(db port.DbPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) CreateUser(username, password, firstname, lastname, email, dob, avatar, address string) (core.User, error) {
	user, err := a.db.CreateUser(username, password, firstname, lastname, email, dob, avatar, address)
	if err != nil {
		return core.User{}, err
	}
	return user, nil
}
func (a Application) GetUsers() ([]core.User, error) {
	users, err := a.db.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (a Application) UpdateUser(username, password string, updateContent core.User) (core.User, error) {
	user, err := a.db.UpdateUser(username, password, updateContent)
	if err != nil {
		return core.User{}, err
	}
	return user, nil
}
func (a Application) DeleteUser(username, password string) (core.User, error) {
	user, err := a.db.DeleteUser(username, password)
	if err != nil {
		return core.User{}, err
	}
	return user, nil
}
