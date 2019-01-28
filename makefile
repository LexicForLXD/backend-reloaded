#!make
include .env
export

prepare:
	dep ensure
generate:
	go run scripts/gqlgen.go generate 
run: 
	go run main.go
test:
	go test -v -cover -covermode=atomic ./...