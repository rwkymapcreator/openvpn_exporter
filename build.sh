#!/bin/sh
docker run --rm -w /src -v $PWD:/src -u $(id -u) -e HOME=/src golang:1.22-bullseye sh -c 'make generate && make build'
