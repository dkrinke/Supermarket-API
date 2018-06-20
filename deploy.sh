#!/bin/bash
docker ps
docker stop supermarket-api
docker run -d -p 8080:8080 --name supermarket-api dkrinke/supermarketapi:latest
