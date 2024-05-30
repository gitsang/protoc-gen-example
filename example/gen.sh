#!/bin/bash

pushd ..
    bash ./build/build.sh
popd

protoc \
    -I . \
    -I ../proto \
    --plugin=protoc-gen-example=../dist/bin/protoc-gen-example \
    --example_out=. \
    hello.proto
