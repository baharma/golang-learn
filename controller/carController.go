package controller

import (
	"belajar-go/database"
	"belajar-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var cars []models.Car


func GetCars(ctx *gin.Context) {

	result := database.DB.Find(&cars)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cars"})
		return
	}

	ctx.JSON(http.StatusOK, cars)

}



func createCar(ctx *gin.Context) {

}