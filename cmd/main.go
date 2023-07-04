package main

import (
	_ "github.com/go-sql-driver/mysql"
	"lab5.cmo/internal/adapters/framework/left/gingonic"
	// "lab5.cmo/internal/adapters/framework/left/gofiber"
	// "lab5.cmo/internal/adapters/framework/right/mongodbframework"
	mysqlframework "lab5.cmo/internal/adapters/framework/right/mysqlFramework"
	"lab5.cmo/internal/application/api"
)

func main() {
	// mongodb := mongodbframework.NewAdapter()
	mysqldb := mysqlframework.NewAdapter("mysql", "root:ThanhTung2!@tcp(localhost:3306)/goauth")

	// using mongodb as database
	// api := api.NewApplication(mongodb)

	// using mysql as database
	api := api.NewApplication(mysqldb)

	// using gin as http framework
	server := gingonic.NewAdapter(api)

	// using gofiber as http framework
	// server := gofiber.NewAdapter(api)

	server.Run()
}