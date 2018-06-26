package supermarketService

import (
	"encoding/json"          //For creating JSON responses
	"github.com/gorilla/mux" //Router that will take requests and decide what should be done (go get github.com/gorilla/mux)
	"net/http"               //Provides the representation of HTTP requests, responses, and is Responsible for running the server
	"regexp"                 //Regular Expressions
	"supermarketDB"          //Provides methods to interact with the "Database"
	"supermarketProduce"     //Provides the Produce Struct
)

//Strings to be used throughout the service
const (
	/*****************Responses************************/
	success = "Success"
	failure = "Failure"
	/*****************Messages*************************/
	badRequest   = "Bad request"
	located      = "Produce located"
	notLocated   = "Could not locate Produce with the provided code"
	added        = "Produce added"
	removed      = "Produce removed"
	notUnique    = "Code is not unique"
	nameFailVal  = "The Produce name failed validation"
	codeFailVal  = "The Produce code failed validation"
	priceFailVal = "The Produce price failed validation"
	/*****************Regex*****************************/
	nameRegex  = "^[a-zA-Z0-9 ]+$"
	codeRegex  = "^([a-zA-Z0-9]{4}-){3}[a-zA-Z0-9]{4}$"
	priceRegex = `^\$[0-9]*\.[0-9]{2}$`
)

//Reponse object containing produce
//Result: Status of the request (Success/Failure)
//Message: Additional information
//Produce: Produce object
type ProduceResponseObject struct {
	Result  string
	Message string
	Produce supermarketProduce.Produce
}

//Reponse object
//Result: Status of the request (Success/Failure)
//Message: Additional information
type ResponseObject struct {
	Result  string
	Message string
}

//Delete request
//Code: Produce code
type DeleteRequestObject struct {
	Code string
}

//GetAllProduce returns a list of produce back to the client
func GetAllProduce(w http.ResponseWriter, r *http.Request) {
	locatedProduce := supermarketDB.ReadAll()            //Retrieve all produce from db
	w.WriteHeader(http.StatusOK)              //Set status to OK
	json.NewEncoder(w).Encode(locatedProduce) //Return list of produce
}

//GetProduceByCode returns a produce back to the client
//Produce returned is based on the provided code
func GetProduceByCode(w http.ResponseWriter, r *http.Request) {

	//Get code from request
	code := getCode(r)

	//Retrieve produce(locatedProduce) with matching code from db
	//locatedBoolean is true/false indicating if the produce was found
	locatedBoolean, locatedProduce := supermarketDB.Read(code)

	if locatedBoolean { //If produce was found
		var successObject ProduceResponseObject  //Create Response object
		successObject.Result = success           //Indicate success
		successObject.Message = located          //Indicate Produce was located
		successObject.Produce = locatedProduce   //Set the produce to be returned
		w.WriteHeader(http.StatusOK)             //Set status to OK
		json.NewEncoder(w).Encode(successObject) //Return Response object
	} else { //If Produce was not found
		var failureObject ProduceResponseObject  //Create Response object
		failureObject.Result = failure           //Indicate failure
		failureObject.Message = notLocated       //Indicate Produce was not located
		w.WriteHeader(http.StatusNotFound)       //Set status to NotFound
		json.NewEncoder(w).Encode(failureObject) //Return Response object
	}
}

//Resets the db data to default (For demo purposes)
func ResetProduceData(w http.ResponseWriter, r *http.Request) {
	supermarketDB.ResetData()                      //Resets the db data to default
	w.WriteHeader(http.StatusNoContent) //Set status to No Content
}

