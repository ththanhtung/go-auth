package gofiber

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func (a Adapter) createUser(c *fiber.Ctx) error {

	var req User

	c.BodyParser(&req)

	log.Println(req)

	newUser, err := a.api.CreateUser(req.Username, req.Password, req.Firstname, req.Lastname, req.Email, req.DoB, req.Avatar, req.Address)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"err": err.Error(),
		})
		return err
	}

	c.Status(http.StatusCreated).JSON(newUser)
	return nil
}

func (a Adapter) updateUser(c *fiber.Ctx) error {
	type UpdatedContent struct {
		Update   json.RawMessage `json:"update"`
		Password string          `json:"password"`
	}

	var updatedUser core.User

	username := c.Params("username")
	log.Println("param:", username)

	var updateContent UpdatedContent

	c.BodyParser(&updateContent)

	err := json.Unmarshal([]byte(updateContent.Update), &updatedUser)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return err
	}

	user, err := a.api.UpdateUser(username, updateContent.Password, updatedUser)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK).JSON(user)

	return nil
}

func (a Adapter) getUsers(c *fiber.Ctx) error {
	users, err := a.api.GetUsers()
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK).JSON(users)

	return nil
}

func (a Adapter) deleteUser(c *fiber.Ctx) error {
	username := c.Params("username")

	type DeleteUser struct {
		Password string `json:"password"`
	}

	var user DeleteUser

	c.BodyParser(&user)

	deletedUser, err := a.api.DeleteUser(username, user.Password)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK).JSON(deletedUser)

	return nil
}
