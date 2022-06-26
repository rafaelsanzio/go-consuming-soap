FROM golang:1.18.2

WORKDIR /go-consuming-soap

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build ./cmd/api/main.go

CMD ["./main"]