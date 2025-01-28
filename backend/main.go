package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Transaction struct {
	ID     string  `json:"id"`
	User   string  `json:"user"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	// ...
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/transactions", getTransactions).Methods("GET")
	router.HandleFunc("/suspicious", getSuspiciousTransactions).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Enable CORS
	corsObj := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(corsObj, headersOk, methodsOk)(router)))
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := []Transaction{
		{
			ID:     "txn_001",
			User:   "John Doe",
			Amount: 150.75,
			Date:   "2025-01-27T12:00:00Z",
		},
		{
			ID:     "txn_002",
			User:   "Jane Doe",
			Amount: 250.00,
			Date:   "2025-01-27T13:30:00Z",
		},
		// ...
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func getSuspiciousTransactions(w http.ResponseWriter, r *http.Request) {
	suspicious := []Transaction{
		{
			ID:     "txn_002",
			User:   "Jane Doe",
			Amount: 250.00,
			Date:   "2025-01-27T13:30:00Z",
		},
		// ...
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suspicious)
}
