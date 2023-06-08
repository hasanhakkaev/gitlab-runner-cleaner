# ==================================================================================== #
# HELPERS
# ==================================================================================== #
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

VERSION=$(shell git rev-parse --short HEAD)

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
.PHONE: init
init:
	go mod tidy -v
	go mod verify
	go mod download

## lint : run linters
.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...


## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go test -race -vet=off ./...
	go mod verify


# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build: build the cmd/cleaner application
.PHONY: build
build:
	go mod verify
	go build -o=./bin/cleaner .

## run: run the cmd/cleaner application
.PHONY: run
run:
	go run main.go 

.PHONE: test
test: 
	go test -v ./...

## clean: clean up the build artifacts
.PHONY: clean
clean:
	rm -rf ./bin

## docker: build the docker image
.PHONY: docker
docker:
	docker buildx build --load -t cleaner:$(VERSION) .