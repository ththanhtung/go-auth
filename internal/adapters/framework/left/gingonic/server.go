package gingonic

import (
	"github.com/gin-gonic/gin"
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

func (a Adapter) Run() {
	server := gin.Default()

	server.POST("auth/signup", a.createUser)
	server.PATCH("users/:username", a.updateUser)
	server.GET("users", a.getUsers)
	server.DELETE("users/:username", a.deleteUser)

	server.Run(":8080")
}
