#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o peerdiscovery ./*.go
docker build -t sadlil/exp-peerdiscovery:1.0 .
minducker sadlil/exp-peerdiscover
kubectl apply -f deployment.yaml
