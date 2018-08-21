package main

import (
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"github.com/just1689/fun-with-chan/fun"
)

const (
	address = "localhost:8000"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := fun.NewFunClient(conn)
	stream, err := c.Put(context.Background())
	if err != nil {
		log.Fatalf("did not connect to stream: %v", err)
	}

	stream.Send(&fun.PutMessage{Topic: "Le queue", Msg: "Just one message"})
	stream.CloseSend()

	log.Printf("Put complete!")
}
