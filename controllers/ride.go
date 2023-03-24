package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ride-kaki-backend/models"
	"ride-kaki-backend/services"
	"strconv"
)

type RideController struct{}

func (t RideController) CreateRideKakiRide(c *gin.Context) {
	rideKaki := new(models.RideKaki)
	err := c.BindJSON(rideKaki)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid RideKaki Object" + err.Error()})
		return
	}

	startingLatitude := rideKaki.StartingLatitude
	startingLongitude := rideKaki.StartingLongitude
	endingLatitude := rideKaki.EndingLatitude
	endingLongitude := rideKaki.EndingLongitude

	gojek := new(models.Gojek)
	gojek.SendPrioritisedOrder = "true"
	gojek.UserSelectedServiceType = "50"
	gojek.StartingLatitude = startingLatitude
	gojek.StartingLongitude = startingLongitude
	gojek.EndingLatitude = endingLatitude
	gojek.EndingLongitude = endingLongitude

	gojek4EconomyPrice, gojek6EconomyPrice, gojek4PremiumPrice, err := services.CreateGojekService(gojek, c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	tada := new(models.Tada)
	tada.IsStudentFleet = false
	startingLocation := new(models.Locations)
	endingLocation := new(models.Locations)
	startingLocation.Latitude, _ = strconv.ParseFloat(startingLatitude, 64)
	startingLocation.Longitude, _ = strconv.ParseFloat(startingLongitude, 64)
	endingLocation.Latitude, _ = strconv.ParseFloat(endingLatitude, 64)
	endingLocation.Longitude, _ = strconv.ParseFloat(endingLongitude, 64)

	var locations []models.Locations
	locations = append(locations, *startingLocation)
	locations = append(locations, *endingLocation)
	tada.Locations = locations

	tada4EconomyPrice, tada6EconomyPrice, tada4PremiumPrice, err := services.CreateTadaService(tada, c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rideKakiOutput := new(models.RideKakiOutput)
	rideKakiOutput.Gojek4EconomyPrice = gojek4EconomyPrice
	rideKakiOutput.Gojek6EconomyPrice = gojek6EconomyPrice
	rideKakiOutput.Gojek4PremiumPrice = gojek4PremiumPrice
	rideKakiOutput.Tada4EconomyPrice = tada4EconomyPrice
	rideKakiOutput.Tada6EconomyPrice = tada6EconomyPrice
	rideKakiOutput.Tada4PremiumPrice = tada4PremiumPrice

	c.JSON(http.StatusOK, rideKakiOutput)
}