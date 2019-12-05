package main

import (
	"context"
	pb "github.com/ishankhare07/grpc-test/pkg/services/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port = ":8888"

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("New request from %s", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf(err.Error())
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf(err.Error())
	}
}
