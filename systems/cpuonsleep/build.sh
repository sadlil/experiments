#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o bin/withsleep withsleep/main.go
CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o bin/withoutsleep withoutsleep/main.go

cp bin/withsleep ./withsleep/ws
cp bin/withoutsleep ./withoutsleep/wos

docker build -t sadlil/experiments:cpuonsleep-wos ./withoutsleep/
docker build -t sadlil/experiments:cpuonsleep-ws ./withsleep/

rm ./withsleep/ws
rm ./withoutsleep/wos
rm -rf ./bin/

minducker sadlil/experiments

