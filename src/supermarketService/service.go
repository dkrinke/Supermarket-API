package service

import (
	"encoding/json" //For creating JSON responses
	"github.com/gorilla/mux"
	"net/http"
	"supermarketDB"
	"supermarketProduce"
)

func GetAllProduce(w http.ResponseWriter, r *http.Request) {

  // Retrieve all produce from db
	var locatedProduce []produce.Produce = db.ReadAll()
	json.NewEncoder(w).Encode(locatedProduce)
}

func GetProduceByCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) // Get url variables
	code := vars["code"] // Get variable which matching key "code"

  // Retrieve produce with matching code from db
	var locatedProduce produce.Produce = db.Read(code)
	json.NewEncoder(w).Encode(locatedProduce)
}
