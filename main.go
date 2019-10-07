package main

import (
	"log"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
)

//Customer is the data structure that defines a customer
type Customer struct {
	CustID int `json:"cust_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
	FaveGame string `json:"fave_game"`
}

//Customers is an array of all instances of Customer
type Customers []Customer

func allCustomers(w http.ResponseWriter, r *http.Request) {
	customers := Customers{
		Customer{CustID:00000, FirstName:"Test", LastName:"Tested", CreatedAt: time.Now() , LastLogin: time.Now(), FaveGame:"Test Game"},
	}
	
	fmt.Println("Reached all customer api endpoint")
	json.NewEncoder(w).Encode(customers)
}


func home (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/customers", allCustomers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
