#!/bin/sh
docker run --rm -w /src -v $PWD:/src -u $(id -u) -e HOME=/src golang:1.24-bookworm sh -c 'make generate && make clean && make build'
