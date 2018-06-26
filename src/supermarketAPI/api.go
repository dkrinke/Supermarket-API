package main

import (
	"encoding/json"          //For creating JSON responses
	"fmt"                    //STDOUT
	"github.com/gorilla/mux" //Router that will take requests and decide what should be done (go get github.com/gorilla/mux)
	"io/ioutil"              //Implements some I/O utility functions
	"log"                    //Logs when the server exits
	"net/http"               //Provides the representation of HTTP requests, responses, and is Responsible for running the server
	"supermarketService"     //Service used to process Produce requests
)

//handler returns html to the client.
func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("html/supermarketAPI.html") //Read from html file into bytes
	if err != nil {
		fmt.Println(err)
	}

	str := string(b) // convert content to a 'string'

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, str)
	return
}

//healthCheck returns "Still alive!" to the client
//Can be used to determine if the application is running
func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!") //Health check to confirm the server is running
}

func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/", handler).Methods("GET")                //Will provide informaton about the application
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET") //Health Check

	/**********************Produce API Version One*******************************/
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce", supermarketService.GetAllProduce).Methods("GET")
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce/{code}", supermarketService.GetProduceByCode).Methods("GET")
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce", supermarketService.AddProduce).Methods("POST")
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce/reset", supermarketService.ResetProduceData).Methods("POST") //Reset date to default (Demo purposes)
	router.PathPrefix("/api/v1").Subrouter().HandleFunc("/produce", supermarketService.DeleteProduce).Methods("DELETE")
	/****************************************************************************/

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
