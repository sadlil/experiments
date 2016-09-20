#!/usr/bin/env bash

CGO_ENABLED=0 go build -o dist/os-signal *.go
docker build -t sadlil/os-signal:demo .
docker push sadlil/os-signal:demo