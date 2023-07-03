package api

import (
	"lab5.cmo/internal/adapters/framework/right/db"
	"lab5.cmo/internal/port"
)

type Application struct {
	db port.DbPort
}

func NewApplication(db port.DbPort)*Application {
	return &Application{
		db: db,
	}
}

func (a Application) CreateUser(username, password, firstname, lastname, email, dob, avatar, address string) (db.User ,error) {
	user, err := a.db.CreateUser(username, password, firstname, lastname, email, dob, avatar, address)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
func (a Application) GetUsers() ([]db.User,error) {
	users, err := a.db.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (a Application) UpdateUser(username, password string, updateContent db.User) (db.User ,error) {
	user, err := a.db.UpdateUser(username, password, updateContent)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
func (a Application) DeleteUser(username, password string) (db.User,error) {
	user, err := a.db.DeleteUser(username, password)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}