package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var customers []Customer

func main() {
	http.HandleFunc("/Customers", getcustomers)
	http.HandleFunc("/Customer/add", addCustomer)
	http.HandleFunc("/Customer/update", updateCustomer)

	http.ListenAndServe(":8080", nil)
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
