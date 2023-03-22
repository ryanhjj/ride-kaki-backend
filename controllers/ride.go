package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ride-kaki-backend/models"
	"ride-kaki-backend/services"
)

type RideController struct{}

func (t RideController) CreateTadaRide(c *gin.Context) {
	dataTada := new(models.Tada)
	err := c.BindJSON(dataTada)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Tada Object" + err.Error()})
		return
	}

	data, err := services.CreateTadaRide(dataTada, c)

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

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Gojek Object" + err.Error()})
		return
	}

	data, err := services.CreateGojekRide(dataGojek, c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
