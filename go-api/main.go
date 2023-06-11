package main

import (
	"go-api/app"
	"go-api/db"
)

func main() {

	db.StartDbEngine()
	app.StartApp()

}
