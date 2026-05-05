package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CardId int
	Brand  string
	Model  string
	Year   int
}

var cars []Car

func CreateCar(ctx *gin.Context) {
	var newCar Car
	if err := ctx.BindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid request body"))
		return
	}
	newCar.CardId = len(cars) + 1
	cars = append(cars, newCar)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Car created successfully", "car": newCar})
}

func UpdateCar(ctx *gin.Context) {
	carId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid car id"))
		return
	}

	condition := false
	var updatedCar Car

	if err := ctx.BindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid request body"))
		return
	}
	for i, car := range cars {
		if carId == car.CardId {
			updatedCar.CardId = car.CardId
			cars[i] = updatedCar
			condition = true
			break
		}
	}

	if condition {
		ctx.JSON(http.StatusOK, gin.H{"message": "Car updated successfully", "car": updatedCar})
	} else {
		ctx.AbortWithError(http.StatusNotFound, errors.New("Car not found"))
	}

}

func GetCars(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Cars found", "cars": cars})
}

func GetCarByID(ctx *gin.Context) {
	carID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid car id"))
		return
	}

	condition := false
	var carData Car

	for _, car := range cars {
		if carID == car.CardId {
			carData = car
			condition = true
			break
		}
	}

	if condition {
		ctx.JSON(http.StatusOK, gin.H{"message": "Car found", "car": carData})
	} else {
		ctx.AbortWithError(http.StatusNotFound, errors.New("Car not found"))
	}
}

func DeleteCar(ctx *gin.Context) {
	carID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid car id"))
		return
	}

	for i, car := range cars {
		if carID == car.CardId {
			cars = append(cars[:i], cars[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
			return
		}
	}

	ctx.AbortWithError(http.StatusNotFound, errors.New("Car not found"))
}
