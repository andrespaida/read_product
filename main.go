package main

import (
	"read_product/database"
	"read_product/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	router := gin.Default()
	router.GET("/products", handlers.ReadProducts)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4001"
	}

	router.Run(":" + port)
}