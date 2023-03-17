package main

import (
	// "fmt"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
	"bytes"
	// "strings"
)

func main() {
	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ride-kaki")
	// if err != nil {
	//     panic(err.Error())
	// }
	// defer db.Close()

	// fmt.Println("Hello, Modules!")

	r := gin.Default()
	// GET /users
	r.GET("/ridesvc/v1/products/baseSearch", handleBaseSearch)

	r.Run()
}

const bearerToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InRhZGEtYXBwLWtleS1pZCJ9.eyJleHAiOjE2ODAwOTgzOTMsInVzZXJfbmFtZSI6ImM1ZTQ0Y2UyLWQ5M2QtNGMyZi04YTBkLWUwMjhhODgyODkyYyIsImF1dGhvcml0aWVzIjpbIlJPTEVfVVNFUiJdLCJqdGkiOiIwNDFhNGNlMi1kYjhhLTRlODQtOGZmMS03YWU5MzQzOGQxYjQiLCJjbGllbnRfaWQiOiJ0YWRhLXJpZGVyLWFwcCIsInNjb3BlIjpbInJlYWQiLCJ3cml0ZSIsInRydXN0Il19.k3o2QlxPgorN6I6Y_BhzXPtv17h_FO0u3sIk-RZ5ERK_hG5kC1SG4DRwlBuzwPlwdBwDUN0bdDgMAoXIZAp1Y9uwmmlBgdTtBn5qXCABR7N3MWjJ7sXEl6WKeRXKMIh35FMy5Ob2g6rJGpNGJpThTtPrG34PQBfXi6jCAt3ST7U5NWcfEkJb6MSo6apS2WXMQWpogUmDfXZO5gtgy12L6FlGIrcdw_gJcaolZRIwfLwppzweqUQv_nj79ecoQVHIpUNOPWMpjj6mUo3qFsS01w4Iot3CcwETChSndDSrhnXPVTjIq2uhEKH53Xrc6tkb88mzDCcoBNicNr_M1iE3Rg"

func handleBaseSearch(c *gin.Context) {
    url := "https://backend.tada.global/ridesvc/v1/products/baseSearch"
    
    // create the request body
    requestBody := map[string]interface{}{
        "isStudentFleet": false,
        "locations": []map[string]interface{}{
            {
                "address": "81 Victoria St, Singapore 188065",
                "googlePlaceId": "ChIJGddBg6MZ2jERACsxW7Ovm_4",
                "latitude": 1.2962727,
                "longitude": 103.8501578,
                "name": "SMU",
            },
            {
                "address": "68 Orchard Rd, Singapore 238839",
                "googlePlaceId": "ChIJR1Fyfr0Z2jERzwO-AZiJ-HM",
                "latitude": 1.3005317,
                "longitude": 103.8452356,
                "name": "Plaza Singapura",
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
    req.Header.Set("Authorization", "Bearer "+bearerToken)

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

