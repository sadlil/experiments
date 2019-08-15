#!/usr/bin/env bash

set -e

start() {
    docker run -d -p 5000:5000 --restart=always --name registry registry:2
}

stop() {
    docker container stop registry
    docker container rm -v registry
}

check() {
    docker pull alpine:latest
    docker tag alpine:latest localhost:5000/alpine:test
    docker push localhost:5000/alpine:test

    docker image remove alpine:latest
    docker image remove localhost:5000/alpine:test

    docker pull localhost:5000/alpine:test

    docker image remove localhost:5000/alpine:test
}

logs() {
    docker logs -f registry
}

RETVAL=0
if [ $# -eq 0 ]; then
  echo "No Target specified"
  exit 1
fi

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  check)
    check
    ;;
  logs)
    logs
    ;;
  *)
    (10)
    echo $"Usage: $0 {up|down}"
    RETVAL=1
    ;;
esac

exit ${RETVAL}