//Add Produce to the db
//The produce added is return upon a successful add
func AddProduce(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body) //Decode the body

	var incomingProduce supermarketProduce.Produce //Create Produce object
	decoder.Decode(&incomingProduce)    //Save body data to the Produce object

	//Perform validation on incoming Produce
	//resultString: Success/Failure
	//resultMessage: Reason for Success/Failure
	//resultBoolean: booleaning indicating if passed validation
	resultString, resultMessage, resultBoolean := validateProduce(incomingProduce)

	if resultBoolean { //If Passed validation
		//Add Produce to DB
		var added, savedProduce = supermarketDB.AddProduce(incomingProduce)
		if added { //If added successfully //Return Response object
			var resultObject ProduceResponseObject  //Create Response object
			resultObject.Result = resultString      //Set result (Indicate success)
			resultObject.Message = resultMessage    //Set message (Indicate added successfully)
			resultObject.Produce = savedProduce     //Set the produce to be returned
			w.WriteHeader(http.StatusCreated)       //Set status to Created
			json.NewEncoder(w).Encode(resultObject) //Return Response object
		} else { //If add was unsuccessful
			var resultObject ResponseObject         //Create Response object
			resultObject.Result = failure           //Indicate failure
			resultObject.Message = notUnique        //Indicate that the code was not unique
			w.WriteHeader(http.StatusBadRequest)    //Set status to BadRequest
			json.NewEncoder(w).Encode(resultObject) //Return Response object
		}
	} else { //If Failed validation
		var resultObject ResponseObject         //Create Response object
		resultObject.Result = resultString      //Set result (Indicate failure)
		resultObject.Message = resultMessage    //Set message (Indicate failed validation)
		w.WriteHeader(http.StatusBadRequest)    //Set status to BadRequest
		json.NewEncoder(w).Encode(resultObject) //Return Response object
	}
}

//Delete Produce from the db
//The produce deleted is based on the provided code
func DeleteProduce(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body) //Decode the body

	var incomingDeleteRequest DeleteRequestObject //Create Delete request
	decoder.Decode(&incomingDeleteRequest)        //Save body to the Delete request

	var code = incomingDeleteRequest.Code //Save Code from the Delete request

	if validateCode(code) { //If the code validates
		if supermarketDB.DeleteProduce(code) { //If produce is found and deleted with the provided code
			var successObject ResponseObject         //Create Response object
			successObject.Result = success           //Indicate success
			successObject.Message = removed          //Indicate Produce was removed
			w.WriteHeader(http.StatusOK)             //Set status to OK
			json.NewEncoder(w).Encode(successObject) //Return Response object
		} else { //If deletion fails with the provided code
			var failureObject ResponseObject         //Create Response object
			failureObject.Result = failure           //Indicate failure
			failureObject.Message = notLocated       //Indicate Produce was not located
			w.WriteHeader(http.StatusNotFound)       //Set status to NotFound
			json.NewEncoder(w).Encode(failureObject) //Return Response object
		}
	} else { //If the code fails validation
		var failureObject ResponseObject         //Create Response object
		failureObject.Result = failure           //Indicate failure
		failureObject.Message = codeFailVal      //Indicate that the provide code failed validation
		w.WriteHeader(http.StatusBadRequest)     //Set the status to BadRequest
		json.NewEncoder(w).Encode(failureObject) //Return Response object
	}
}

//getCode extracts and returns the code provided in te request
func getCode(r *http.Request) string {
	vars := mux.Vars(r) //Get url variables
	return vars["code"] //Get variable which matching key "code"
}

//validateProduce is used to validate Produce objects
func validateProduce(produce supermarketProduce.Produce) (string, string, bool) {

	var resultBoolean bool   //Is true if validation passes
	var resultString string  //Indicates Success/Failure
	var resultMessage string //Indicates the reason for failure

	switch {
	case validateName(produce.Name) == false: //If Name fails validation
		resultBoolean = false       //Set Result to false
		resultString = failure      //Set the Result string to indicate Failure
		resultMessage = nameFailVal //Set the message to indicate name failed validation
	case validateCode(produce.Code) == false: //If Code fails Validation
		resultBoolean = false       //Set Result to false
		resultString = failure      //Set the Result string to indicate Failure
		resultMessage = codeFailVal //Set the message to indicate code failed validation
	case validatePrice(produce.Price) == false: //If Price fails Validation
		resultBoolean = false        //Set Result to false
		resultString = failure       //Set the Result string to indicate Failure
		resultMessage = priceFailVal //Set the message to indicate price failed validation
	default: //If all pass validation
		resultBoolean = true   //Set Result to true
		resultString = success //Set the Result string to indicate Success
		resultMessage = added  //Set the message to indicate Produce was added
	}

	//Return the resultString, resultMessage, and resultBoolean
	return resultString, resultMessage, resultBoolean
}

//Returns boolean indicating if name passed validation
func validateName(name string) bool {
	//Validate the Name against regex
	return regexp.MustCompile(nameRegex).MatchString(name)
}

//Returns boolean indicating if code passed validation
func validateCode(code string) bool {
	//Validate the Code against regex
	return regexp.MustCompile(codeRegex).MatchString(code)
}

//Returns boolean indicating if price passed validation
func validatePrice(price string) bool {
	//Vaidate the Price against regex
	return regexp.MustCompile(priceRegex).MatchString(price)
}
