package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"ride-kaki-backend/controllers"
)

func InitRoutes() {
	godotenv.Load(".env")
	PORT := os.Getenv("SERVER_PORT")
	r := gin.Default()

	ride := new(controllers.RideController)

	// GET /users
	r.POST("/ridekaki/v1/products/baseSearch", ride.HandleBaseSearch)
	//r.GET("/transport/v1/estimate", ride.HandleEstimate)
	r.Run(":" + PORT)
}
