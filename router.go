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

	// POST
	r.POST("/ridekaki/v1/RideKaki", ride.CreateRideKakiRide)
	r.Run(":" + PORT)
}
