package main

import (
	"belajar-go/database"
	"belajar-go/models"
	"fmt"
	"log"
)

func main() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(&models.Product{}, &models.Car{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrasi Berhasil!")

}
