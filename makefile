#!make
prepare:
	dep ensure
generate:
	go run scripts/gqlgen.go generate 
run: 
	go run main.go
test:
	go test -v -cover -covermode=atomic ./...
tls:
	openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout ./key.pem -out ./cert.pem