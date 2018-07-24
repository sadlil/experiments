package main

import (
	"log"
	"net"
	"net/http"

	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sadlil/experiments/grpc-examples/_proto"
	"github.com/sadlil/experiments/grpc-examples/pkg"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)

	// We first match the connection against HTTP2 fields. If matched, the
	// connection will be sent through the "grpcl" listener.
	grpcl := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	// The next Line Incresses CPU uses.
	// grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// Otherwise, we match it againts HTTP1 methods. If matched,
	// it is sent through the "httpl" listener.
	httpl := m.Match(cmux.Any())

	go ServeGRPC(grpcl)
	go ServeHTTP(httpl)
	m.Serve()
}

func ServeGRPC(l net.Listener) {
	gRPCServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(gRPCServer, pkg.NewEchoServer())
	gRPCServer.Serve(l)
}

var grpcDialOptions = []grpc.DialOption{grpc.WithInsecure()}

func ServeHTTP(l net.Listener) {
	gwMux := gwrt.NewServeMux()
	helloworld.RegisterGreeterHandlerFromEndpoint(context.Background(), gwMux, "localhost:7070", grpcDialOptions)
	srv := &http.Server{
		Addr:    ":1010",
		Handler: gwMux,
	}
	srv.Serve(l)
}
