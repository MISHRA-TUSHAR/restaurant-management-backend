package main

import (
	"os"

	"github.com/MISHRA-TUSHAR/restaurant-management-backend/database"
	"github.com/MISHRA-TUSHAR/restaurant-management-backend/middleware"
	"github.com/MISHRA-TUSHAR/restaurant-management-backend/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

}
