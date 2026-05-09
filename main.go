package main

import (
	"belajar-go/database"
	"belajar-go/models"
	"belajar-go/routers"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "baharma1899"
	dbname   = "car_db"
)

var PORT = ":8030"

type Employee struct {
	ID   int
	Name string
	Age  int
}

func main() {
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	var err error

	database.DB, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	database.DB.AutoMigrate(
		&models.Car{},
		&models.Product{},
	)

	fmt.Println("Successfully connected to the database!")

	server := routers.StartServer()

	fmt.Printf("Server is running on port %s\n", PORT)
	if err := server.Run(PORT); err != nil {
		log.Fatal("Error starting server: ", err)
	}

	err = server.Run(":8030")
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, World!"
	fmt.Fprintln(w, msg)
}
