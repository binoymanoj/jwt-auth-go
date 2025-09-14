# Variables
MAINFILE 		 = main.go
BINARY_NAME  = fireworks-api
BINARY_DIR   = bin

# Default Target (executes when no targets specified)
.DEFAULT_GOAL := all

# Phony targets
.PHONY: all fmt vet run build clean help db-up db-down

## help: help menu
help: 
	@echo "Usage:" 
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## all: format, vet and build the application
all: fmt vet build

## db-up: docker compose up -d (run the database)
db-up: 
	@echo "--------------------------"
	@echo "- Running Docker Compose -"
	@echo "--------------------------"
	docker compose up -d

## db-down: docker compose down (stop the database)
db-down: 
	@echo "---------------------------"
	@echo "- Stopping Docker Compose -"
	@echo "---------------------------"
	docker compose down

## fmt: formats the application
fmt:
	@echo "------------------------------"
	@echo "- Formatting the Application -"
	@echo "------------------------------"
	go fmt

## vet: check for errors using go's vet command
vet:
	@echo "-----------------------"
	@echo "- Checking for Errors -"
	@echo "-----------------------"
	go vet .

# TODO: uncomment this block when adding test 
## test: run all tests 
# run: fmt vet
# 	@echo "-----------------"
# 	@echo "- Running Tests -"
# 	@echo "-----------------"
# 	go test

## run: run the application
run: fmt vet
	@echo "------------------------------"
	@echo "- Running the Go Application -"
	@echo "------------------------------"
	go run $(MAINFILE)

## build: build the application (also does fmt and vet)
build: fmt vet
	@echo "-------------------------------"
	@echo "- Building the Go Application -"
	@echo "-------------------------------"
	@mkdir -p $(BINARY_DIR)
	GOARCH=amd64 GOOS=linux go build -o $(BINARY_DIR)/$(BINARY_NAME)-linux $(MAINFILE)
	GOARCH=amd64 GOOS=darwin go build -o $(BINARY_DIR)/$(BINARY_NAME)-darwin $(MAINFILE)
	GOARCH=amd64 GOOS=windows go build -o $(BINARY_DIR)/$(BINARY_NAME)-windows.exe $(MAINFILE)
	chmod +x $(BINARY_DIR)/$(BINARY_NAME)-linux

## clean: cleans all the temporary and binary (output) directories
clean:
	@echo "----------------------------------------------"
	@echo "- Removing temporary and bin directories     -"
	@echo "----------------------------------------------"
	rm -rf tmp/ $(BINARY_DIR)/

