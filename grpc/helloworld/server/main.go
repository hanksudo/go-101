package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "helloworld/proto/helloworld"

	"google.golang.org/grpc"
)

var gRPCPort = flag.Int("grpc-port", 50051, "The gRPC server port")

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello, " + in.Name}, nil
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *gRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
