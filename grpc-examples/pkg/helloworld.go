package pkg

import (
	"fmt"
	"github.com/sadlil/grpc-examples/_proto"
	"golang.org/x/net/context"
)

// Implements of EchoServiceServer

type greeterServerServer struct{}

func NewEchoServer() helloworld.GreeterServer {
	return new(greeterServerServer)
}

func (s *greeterServerServer) SayHello(ctx context.Context, msg *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello %s!", msg.Name),
	}, nil
}
