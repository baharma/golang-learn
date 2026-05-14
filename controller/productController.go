package controller

import (
	"belajar-go/database"
	"belajar-go/models"

	"github.com/gin-gonic/gin"
)

var products []models.Product

func GetProducts(ctx *gin.Context) {
	database.DB.Find(&products)
	ctx.JSON(200, products)
}

func CreateProduct(ctx *gin.Context) {
	var newProduct models.Product
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	database.DB.Create(&newProduct)
	ctx.JSON(201, newProduct)
}

func GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product
	database.DB.First(&product, id)
	ctx.JSON(200, product)
}

func UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	var updatedProduct models.Product
	if err := ctx.ShouldBindJSON(&updatedProduct); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	database.DB.Save(&product)

	ctx.JSON(200, product)
}

func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	database.DB.Delete(&product)
	ctx.JSON(204, nil)
}
