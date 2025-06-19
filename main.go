package main

import (
	"log"
	"os"
	"read_product/database"
	"read_product/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Conexión a la base de datos
	database.ConnectDB()

	// Inicializar router Gin
	router := gin.Default()

	// Habilitar CORS globalmente
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // o reemplaza con []string{"http://98.85.86.231"} para mayor seguridad
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Ruta principal
	router.GET("/products", handlers.ReadProducts)

	// Puerto desde .env o por defecto 4001
	port := os.Getenv("PORT")
	if port == "" {
		port = "4001"
	}

	// Iniciar servidor
	router.Run(":" + port)
}
