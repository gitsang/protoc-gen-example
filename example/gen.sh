#!/bin/bash

git_root=$(git rev-parse --show-toplevel)

if [[ ! -f "${git_root}/dist/bin/protoc-gen-example" ]]; then
    pushd "${git_root}" || return
        bash ./build/build.sh
    popd || return
fi

protoc \
    -I . \
    -I "${git_root}/proto" \
    --plugin=protoc-gen-example="${git_root}/dist/bin/protoc-gen-example" \
    --example_out=./pkg/api/v1 \
    hello.proto
