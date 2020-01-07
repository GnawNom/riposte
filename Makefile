APP=riposte
BINARY=$(APP)
BINARY_UNIX=$(BINARY)_unix
.PHONY: build
## build: build the application
build: clean
	@go build -o ${BINARY} main.go
	##@go build -o ${BINARY} cmd/${APP}/main.go

build-linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) cmd/${APP}/main.go

## run: runs application and checks for race conditions
.PHONY: run
run:
	go run -race main.go

## clean: cleans the binary
.PHONY: clean
clean:
	@rm -rf ${APP}

## setup: setup go modules
.PHONY: setup
setup:
	@go mod init \
		&& go mod tidy \

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
