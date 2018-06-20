#!/bin/bash
docker ps
docker pull dkrinke/supermarketapi:latest
docker stop supermarket-api
docker rm supermarket-api
docker run -d -p 8080:8080 --name supermarket-api dkrinke/supermarketapi:latest
