package main

import (
	"encoding/json" //For creating JSON responses
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"supermarketService"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<!DOCTYPE html><html><body><h1>SupermarketAPI</h1></body></html>")
	return
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!") //Health check to confirm the server is running
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce", service.GetAllProduce).Methods("GET")
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce/{code}", service.GetProduceByCode).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
