#!/bin/bash

protoc \
    -I ./proto \
    --go_out=paths=source_relative:./proto \
    $(find ./proto -name "*.proto")

mkdir -p ./dist/bin
go build -o ./dist/bin/protoc-gen-example .
