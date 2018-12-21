generate:
	go run scripts/gqlgen.go generate 
run: 
	DB_HOST=localhost DB_USER=lexic DB_NAME=lexic DB_PASSWORD=password go run main.go