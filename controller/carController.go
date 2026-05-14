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

func CreateCar(ctx *gin.Context) {
	var newCar models.Car
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	result := database.DB.Create(&newCar)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	ctx.JSON(http.StatusCreated, newCar)
}

func GetCarById(ctx *gin.Context) {
	id := ctx.Param("id")
	var car models.Car
	result := database.DB.First(&car, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	ctx.JSON(http.StatusOK, car)
}

func UpdateCar(ctx *gin.Context) {
	id := ctx.Param("id")
	var car models.Car

	// 1. Cari data lama di database berdasarkan ID
	if err := database.DB.First(&car, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found (ID " + id + " tidak ada)"})
		return
	}

	// 2. Bind data baru dari request body ke struct 'input'
	var input models.Car
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data yang dikirim tidak valid"})
		return
	}

	// 3. Update field dan gunakan .Scan(&car) untuk mengisi variabel 'car' dengan data terbaru
	// Ini adalah kunci agar 'car' berisi data setelah diupdate
	result := database.DB.Model(&car).Updates(models.Car{
		Name:       input.Name,
		ProductsId: input.ProductsId,
	}).Scan(&car) // <--- PENTING: Mengambil data hasil update kembali ke struct 'car'

	// Cek jika ada error saat proses update
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}

	// 4. Sekarang 'car' sudah berisi data terbaru, kirim sebagai respon
	ctx.JSON(http.StatusOK, car)
}

func DeleteCar(ctx *gin.Context) {
	id := ctx.Param("id")
	var car models.Car
	result := database.DB.First(&car, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	database.DB.Delete(&car)
	ctx.JSON(http.StatusNoContent, nil)
}
