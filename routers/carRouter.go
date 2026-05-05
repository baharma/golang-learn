package routers

import (
	"belajar-go/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/cars", controller.CreateCar)
	router.GET("/cars", controller.GetCars)
	router.GET("/cars/:id", controller.GetCarByID)
	router.PUT("/cars/:id", controller.UpdateCar)
	router.DELETE("/cars/:id", controller.DeleteCar)

	return router
}
