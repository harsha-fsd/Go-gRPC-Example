package main

import (
	"log"
	"net"

	pb "github.com/harsha-fsd/Go-gRPC-Example/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "127.0.0.1:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)

	}

}
