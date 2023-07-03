package main

import (
	"lab5.cmo/internal/adapters/framework/left/gingonic"
	// "lab5.cmo/internal/adapters/framework/left/gofiber"
	"lab5.cmo/internal/adapters/framework/right/mongodbframework"
	"lab5.cmo/internal/application/api"
)

func main() {
	mongodb := mongodbframework.NewAdapter()

	api := api.NewApplication(mongodb)

	// gin framework
	server := gingonic.NewAdapter(api)

	// gofiber framework
	// server := gofiber.NewAdapter(api)

	server.Run()
}