package main

import (
	"encoding/json" //For creating JSON responses
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Produce struct {
	Name  string
	Code  string
	Price float64
}

var (
	produceArray = []Produce{
		Produce{
			Name:  "Lettuce",
			Code:  "A12T-4GH7-QPL9-3N4M",
			Price: 3.46,
		},
		Produce{
			Name:  "Peach",
			Code:  "E5T6-9UI3-TH15-QR88",
			Price: 2.99,
		},
		Produce{
			Name:  "Green Pepper",
			Code:  "YRT6-72AS-K736-L4AR",
			Price: 0.79,
		},
		Produce{
			Name:  "Gala Apple",
			Code:  "TQ4C-VV6T-75ZX-1RMR",
			Price: 3.59,
		},
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, "<!DOCTYPE html><html><body><h1>SupermarketAPI</h1></body></html>")
		return
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!") //Health check to confirm the server is running
}

func getAllProduce(w http.ResponseWriter, r *http.Request) {
	//Send json response
	json.NewEncoder(w).Encode(map[string][]Produce{"produce": produceArray})
}

func getProduceByCode(w http.ResponseWriter, r *http.Request) {
	//Send json response
	vars := mux.Vars(r)
	code := vars["code"]
	var locatedProduce Produce

	for _, produce := range produceArray {
    if produce.Code == code {
        locatedProduce = produce
    }
	}

	json.NewEncoder(w).Encode(map[string]Produce{"produce": locatedProduce})
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce", getAllProduce).Methods("GET")
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce/{code}", getProduceByCode).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func returnTrueForFirstTest() bool {
	return true
}
