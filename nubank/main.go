package main

import (
	"nubank/database"
	"nubank/service"
)

func main() {

	// assure server closing at end of execution
	defer service.StopServer()

	// call db client constructor
	database.CreateClient()

	// call start server function
	service.StartServer()
}
