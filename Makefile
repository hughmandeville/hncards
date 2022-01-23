# Makefile

## HELP:
.PHONY: help
## help: Show this help message.
help:
	@echo "Usage: make [target]\n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## :
## BUILD:

.PHONY: build-client
## build-client: Build React client code.
build-client:
	cd client; yarn && yarn build

.PHONY: build-server
## build-server: Build Go server code.
build-server:
	go build -o /dev/null ./...

.PHONY: clean-client
## clean-client: Remove React build directory and Node modules.
clean-client:
	@rm -rf client/build
	@rm -rf client/node_modules

## :
## DEPENDENCIES:

.PHONY: dep-clean-server
## dep-clean-server: Clean up dependency files.
dep-clean-server:
	rm go.mod
	rm go.sum

.PHONY: dep-get-client
## dep-get-client: Get Node modules.
dep-get-client:
	cd client; yarn

.PHONY: dep-get-server
## dep-get-server: Get Go modules.
dep-get-server:
	go mod tidy

.PHONY: dep-init-server
## dep-init-server: Initialize Go modules.
dep-init-server:
	go mod init

.PHONY: dep-update-all
## dep-update-all: Update Go modules and Node modules.
dep-update-all: dep-update-server dep-update-client

.PHONY: dep-update-client
## dep-update-client: Update Node modules.
dep-update-client:
	cd client; yarn upgrade

.PHONY: dep-update-server
## dep-update-server: Update Go modules.
dep-update-server:
	go get -u ./...

## :
## RUN:

.PHONY: run-client
## run-client: Run React client locally (on port 3000).
run-client:
	cd client; yarn start

.PHONY: run-server
## run-server: Run Go server locally (on port 8080).
run-server:
	cd cmd/server; go run ./...

## :
