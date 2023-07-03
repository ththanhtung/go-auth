package main

import (
	"lab5.cmo/internal/adapters/framework/left/http"
	"lab5.cmo/internal/adapters/framework/right/db"
	"lab5.cmo/internal/application/api"
)

func main() {
	db := db.NewAdapter()

	api := api.NewApplication(db)

	server := http.NewAdapter(api)

	server.Run()
}