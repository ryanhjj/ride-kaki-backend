package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"ride-kaki-backend/models"
)

type RideController struct{}

func (t RideController) HandleBaseSearch(c *gin.Context) {
	url := "https://backend.tada.global/ridesvc/v1/products/baseSearch"

	dataCampaign := new(models.Fare)
	err := c.BindJSON(dataCampaign)

	// create the HTTP client
	client := &http.Client{}

	// create the request object
	requestBodyBytes, err := json.Marshal(dataCampaign)
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

//func (t RideController) HandleEstimate(c *gin.Context) {
//	// Get the query parameters from the request
//	sendPrioritisedOrder := c.Query("send_prioritised_order")
//	serviceType := c.Query("user_selected_service_type")
//	waypoints := c.Query("waypoints")
//
//	// Set the query parameters
//	params := url.Values{}
//	params.Set("send_prioritised_order", sendPrioritisedOrder)
//	params.Set("user_selected_service_type", serviceType)
//	params.Set("waypoints", waypoints)
//
//	// Build the request URL
//	reqURL := fmt.Sprintf("https://api.gojekapi.com/transport/v1/estimate?%s", params.Encode())
//
//	// create the HTTP client
//	client := &http.Client{}
//
//	// Create a new GET request with the bearer token in the header
//	req, err := http.NewRequest("GET", reqURL, nil)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	req.Header.Set("Authorization", "Bearer "+"eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJhdWQiOlsiZ29qZWs6Y29uc3VtZXI6YXBwIl0sImRhdCI6eyJhY3RpdmUiOiJ0cnVlIiwiYmxhY2tsaXN0ZWQiOiJmYWxzZSIsImNvdW50cnlfY29kZSI6Iis2NSIsImNyZWF0ZWRfYXQiOiIyMDIxLTA2LTMwVDEyOjI3OjA0WiIsImVtYWlsIjoiIiwiZW1haWxfdmVyaWZpZWQiOiJmYWxzZSIsImdvcGF5X2FjY291bnRfaWQiOiIwMS1mOWRkMGYzNzFlZWY0MTM5YTBhZTkzZGIxNGRjYjI3YS0yMSIsIm5hbWUiOiJSeWFuIEhvbmciLCJudW1iZXIiOiI4ODE0NzAxOCIsInBob25lIjoiKzY1ODgxNDcwMTgiLCJzaWduZWRfdXBfY291bnRyeSI6IlNHIiwid2FsbGV0X2lkIjoiMjExODEwNzQ3MTE3MTc0ODM0In0sImV4cCI6MTY4MDY2MTI1NCwiaWF0IjoxNjc3NzQ0MDMyLCJpc3MiOiJnb2lkIiwianRpIjoiY2FkNGVkZWEtMzEyYi00ZDI3LTkzNmYtNTI5MzMwMGJjNGY4Iiwic2NvcGVzIjpbXSwic2lkIjoiMGQ1YzA0MTItYmRkMi00OTVjLTgyMjItZmQ3ZGIxM2QxNmI2Iiwic3ViIjoiOTNiMmI3NDctMGY1Zi00NTA5LTk5MWQtODViZGJkNGZmYzE1IiwidWlkIjoiNzE3NzMwOTU3IiwidXR5cGUiOiJjdXN0b21lciJ9.O-5QqZL7jTZEG9kUfH1h1-uktEX7x7t2q4fnxxYP0owRFIqbBwrzKzKqijg2ygT4VlTqabH4rbYR2EsmIlsYRLZ-pWbDONRkO2KhcKN0-Gfi0LL8HSivMHw-GSX1FZLepI2_Eq2Q-QzqpTHFx5gGgmQw8_sW3ogdVb4cQdk4sz4")
//	req.Header.Set("Content-Type", "application/json")
//	// Send the request
//	resp, err := client.Do(req)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer resp.Body.Close()
//
//	var data interface{}
//	err = json.NewDecoder(resp.Body).Decode(&data)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"data": data,
//	})
//}
