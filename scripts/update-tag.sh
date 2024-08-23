#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd "${SCRIPT_DIR}/.."

BASE_DIR="${PWD}"

go build -o bin/bc-versioner "${SCRIPT_DIR}/versioning.go" && \
    ./bin/bc-versioner ${@}
