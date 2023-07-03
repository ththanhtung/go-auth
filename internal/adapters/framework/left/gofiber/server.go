package gofiber

import (
	"github.com/gofiber/fiber/v2"
	"lab5.cmo/internal/port"
)

type Adapter struct {
	api port.ApiPort
}

func NewAdapter(a port.ApiPort) *Adapter {
	return &Adapter{
		api: a,
	}
}

func (a *Adapter) Run(){
	server := fiber.New()

	
	server.Post("auth/signup", a.createUser)
	server.Patch("users/:username", a.updateUser)
	server.Get("users", a.getUsers)
	server.Delete("users/:username", a.deleteUser)

	server.Listen(":8080")
}