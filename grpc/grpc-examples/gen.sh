#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

RETVAL=0
ROOT=$PWD
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

clean() {
	(find . | grep pb.go | xargs rm) || true
	(find . | grep pb.gw.go | xargs rm) || true
}

gen_proto() {
  if [ $(ls -1 *.proto 2>/dev/null | wc -l) = 0 ]; then
    return
  fi
  rm -rf *.pb.go
  protoc -I /usr/local/include -I . \
         -I ${GOPATH}/src/github.com \
         -I ${GOPATH}/src/github.com/googleapis/googleapis/ \
         --go_out=plugins=grpc:. *.proto
}

gen_gateway_proto() {
  if [ $(ls -1 *.proto 2>/dev/null | wc -l) = 0 ]; then
    return
  fi
  rm -rf *.pb.gw.go
  protoc -I /usr/local/include -I . \
         -I ${GOPATH}/src/github.com \
         -I ${GOPATH}/src/github.com/googleapis/googleapis/ \
         --grpc-gateway_out=logtostderr=true:. *.proto
}

gen_server_protos() {
	echo "Generating server protobuf files"
    for d in */ ; do
      pushd ${d}
      gen_proto
      if dirs=$( ls -1 -F | grep "^v.*/" | tr -d "/" ); then
        for dd in $dirs ; do
          pushd ${dd}
          gen_proto
          popd
        done
      fi
      popd
    done
}

gen_proxy_protos() {
    echo "Generating gateway protobuf files"
    for d in */ ; do
      pushd ${d}
      gen_gateway_proto
      if dirs=$( ls -1 -F | grep "^v.*/" | tr -d "/" ); then
        for dd in $dirs ; do
          pushd ${dd}
          gen_gateway_proto
          popd
        done
      fi
      popd
    done
}

gen_protos() {
  clean
  gen_server_protos
  gen_proxy_protos
}

if [ $# -eq 0 ]; then
	gen_protos
	exit $RETVAL
fi

case "$1" in
    compile)
        compile
        ;;
    server)
		gen_server_protos
		;;
	proxy)
		gen_proxy_protos
		gen_cors_patterns
		;;
	*)  echo $"Usage: $0 {compile|server|proxy|js|swagger|json-schema|all|clean}"
		RETVAL=1
		;;
esac
exit $RETVAL