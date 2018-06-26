package main

import (
	"flag"      //Implements command-line flag parsing.
	"io/ioutil" //Implements some I/O utility functions
	"net/http"  //Provides the representation of HTTP requests, responses, and is Responsible for running the server
	"testing"
  "bytes"
)

const (
	baseUrl        = "http://localhost:9000" //Port 9000 for integration tests only
	healthCheckUrl = "/healthcheck"
	apiUrl         = "/api"
	v1Url          = "/v1"
	produceUrl     = "/produce"

  jsonFormat = "application/json"

  produceCode = "/A12T-4GH7-QPL9-3N4M"

  deleteProduceCode = "/1213-456B-789C-DEF5"

  healthCheckExpected = `"Still alive!"`+"\n"
  getAllProduceExpected = `[{"Name":"Lettuce","Code":"A12T-4GH7-QPL9-3N4M","Price":"$3.46"},{"Name":"Peach","Code":"E5T6-9UI3-TH15-QR88","Price":"$2.99"},{"Name":"Green Pepper","Code":"YRT6-72AS-K736-L4AR","Price":"$0.79"},{"Name":"Gala Apple","Code":"TQ4C-VV6T-75ZX-1RMR","Price":"$3.59"}]`+"\n"

  getProduceByCodeExpected = `{"Result":"Success","Message":"Produce located","Produce":{"Name":"Lettuce","Code":"A12T-4GH7-QPL9-3N4M","Price":"$3.46"}}`+"\n"

  addProduceExpected = `{"Result":"Success","Message":"Produce added","Produce":{"Name":"Green Apples","Code":"1213-456B-789C-DEF5","Price":"$1.12"}}`+"\n"
  deleteProduceExpected = `{"Result":"Success","Message":"Produce removed"}`+"\n"
  addProduceRequest = `{
    "Name": "Green Apples",
  	"Code": "1213-456B-789C-DEF5",
  	"Price": "$1.12"
  }`

)

var (
	runIntegrationTests = flag.Bool("integration", false, "Run the integration tests (in addition to the unit tests)")
)

/********************************Test HealthCheck*******************************/
func TestHealthCheck(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
	}

  var url = baseUrl + healthCheckUrl
	response, err := http.Get(url)
	if err != nil {
		t.Errorf("HealthCheck endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
      t.Errorf("HealthCheck endpoint returned content that could not be parsed: %s", err)
		}

    if string(contents) != healthCheckExpected {
      t.Errorf(`Expected %s but was %s`, healthCheckExpected, string(contents))
    }
	}
}

/********************************Test GetAllProduce*****************************/
func TestGetAllProduce(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
	}

  var url = baseUrl + apiUrl + v1Url + produceUrl
	response, err := http.Get(url)
	if err != nil {
		t.Errorf("GetAllProduce endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
      t.Errorf("GetAllProduce endpoint returned content that could not be parsed: %s", err)
		}

    if string(contents) != getAllProduceExpected {
      t.Errorf(`Expected %s but was %s`, getAllProduceExpected, string(contents))
    }
	}
}

/********************************Test GetProduceByCode**************************/
func TestGetProduceByCode(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
	}

  var url = baseUrl + apiUrl + v1Url + produceUrl + produceCode
	response, err := http.Get(url)
	if err != nil {
		t.Errorf("GetProduceByCode endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
      t.Errorf("GetProduceByCode endpoint returned content that could not be parsed: %s", err)
		}

    if string(contents) != getProduceByCodeExpected {
      t.Errorf(`Expected %s but was %s`, getProduceByCodeExpected, string(contents))
    }
	}
}

/********************************Test AddProduce********************************/
func TestAddProduce(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
	}

  var url = baseUrl + apiUrl + v1Url + produceUrl
  var jsonStr = []byte(addProduceRequest)
	response, err := http.Post(url, jsonFormat,  bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Errorf("AddProduce endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
      t.Errorf("AddProduce endpoint returned content that could not be parsed: %s", err)
		}

    if string(contents) != addProduceExpected {
      t.Errorf(`Expected %s but was %s`, addProduceExpected, string(contents))
    }
	}
}

/********************************Test DeleteProduce*****************************/
func TestDeleteProduce(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
	}

  var url = baseUrl + apiUrl + v1Url + produceUrl + deleteProduceCode
  t.Log(url)
  data := []byte("")
  requestBody := bytes.NewBuffer(data)
	request, _ := http.NewRequest("DELETE", url, requestBody)

  response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("AddProduce endpoint returned an error: %s", err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
      t.Errorf("AddProduce endpoint returned content that could not be parsed: %s", err)
		}

    if string(contents) != deleteProduceExpected {
      t.Errorf(`Expected %s but was %s`, deleteProduceExpected, string(contents))
    }
	}
}
