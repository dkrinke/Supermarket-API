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

  code := getCode(r)

  // Retrieve produce with matching code from db
	var locatedProduce produce.Produce = db.Read(code)
	json.NewEncoder(w).Encode(locatedProduce)
}

func getCode(r *http.Request) string {
  vars := mux.Vars(r) // Get url variables
	return vars["code"] // Get variable which matching key "code"
}
