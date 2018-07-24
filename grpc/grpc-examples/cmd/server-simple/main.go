package main

import (
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sadlil/experiments/grpc-examples/_proto"
	"github.com/sadlil/experiments/grpc-examples/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	go runGRPCServer()
	runProxy()
}

func runGRPCServer() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, pkg.NewEchoServer())
	s.Serve(l)
}

func runProxy() {
	httpMux := http.NewServeMux()
	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := helloworld.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, "localhost:9090", dialOpts)
	if err != nil {
		log.Fatal(err)
	}

	httpMux.Handle("/", mux)
	http.ListenAndServe(":9191", mux)
}
