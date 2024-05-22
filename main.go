package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Prince2347X/hope-backend/handlers"
)

func main() {
	godotenv.Load("secrets/.env")

	router := gin.Default()
	router.GET("/ping", handlers.HandlePing)
	router.GET("/hospitals", handlers.HandleHospitals)
	router.Run(":8080")
}
