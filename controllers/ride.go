package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"ride-kaki-backend/models"
)

type RideController struct{}

func (t RideController) CreateTadaRide(c *gin.Context) {
	url := "https://backend.tada.global/ridesvc/v1/products/baseSearch"

	dataTada := new(models.Tada)
	err := c.BindJSON(dataTada)

	// create the HTTP client
	client := &http.Client{}

	// create the request object
	requestBodyBytes, err := json.Marshal(dataTada)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// set the Authorization header
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TADA_TOKEN"))

	// set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// make the request
	resp, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// parse the response
	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (t RideController) CreateGoJekRide(c *gin.Context) {
	dataGojek := new(models.Gojek)
	err := c.BindJSON(dataGojek)

	// Get the query parameters from the request
	sendPrioritisedOrder := dataGojek.SendPrioritisedOrder
	serviceType := dataGojek.UserSelectedServiceType
	startingLatitude := dataGojek.StartingLatitude
	startingLongitude := dataGojek.StartingLongitude
	endingLatitude := dataGojek.EndingLatitude
	endingLongitude := dataGojek.EndingLongitude
	waypoints := fmt.Sprintf("%s%%2C%s%%7C%s%%2C%s", startingLatitude, startingLongitude, endingLatitude, endingLongitude)

	// Set the query parameters
	params := url.Values{}
	params.Set("send_prioritised_order", sendPrioritisedOrder)
	params.Set("user_selected_service_type", serviceType)

	// Build the request URL
	reqURL := fmt.Sprintf("https://api.gojekapi.com/transport/v1/estimate?%s&waypoints=%s", params.Encode(), waypoints)

	// create the HTTP client
	client := &http.Client{}

	// Create a new GET request with the bearer token in the header
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GOJEK_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
