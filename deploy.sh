docker ps
docker stop supermarket-api
docker run -d -p 8080:8080 dkrinke/supermarketapi:latest
