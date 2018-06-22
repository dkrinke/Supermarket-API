package service

import (
	"encoding/json" //For creating JSON responses
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"supermarketDB"
	"supermarketProduce"
)

// Returned on success
type ProduceResponseObject struct {
	Result  string
	Message string
	Produce produce.Produce
}

// Returned on failure
type ResponseObject struct {
	Result  string
	Message string
}

// Delete request
type DeleteRequestObject struct {
	Code string
}

func GetAllProduce(w http.ResponseWriter, r *http.Request) {

	// Retrieve all produce from db
	var locatedProduce = db.ReadAll()
	json.NewEncoder(w).Encode(locatedProduce)
}

func GetProduceByCode(w http.ResponseWriter, r *http.Request) {

	code := getCode(r)

	// Retrieve produce with matching code from db
	var locatedProduce = db.Read(code)
	json.NewEncoder(w).Encode(locatedProduce)
}

func AddProduce(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var incomingProduce produce.Produce
	err := decoder.Decode(&incomingProduce)
	if err != nil {
		panic(err)
	}

	// Perform validation on incoming Produce
	resultString, resultMessage, resultBoolean := validateProduce(incomingProduce)

	if resultBoolean {
		// Build success object
		var savedProduce = db.AddProduce(incomingProduce)
		var resultObject ProduceResponseObject
		resultObject.Result = resultString
		resultObject.Message = resultMessage
		resultObject.Produce = savedProduce
		json.NewEncoder(w).Encode(resultObject)
	} else {
		// Build failure object
		var resultObject ResponseObject
		resultObject.Result = resultString
		resultObject.Message = resultMessage
		json.NewEncoder(w).Encode(resultObject)
	}
}

func DeleteProduce(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var incomingDeleteRequest DeleteRequestObject
	err := decoder.Decode(&incomingDeleteRequest)
	if err != nil {
		panic(err)
	}

	var code = incomingDeleteRequest.Code

	if(validateCode(code)) {
		if(db.DeleteProduce(code)) {
			var successObject ResponseObject
			successObject.Result = "Success"
			successObject.Message = "Produce removed"
			json.NewEncoder(w).Encode(successObject)
		} else {
			var failureObject ResponseObject
			failureObject.Result = "Failure"
			failureObject.Message = "Could not locate Produce with this code"
			json.NewEncoder(w).Encode(failureObject)
		}
	} else {
		var failureObject ResponseObject
		failureObject.Result = "Failure"
		failureObject.Message = "The Produce code failed validation"
		json.NewEncoder(w).Encode(failureObject)
	}
}

func getCode(r *http.Request) string {
	vars := mux.Vars(r) // Get url variables
	return vars["code"] // Get variable which matching key "code"
}

func validateProduce(produce produce.Produce) (string, string, bool) {

	var resultBoolean bool   // true if validation passes
	var resultString string  // Success/Failure
	var resultMessage string // Reason for failure

	switch {
	// If Name fails validation
	case validateName(produce.Name) == false:
		resultBoolean = false
		resultString = "Failure"
		resultMessage = "The Produce name failed validation"
	// If Code fails validation
	case validateCode(produce.Code) == false:
		resultBoolean = false
		resultString = "Failure"
		resultMessage = "The Produce code failed validation"
	// If Price fails validation
	case validatePrice(produce.Price) == false:
		resultBoolean = false
		resultString = "Failure"
		resultMessage = "The Produce price failed validation"
	// If all pass validation
	default:
		resultBoolean = true
		resultString = "Success"
		resultMessage = "Produce added"
	}

	return resultString, resultMessage, resultBoolean
}

func validateName(name string) bool {
	// Validate the Name
	return regexp.MustCompile("^[a-zA-Z0-9 ]+$").MatchString(name)
}

func validateCode(code string) bool {
	// Validate the Code (Assuming the code provided is unique)
	return regexp.MustCompile("([a-zA-Z0-9]{4}-){3}[a-zA-Z0-9]{4}").MatchString(code)
}

func validatePrice(price string) bool {
	// Vaidate the price
	return regexp.MustCompile(`^\$[0-9]*\.[0-9]{2}$`).MatchString(price)
}
