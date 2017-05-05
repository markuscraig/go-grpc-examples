package main

import (
	"log"
	"net"

	pb "github.com/markuscraig/go-grpc-examples/howdy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9999"
)

// server implementation
type server struct{}

// server SayHowdy api implementation
func (s *server) SayHowdy(ctx context.Context, req *pb.HowdyRequest) (*pb.HowdyReply, error) {
	return &pb.HowdyReply{
		Msg: "Howdy " + req.Name + "!",
	}, nil
}

func main() {
	// listen on tcp port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create the server
	s := grpc.NewServer()

	// register the server api
	pb.RegisterGreeterServer(s, &server{})

	// register reflection service on the grpc server
	reflection.Register(s)

	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
