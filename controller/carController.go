package controller

import (
	"belajar-go/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var cars []Car

func GetCars(ctx *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name FROM cars")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query cars"})
		return
	}
	defer rows.Close()

	cars = []Car{}
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.ID, &car.Name); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan car"})
			return
		}
		cars = append(cars, car)
	}

	if len(cars) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No cars found"})
		return
	}

	ctx.JSON(http.StatusOK, cars)
}

func CreateCar(ctx *gin.Context) {
	var newCar Car
	if err := ctx.BindJSON(&newCar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := database.DB.Exec("INSERT INTO cars (name) VALUES (?)", newCar.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	id, _ := result.LastInsertId()
	newCar.ID = int(id)
	cars = append(cars, newCar)

	ctx.JSON(http.StatusCreated, newCar)
}

func GetCarByID(ctx *gin.Context) {
	var getCar Car
	id := ctx.Param("id")

	err := database.DB.QueryRow("SELECT id, name FROM cars WHERE id = ?", id).Scan(&getCar.ID, &getCar.Name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	ctx.JSON(http.StatusOK, getCar)
}

func UpdateCar(ctx *gin.Context) {
	var updateCar Car
	id := ctx.Param("id")

	if err := ctx.BindJSON(&updateCar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := database.DB.Exec("UPDATE cars SET name = ? WHERE id = ?", updateCar.Name, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	updateCar.ID = int(rowsAffected)
	ctx.JSON(http.StatusOK, updateCar)
}

func DeleteCar(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := database.DB.Exec("DELETE FROM cars WHERE id = ?", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
