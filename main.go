package main

import (
	"log"

	"github.com/agnjunio/go-rest-demo/database"
	"github.com/agnjunio/go-rest-demo/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found")
	}

	mongo, err := database.Connect()
	if err != nil {
		log.Fatal("Unable to connect to database.")
	}
	defer database.Disconnect(mongo)

	router := routes.InitRouter(mongo)

	err = router.Run()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
