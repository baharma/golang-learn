package routers

import (
	"belajar-go/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controller.GetCars)

	return router
}
