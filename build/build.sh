#!/bin/bash

mkdir -p ./dist/bin
go build -o ./dist/bin/protoc-gen-example .
