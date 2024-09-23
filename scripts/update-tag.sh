#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd "${SCRIPT_DIR}/.."

BASE_DIR="${PWD}"
ARGS="${@}"


./bin/bc-versioner -update-type minor && git tag $(cat "${BASE_DIR}/version") \
    && go build -o bin/bc-version-control "${SCRIPT_DIR}/vc.go" \
    && git push --tags