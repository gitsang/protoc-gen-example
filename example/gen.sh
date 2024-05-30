#!/bin/bash

pushd ..
    bash ./build.sh
popd

protoc \
    -I . \
    -I ../.. \
    --plugin=protoc-gen-example=../protoc-gen-example \
    --example_out=. \
    hello.proto
