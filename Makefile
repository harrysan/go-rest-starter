.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

RELEASE_VERSION = v1.0.0

APP 			= financetracker

all: start

start:
	@go run main.go start

build:
	@go build main.go

swagger:
	@swag init --parseDependency --generalInfo ./main.go --output ./docs

# go install github.com/google/wire/cmd/wire@latest
wire:
	@wire gen ./internal/wirex

clean:
	rm -rf data

seed:
	@go run main.go seed