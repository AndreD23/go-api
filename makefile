.PHONY: default run build test docs clear

# Variables
APP_NAME=gopportunities

# Tasks
default: run-with-docs

run:
	@go run main.go

run-with-docs:
	@swag init
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./ ...

docs:
	@swag init

clear:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
