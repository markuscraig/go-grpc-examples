package main

import (
	"log"

	pb "github.com/markuscraig/go-grpc-examples/howdy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9999"
)

func main() {
	// connect to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// create the client
	c := pb.NewGreeterClient(conn)

	// invoke the api
	req := &pb.HowdyRequest{
		Name: "Mark",
	}
	resp, err := c.SayHowdy(context.Background(), req)
	if err != nil {
		log.Fatalf("could not invoke api: %v", err)
	}

	log.Printf("\nRESPONSE: %s", resp.Msg)
}
