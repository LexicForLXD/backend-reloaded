#!make
prepare:
	dep ensure
generate:
	go run scripts/gqlgen.go generate 
run: 
	go run main.go
test:
	ifdef ci
		go test -coverprofile c.out -v ./...
	else
		go test -v -cover -covermode=atomic ./...
	endif
tls:
	ifdef ci
		openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout ./key.pem -out ./cert.pem -subj "/C=DE/ST=NRW/L=Cologne/O=LexicForLXD/OU=IT Department/CN=lexic.test.de"
	else
		openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout ./key.pem -out ./cert.pem
	endif