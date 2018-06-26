package main

import (
	"bytes"     //Implements functions for the manipulation of byte slices. It is analogous to the facilities of the strings package
	"flag"      //Implements command-line flag parsing.
	"io/ioutil" //Implements some I/O utility functions
	"net/http"  //Provides the representation of HTTP requests, responses, and is Responsible for running the server
	"testing"
)

const (
	/*********Application Endpoints***********************************************/
	baseUrl        = "http://localhost:9000" //Port 9000 for integration tests only
	healthCheckUrl = "/healthcheck"
	apiUrl         = "/api"
	v1Url          = "/v1"
	produceUrl     = "/produce"

	/*********Request Body Format*************************************************/
	jsonFormat = "application/json"

	/*********Test Input**********************************************************/
	produceCode       = "/A12T-4GH7-QPL9-3N4M"
	addProduceRequest = `{
    "Name": "Green Apples",
  	"Code": "1213-456B-789C-DEF5",
  	"Price": "$1.12"
  }`
	deleteProduceCode = "/1213-456B-789C-DEF5"

	/*********Test Expected Output************************************************/
	healthCheckExpected      = `"Still alive!"` + "\n"
	getAllProduceExpected    = `[{"Name":"Lettuce","Code":"A12T-4GH7-QPL9-3N4M","Price":"$3.46"},{"Name":"Peach","Code":"E5T6-9UI3-TH15-QR88","Price":"$2.99"},{"Name":"Green Pepper","Code":"YRT6-72AS-K736-L4AR","Price":"$0.79"},{"Name":"Gala Apple","Code":"TQ4C-VV6T-75ZX-1RMR","Price":"$3.59"}]` + "\n"
	getProduceByCodeExpected = `{"Result":"Success","Message":"Produce located","Produce":{"Name":"Lettuce","Code":"A12T-4GH7-QPL9-3N4M","Price":"$3.46"}}` + "\n"
	addProduceExpected       = `{"Result":"Success","Message":"Produce added","Produce":{"Name":"Green Apples","Code":"1213-456B-789C-DEF5","Price":"$1.12"}}` + "\n"
	deleteProduceExpected    = `{"Result":"Success","Message":"Produce removed"}` + "\n"
)

var (
	//This prevents the tests from running if the integration flag is not provided
	runIntegrationTests = flag.Bool("integration", false, "Run the integration tests (in addition to the unit tests)")
)

/********************************Test HealthCheck*******************************/
func TestHealthCheck(t *testing.T) {
	if !*runIntegrationTests { //Check if integration
		t.Skip("To run this test, use: go test -integration")
	}

	var url = baseUrl + healthCheckUrl //Build the url endpoint

	response, err := http.Get(url) //GET request to url endpoint
	if err != nil {                //Check if error was returned
		t.Errorf("HealthCheck endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close() //Close body after completion

		contents, err := ioutil.ReadAll(response.Body) //Read the response from the body
		if err != nil {                                //Check if error was returned
			t.Errorf("HealthCheck endpoint returned content that could not be parsed: %s", err)
		}

		if string(contents) != healthCheckExpected { //Check if the result matches expected
			t.Errorf(`Expected %s but was %s`, healthCheckExpected, string(contents))
		}
	}
}

/********************************Test GetAllProduce*****************************/
func TestGetAllProduce(t *testing.T) {
	if !*runIntegrationTests { //Check if integration
		t.Skip("To run this test, use: go test -integration")
	}

	var url = baseUrl + apiUrl + v1Url + produceUrl //Build the url endpoint
	response, err := http.Get(url)                  //GET request to url endpoint
	if err != nil {                                 //Check if error was returned
		t.Errorf("GetAllProduce endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close() //Close body after completion

		contents, err := ioutil.ReadAll(response.Body) //Read the response from the body
		if err != nil {                                //Check if error was returned
			t.Errorf("GetAllProduce endpoint returned content that could not be parsed: %s", err)
		}

		if string(contents) != getAllProduceExpected { //Check if the result matches expected
			t.Errorf(`Expected %s but was %s`, getAllProduceExpected, string(contents))
		}
	}
}

/********************************Test GetProduceByCode**************************/
func TestGetProduceByCode(t *testing.T) {
	if !*runIntegrationTests { //Check if integration
		t.Skip("To run this test, use: go test -integration")
	}

	var url = baseUrl + apiUrl + v1Url + produceUrl + produceCode //Build the url endpoint
	response, err := http.Get(url)                                //GET request to url endpoint
	if err != nil {                                               //Check if error was returned
		t.Errorf("GetProduceByCode endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close() //Close body after completion

		contents, err := ioutil.ReadAll(response.Body) //Read the response from the body
		if err != nil {                                //Check if error was returned
			t.Errorf("GetProduceByCode endpoint returned content that could not be parsed: %s", err)
		}

		if string(contents) != getProduceByCodeExpected { //Check if the result matches expected
			t.Errorf(`Expected %s but was %s`, getProduceByCodeExpected, string(contents))
		}
	}
}

/********************************Test AddProduce********************************/
func TestAddProduce(t *testing.T) {
	if !*runIntegrationTests { //Check if integration
		t.Skip("To run this test, use: go test -integration")
	}

	var url = baseUrl + apiUrl + v1Url + produceUrl                    //Build the url endpoint
	var data = []byte(addProduceRequest)                               //Build data
	response, err := http.Post(url, jsonFormat, bytes.NewBuffer(data)) //POST request to url endpoint
	if err != nil {                                                    //Check if error was returned
		t.Errorf("AddProduce endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close() //Close body after completion

		contents, err := ioutil.ReadAll(response.Body) //Read the response from the body
		if err != nil {                                //Check if error was returned
			t.Errorf("AddProduce endpoint returned content that could not be parsed: %s", err)
		}

		if string(contents) != addProduceExpected { //Check if the result matches expected
			t.Errorf(`Expected %s but was %s`, addProduceExpected, string(contents))
		}
	}
}

/********************************Test DeleteProduce*****************************/
func TestDeleteProduce(t *testing.T) {
	if !*runIntegrationTests { //Check if integration
		t.Skip("To run this test, use: go test -integration")
	}

	var url = baseUrl + apiUrl + v1Url + produceUrl + deleteProduceCode //Build the url endpoint

	data := []byte("")                                        //Build data (Set to "" because nil would cause NewRequest to panic)
	requestBody := bytes.NewBuffer(data)                      //Build Request body
	request, _ := http.NewRequest("DELETE", url, requestBody) //Build Request object

	response, err := http.DefaultClient.Do(request) //Delete request to irl endpoint
	if err != nil {                                 //Check if error was returned
		t.Errorf("DeleteProduce endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close() //Close body after completion

		contents, err := ioutil.ReadAll(response.Body) //Read the response from the body
		if err != nil {                                //Check if error was returned
			t.Errorf("DeleteProduce endpoint returned content that could not be parsed: %s", err)
		}

		if string(contents) != deleteProduceExpected { //Check if the result matches expected
			t.Errorf(`Expected %s but was %s`, deleteProduceExpected, string(contents))
		}
	}
}
