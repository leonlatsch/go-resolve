#!/bin/bash

if [[ "$1" == "docker" ]]; then
    echo "Building docker image"
    docker build -t ghcr.io/leonlatsch/go-resolve . --target production
else
    echo "Cleaning up"
    rm -rf build/
    mkdir build/

    echo "Building executable"
    go build -o build/go-resolve
fi
