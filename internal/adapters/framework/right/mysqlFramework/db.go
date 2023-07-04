package mysqlframework

import (
	"database/sql"
	"log"

	"lab5.cmo/internal/application/core"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) *Adapter {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Panic(err.Error())
	}

	return &Adapter{
		db: db,
	}
}

func (a *Adapter) CreateUser(username, password, firstname, lastname, email, dob, avatar, address string) (core.User, error) {
	stmt, err := a.db.Prepare("INSERT INTO users (username, password, firstname, lastname, email, dob, avatar, address) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return core.User{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(username, password, firstname, lastname, email, dob, avatar, address)
	if err != nil {
		return core.User{}, err
	}

	var user core.User

	id, err := res.LastInsertId()
	if err != nil {
		return core.User{}, err
	}

	row := a.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	err = row.Scan(&user.UserId, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Email, &user.DoB, &user.Avatar, &user.Address)
	if err != nil {
		return core.User{}, err
	}

	return user, nil
}

func (a *Adapter) UpdateUser(username, password string, updatecontent core.User) (core.User, error) {
	updateStmt := "UPDATE users SET username = ?, firstname = ?, lastname = ?, email = ?, dob = ?, avatar = ?, address = ? WHERE username = ? AND password = ?"

	stmt, err := a.db.Prepare(updateStmt)
	if err != nil {
		return core.User{}, err
	}

	res, err := stmt.Exec(updatecontent.Username, updatecontent.Firstname, updatecontent.Lastname, updatecontent.Email, updatecontent.DoB, updatecontent.Avatar, updatecontent.Address, username, password)
	affectedRows, _ := res.RowsAffected()
	if err != nil || affectedRows != 1 {
		return core.User{}, err
	}

	row := a.db.QueryRow("SELECT * FROM users WHERE username = ? AND password = ?", username, password)

	var updatedUser core.User

	if err = row.Scan(&updatedUser.UserId, &updatedUser.Username, &updatedUser.Password, &updatedUser.Firstname, &updatedUser.Lastname, &updatedUser.Email, &updatedUser.DoB, &updatedUser.Avatar, &updatedUser.Address); err != nil {
		return core.User{}, err
	}

	return updatedUser, nil
}
func (a *Adapter) GetUsers() ([]core.User, error) {
	rows, err := a.db.Query("SELECT * FROM users")
	if err != nil {
		return []core.User{}, err
	}
	defer rows.Close()

	var users []core.User
	for rows.Next() {
		var user core.User

		err = rows.Scan(&user.UserId, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Email, &user.DoB, &user.Avatar, &user.Address)
		if err != nil {
			return []core.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}
func (a *Adapter) DeleteUser(username, password string) (core.User, error) {
	getDeltedUserStmt := "SELECT * FROM users WHERE username = ? AND password = ?"
	deleteStmt := "DELETE FROM users WHERE username = ? AND password = ?"

	var deletedUser core.User
	row := a.db.QueryRow(getDeltedUserStmt, username, password)
	err := row.Scan(&deletedUser.UserId, &deletedUser.Username, &deletedUser.Password, &deletedUser.Firstname, &deletedUser.Lastname, &deletedUser.Email, &deletedUser.DoB, &deletedUser.Avatar, &deletedUser.Address)
	if err != nil {
		return core.User{}, err
	}

	stmt, err := a.db.Prepare(deleteStmt)
	if err != nil {
		return core.User{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(username, password)
	affectedRows, _ := res.RowsAffected()
	if err != nil || affectedRows != 1 {
		return core.User{}, err
	}

	return deletedUser, nil
}
