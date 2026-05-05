package main

import (
	"belajar-go/routers"
	"encoding/json"
	"fmt"
	"net/http"
)

var PORT = ":8030"

type Employee struct {
	ID   int
	Name string
	Age  int
}

var emp = []Employee{
	{Name: "John", Age: 30},
	{Name: "Jane", Age: 25},
}

func main() {
	var PORT = ":8030"
	routers.StartServer().Run(PORT)
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/employees", getEmployees)
	// http.HandleFunc("/employee", createEmployee)
	// fmt.Printf("Server is running on http://localhost%s\n", PORT)
	// http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, World!"
	fmt.Fprintln(w, msg)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}
	var input Employee
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newEmployee := Employee{
		ID:   len(emp) + 1,
		Name: input.Name,
		Age:  input.Age,
	}
	emp = append(emp, newEmployee)
	json.NewEncoder(w).Encode(newEmployee)
	return

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(emp)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
