package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kevingil/coffee-recipe/routes"
)

func main() {
	//Load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	routes.InitRouter()
}
