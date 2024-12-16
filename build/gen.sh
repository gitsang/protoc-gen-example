#!/bin/bash

# shellcheck disable=SC2046
protoc \
    -I ./proto \
    --go_out=paths=source_relative:./proto \
    $(find ./proto -name "*.proto")
