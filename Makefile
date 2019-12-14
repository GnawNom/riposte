APP=riposte
.PHONY: build
## build: build the application
build: clean
	@go build -o ${APP} main.go

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