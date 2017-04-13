package pkg

import (
	"github.com/sadlil/experiments/grpc/server-stream/_proto"
	"fmt"
	"time"
)

// Implements of EchoServiceServer

type greeterServerServer struct{}

func NewEchoServer() helloworld.GreeterServer {
	return new(greeterServerServer)
}

func (s *greeterServerServer) SayHello(req *helloworld.HelloRequest, stream helloworld.Greeter_SayHelloServer) error {
	for i := 0; i < 20; i ++ {
		fmt.Println("Sending Message")
		stream.Send(&helloworld.HelloReply{
			Message: fmt.Sprintf("%s %d", "hello", i+1),
		})
		time.Sleep(time.Second*20)
	}
	return nil
}
