package services

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"os"
	"ride-kaki-backend/models"
)

func CreateTadaService(dataTada *models.Tada, c *gin.Context) (gojekFourEconomySeaterPrice float64, gojekSixEconomySeaterPrice float64, gojekFourPremiumSeaterPrice float64, err error) {
	url := "https://backend.tada.global/ridesvc/v1/products/baseSearch"

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
	//var data interface{}
	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)

	var output models.TadaOutput
	if err = mapstructure.Decode(data, &output); err != nil {
		// handle error
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	tadaFourEconomySeaterPrice := output.Products[0].NetPrice
	tadaSixEconomySeaterPrice := output.Products[2].NetPrice
	tadaFourPremiumSeaterPrice := output.Products[3].NetPrice

	return tadaFourEconomySeaterPrice, tadaSixEconomySeaterPrice, tadaFourPremiumSeaterPrice, err
}