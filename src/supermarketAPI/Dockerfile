FROM golang:1.6.2

WORKDIR /go/src/supermarketAPI
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["supermarketAPI"]
