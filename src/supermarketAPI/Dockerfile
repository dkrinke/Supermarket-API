FROM golang:1.6.2

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["supermarketAPI"]
