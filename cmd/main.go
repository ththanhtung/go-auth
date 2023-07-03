package main

import (
	"lab5.cmo/internal/adapters/framework/left/gingonic"
	// "lab5.cmo/internal/adapters/framework/left/gofiber"
	"lab5.cmo/internal/adapters/framework/right/db"
	"lab5.cmo/internal/application/api"
)

func main() {
	db := db.NewAdapter()

	api := api.NewApplication(db)

	// gin framework
	server := gingonic.NewAdapter(api)

	// gofiber framework
	// server := gofiber.NewAdapter(api)

	server.Run()
}