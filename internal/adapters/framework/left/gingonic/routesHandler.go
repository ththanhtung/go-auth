package gingonic

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"lab5.cmo/internal/application/core"
)

type User struct {
	Username  string `json:"username"`
	Password  string `Json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	DoB       string `json:"dob"`
	Avatar    string `Json:"avatar"`
	Address   string `jsson:"address"`
}

func (a Adapter) createUser(c *gin.Context) {

	var req User

	c.ShouldBindJSON(&req)

	log.Println(req)

	newUser, err := a.api.CreateUser(req.Username, req.Password, req.Firstname, req.Lastname, req.Email, req.DoB, req.Avatar, req.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(201, newUser)
}

func (a Adapter) updateUser(c *gin.Context) {
	type UpdatedContent struct {
		Update   json.RawMessage `json:"update"`
		Password string          `json:"password"`
	}

	var updatedUser core.User

	username := c.Param("username")
	log.Println("param:", username)

	var updateContent UpdatedContent

	c.ShouldBindJSON(&updateContent)

	err := json.Unmarshal([]byte(updateContent.Update), &updatedUser)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}

	user, err := a.api.UpdateUser(username, updateContent.Password, updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a Adapter) getUsers(c *gin.Context) {
	users, err := a.api.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (a Adapter) deleteUser(c *gin.Context) {
	username := c.Param("username")

	type DeleteUser struct {
		Password string `json:"password"`
	}

	var user DeleteUser

	c.ShouldBindJSON(&user)

	deletedUser, err := a.api.DeleteUser(username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	
	c.JSON(http.StatusOK, deletedUser)
}
