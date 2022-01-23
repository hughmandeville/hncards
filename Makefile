# Makefile

## HELP:
.PHONY: help
## help: Show this help message.
help:
	@echo "Usage: make [target]\n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## :
## BUILD:

.PHONY: build-server
## build-server: Build Go server code.
build-server:
	go build -o /dev/null ./...

## :
## DEPENDENCIES:

.PHONY: dep-clean-server
## dep-clean-server: Clean up dependency files.
dep-clean-server:
	rm go.mod
	rm go.sum

.PHONY: dep-get-server
## dep-get-server: Get Go modules.
dep-get-server:
	go mod tidy

.PHONY: dep-init-server
## dep-init-server: Initialize Go modules.
dep-init-server:
	go mod init

.PHONY: dep-update-server
## dep-update-server: Update Go modules.
dep-update-server:
	go get -u ./...

## :
## RUN:

.PHONY: run-server
## run-server: Run Go server locally (on port 8080).
run-server:
	cd cmd/server; go run ./...

## :
