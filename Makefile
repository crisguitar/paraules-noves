include .env
export $(shell sed 's/=.*//' .env)

.PHONY: test build

deps:
	go get -v -t -d ./...

run:
	go run cmd/main.go

test:
	go test ./...

build:
	sh scripts/build.sh
