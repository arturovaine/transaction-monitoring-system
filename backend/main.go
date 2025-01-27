package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/transactions", getTransactions).Methods("GET")
    router.HandleFunc("/suspicious", getSuspiciousTransactions).Methods("GET")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server running on port %s", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Transactions data endpoint")
}

func getSuspiciousTransactions(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Suspicious activity endpoint")
}
