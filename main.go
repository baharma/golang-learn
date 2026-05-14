package main

import (
	"belajar-go/database"
	"belajar-go/routers"
	"fmt"
	"log"
	"net/http"
)

var PORT = ":8030"

type Employee struct {
	ID   int
	Name string
	Age  int
}

func main() {
	database.ConnectDatabase()
	var err error

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

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
