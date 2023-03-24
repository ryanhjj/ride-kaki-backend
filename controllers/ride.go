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

	// Launch a new goroutine for CreateGojekService
	gojekChan := make(chan [3]float64)
	go func() {
		gojek4EconomyPrice, gojek6EconomyPrice, gojek4PremiumPrice, err := services.CreateGojekService(gojek, c)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gojekChan <- [3]float64{gojek4EconomyPrice, gojek6EconomyPrice, gojek4PremiumPrice}
	}()

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

	// Launch a new goroutine for CreateTadaService
	tadaChan := make(chan [3]float64)
	go func() {
		tada4EconomyPrice, tada6EconomyPrice, tada4PremiumPrice, err := services.CreateTadaService(tada, c)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		tadaChan <- [3]float64{tada4EconomyPrice, tada6EconomyPrice, tada4PremiumPrice}
	}()

	// Wait for both goroutines to complete and retrieve their results
	gojekPrices := <-gojekChan
	tadaPrices := <-tadaChan

	rideKakiOutput := new(models.RideKakiOutput)
	rideKakiOutput.Gojek4EconomyPrice = gojekPrices[0]
	rideKakiOutput.Gojek6EconomyPrice = gojekPrices[1]
	rideKakiOutput.Gojek4PremiumPrice = gojekPrices[2]
	rideKakiOutput.Tada4EconomyPrice = tadaPrices[0]
	rideKakiOutput.Tada6EconomyPrice = tadaPrices[1]
	rideKakiOutput.Tada4PremiumPrice = tadaPrices[2]

	c.JSON(http.StatusOK, rideKakiOutput)
}
