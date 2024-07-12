build:
	go build -o bin/bc-helper src/main.go

run: build
	go run src/main.go

setup:	init tidy 

init:
	go mod init FraktalDeFiDAO/bc-helper
	
tidy:
	go mod tidy
	go mod vendor