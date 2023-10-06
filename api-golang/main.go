package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var customers []Customer

func main() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	r.HandleFunc("/Customers", getcustomers).Methods("GET")
	r.HandleFunc("/Customer/add", addCustomer).Methods("POST")
	r.HandleFunc("/Customer/update", updateCustomer).Methods("PUT")
	r.HandleFunc("/Customer/remove/{id}", removeCustomer).Methods("DELETE")

	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}

func getcustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newCustomer Customer
	newCustomer.ID = uuid.New().String()

	err := json.NewDecoder(r.Body).Decode(&newCustomer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customers = append(customers, newCustomer)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updateCustomer Customer

	err := json.NewDecoder(r.Body).Decode(&updateCustomer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, customer := range customers {
		if customer.ID == updateCustomer.ID {
			customers[i] = updateCustomer
			break
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updateCustomer)
	}
}

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	customerID := vars["id"]

	index := -1
	for i, customer := range customers {
		if customer.ID == customerID {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	customers = append(customers[:index], customers[index+1:]...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{ Message string }{"Customer removed successfully"})
}
