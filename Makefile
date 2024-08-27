build: tidy
	go build -o bin/bc-helper src/main.go

build-versioner: tidy
	go build -o bin/bc-versioner scripts/versioning.go

build-version-control: tidy
	go build -o bin/bc-version-control scripts/vc.go

run: build
	./bin/bc-helper

run-versioner: build-versioner
	./bin/bc-versioner -${CMD} ${ARG} 

run-version-control: build-version-control
	echo "New Version: ${VERSION}"
	./bin/bc-version-control ${VERSION}

init:
	go mod init FraktalDeFiDAO/bc-helper

setup:	init tidy


tidy:
	go mod tidy
	go mod vendor

git-add:
	clear
	set -euo pipefail
	[[ "${FILES}" == "" ]] && sh -c "./bin/bc-version-control -add ." || sh -c "./bin/bc-version-control -add ${FILES}"
	
	
git-commit:
	sh -c "./bin/bc-version-control -commit '${COMMENT}'"

