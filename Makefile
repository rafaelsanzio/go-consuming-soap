GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=simple_service
LINTER=golangci-lint
MAIN_PATH=./cmd/api/main.go

all:
		test build

mod:
	  $(GOMOD) tidy -v

test:
		$(GOTEST) ./... -v

build:
		$(GOBUILD) -o $(BINARY_NAME) -v

lint:
		$(LINTER) run -v

start:
		$(GOCMD) run $(MAIN_PATH)