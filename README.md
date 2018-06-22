# Supermarket API

Supermarket application that can add, delete, and fetch all produce in the system.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
1. Install Go: https://golang.org/doc/install

### Installing

#### GitHub
1. Clone this repo
2. Set the GOPATH to the route directory
3. Run go get github.com/gorilla/mux

#### DockerHub
1. Run: docker pull dkrinke/supermarketapi:latest

## Running the tests
1. Run go test ./...

## Deployment

### Build
1. Run go install supermarketAPI
2. Run bin/supermaketAPI

### Docker
1. Build docker image: docker build -t dkrinke/supermarketapi:latest . -f ./src/supermarketAPI/Dockerfile
2. Run docker image: docker run -d -p 8080:8080 --name supermarket-api dkrinke/supermarketapi:latest

### DockerHub
1. Run docker pull dkrinke/supermarketapi:latest
2. Run docker image: docker run -d -p 8080:8080 --name supermarket-api dkrinke/supermarketapi:latest

## Production
http://api.krinkes.com/

### Supported Endpoints
- GET: http://api.krinkes.com/healthcheck
  - Returns: "Still alive!"
- GET http://api.krinkes.com/api/v1/produce
  - Returns: List of Produce
- GET http://api.krinkes.com/api/v1/produce/A12T-4GH7-QPL9-3N4M
  - Returns: A single Produce that matches the provided code
- POST http://api.krinkes.com/api/v1/produce
  - Payload:  
    {  
    &nbsp;&nbsp;&nbsp;&nbsp;"Name": "Green Apples",  
    &nbsp;&nbsp;&nbsp;&nbsp;"Code": "123A-123A-123A-124A", // Assumed that a unique code will be provided  
    &nbsp;&nbsp;&nbsp;&nbsp;"Price": "$1.11"  
    }  
  - Returns: Success or Failure with messages containing the reason



