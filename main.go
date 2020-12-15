package main

import (
	"log"

	routes "github.com/agnjunio/go-rest-demo/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found")
	}

	router := routes.InitRouter()

	err = router.Run()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
