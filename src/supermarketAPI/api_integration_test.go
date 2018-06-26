package main

import (
"flag"
"fmt"                    //STDOUT
"io/ioutil"              //Implements some I/O utility functions
"net/http"               //Provides the representation of HTTP requests, responses, and is Responsible for running the server
"testing"
)

const (
  baseUrl = "http://localhost:9000"
  healthCheckUrl = "/healthcheck"
)

var (
	runIntegrationTests = flag.Bool("integration", false, "Run the integration tests (in addition to the unit tests)")
)

/********************************Test HealthCheck*******************************/
func TestHealthCheck(t *testing.T) {
  if !*runIntegrationTests {
      t.Skip("To run this test, use: go test -integration")
  }

  response, err := http.Get(baseUrl+healthCheckUrl)
    if err != nil {
        t.Errorf("HealthCheck endpoint returned an error: %s", err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Printf("%s\n", string(contents))
    }
}
