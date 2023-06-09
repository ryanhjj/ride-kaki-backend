package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"ride-kaki-backend/models"
)

func CreateGojekService(dataGojek *models.Gojek, c *gin.Context) (gojek4EconomyPrice float64, gojek6EconomyPrice float64, gojek4PremiumPrice float64, err error) {
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
	var output models.GojekOutput
	err = json.NewDecoder(resp.Body).Decode(&output)

	gojek4EconomyPrice = float64(output.Data.ServiceTypes[0].PaymentMethods[0].MinPriceWithoutVoucher) / 100
	gojek6EconomyPrice = float64(output.Data.ServiceTypes[1].PaymentMethods[0].MinPriceWithoutVoucher) / 100
	gojek4PremiumPrice = float64(output.Data.ServiceTypes[4].PaymentMethods[0].MinPriceWithoutVoucher) / 100

	return gojek4EconomyPrice, gojek6EconomyPrice, gojek4PremiumPrice, err
}
