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

.PHONY: clean-client
## clean-client: Remove React build directory and Node modules.
clean-client:
	@rm -rf client/build
	@rm -rf client/node_modules

## :
## DEPENDENCIES:

.PHONY: dep-get-client
## dep-get-client: Get Node modules.
dep-get-client:
	cd client; yarn

.PHONY: dep-update-client
## dep-update-client: Update Node modules.
dep-update-client:
	cd client; yarn upgrade

## :
## GITHUB:

.PHONY: gh-deploy
## gh-deploy: Deploy React client to GitHub Pages.
gh-deploy:
	cd client; yarn deploy

## :
## HACKER_NEWS:

.PHONY: get-hn-ts
## get-hn-ts: Get Hacker News Top Stories.
get-hn-ts:
	go run cmd/get_hn/get_hn.go -out client/public/hn_topstories.json -verbose

## :
## RUN:

.PHONY: run-client
## run-client: Run React client locally (on port 3000).
run-client:
	cd client; yarn start

## :
