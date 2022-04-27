#! /bin/bash

function build() {

 mkdir -p bin; go mod tidy; make linux
}

echo "starting building binary"
build
echo "binary build complete."