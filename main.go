package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// initialize a router
	router := gin.Default()

	// TODO: add routes

	fmt.Println("Starting the server at port: " + port)
	router.Run(":" + port)
}
