package main

import (
	"google.golang.org/grpc"
	"log"
	"time"
	"golang.org/x/net/context"
	"github.com/just1689/fun-with-chan/fun"
	"fmt"
)

const (
	address = "localhost:8000"
	msg     = "Swag for yolo!"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := fun.NewFunClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	re, err := c.Put(ctx, &fun.PutMessage{Topic: "Le queue", Msg: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(re)
	time.Sleep(2 * time.Second)
	log.Printf("Put complete!")
}
