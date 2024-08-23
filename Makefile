build: tidy
	go build -o bin/bc-helper src/main.go

build-versioner: tidy
	go build -o bin/bc-versioner scripts/versioning.go

build-version-control: tidy
	go build -o bin/bc-version-control scripts/vc.go

run: build
	./bin/bc-helper

run-versioner: build-versioner
	./bin/bc-versioner

run-version-control: build-version-control
	echo "New Version: ${VERSION}"
	./bin/bc-version-control ${VERSION}

setup:	init tidy

init:
	go mod init FraktalDeFiDAO/bc-helper

tidy:
	go mod tidy
	go mod vendor

arg-test:
	echo "SHOW ME: ${INT} ${STR}"