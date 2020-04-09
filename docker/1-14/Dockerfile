FROM golang:1.14-alpine

WORKDIR /go/src/app
COPY . .
COPY docker/1-14/go.mod .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["morpheus-go-sample"]