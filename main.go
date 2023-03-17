package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"

	// "fmt"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strings"
)

func main() {
	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ride-kaki")
	// if err != nil {
	//     panic(err.Error())
	// }
	// defer db.Close()

	// fmt.Println("Hello, Modules!")

	godotenv.Load(".env")
	r := gin.Default()
	// GET /users
	r.GET("/ridesvc/v1/products/baseSearch", handleBaseSearch)
	r.Run("localhost:8080")
}

func handleBaseSearch(c *gin.Context) {
	url := "https://backend.tada.global/ridesvc/v1/products/baseSearch"

	// create the request body
	requestBody := map[string]interface{}{
		"isStudentFleet": false,
		"locations": []map[string]interface{}{
			{
				"address":       "81 Victoria St, Singapore 188065",
				"googlePlaceId": "ChIJGddBg6MZ2jERACsxW7Ovm_4",
				"latitude":      1.2962727,
				"longitude":     103.8501578,
				"name":          "SMU",
			},
			{
				"address":       "68 Orchard Rd, Singapore 238839",
				"googlePlaceId": "ChIJR1Fyfr0Z2jERzwO-AZiJ-HM",
				"latitude":      1.3005317,
				"longitude":     103.8452356,
				"name":          "Plaza Singapura",
			},
		},
	}

	// create the HTTP client
	client := &http.Client{}

	// create the request object
	requestBodyBytes, err := json.Marshal(requestBody)
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
