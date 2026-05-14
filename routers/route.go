package routers

import (
	"belajar-go/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controller.GetCars)
	router.POST("/cars", controller.CreateCar)
	router.GET("/cars/:id", controller.GetCarById)
	router.PUT("/cars/:id", controller.UpdateCar)
	router.DELETE("/cars/:id", controller.DeleteCar)

	router.GET("/products", controller.GetProducts)
	router.POST("/products", controller.CreateProduct)
	router.GET("/products/:id", controller.GetProductById)
	router.PUT("/products/:id", controller.UpdateProduct)
	router.DELETE("/products/:id", controller.DeleteProduct)
	return router
}
