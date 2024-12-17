package main

import (
	"fmt"
	"log"
	"server/internal/handler"
	"server/internal/repository/postgres"
)

// @title Fitness API
// @version 1.0
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	db, err := postgres.NewDatabase()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not initialize database connection: %s", err))
	}

	handlers := handler.NewHandler(db)
	
	app := handlers.Router()
	app.Listen(":8080")
}
