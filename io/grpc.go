package io

import (
	"github.com/just1689/fun-with-chan/fun"
	"log"
	"github.com/just1689/fun-with-chan/example"
	"io"
	"fmt"
	"net"
	"google.golang.org/grpc"
)

type funServer struct {
}

func (f funServer) Put(i fun.Fun_PutServer) error {

	ctx := i.Context()

	example.Topic.PutItem("XXX")

	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		item, err := i.Recv()
		if err == io.EOF {
			log.Println("Normal exit", err)
			return err
		}
		if err != nil {
			log.Println("We got an error over GRPC! ", err)
			continue
		}
		fmt.Println("Putting item in the queue: ", item.Msg)
		example.Topic.PutItem(item.Msg)
	}

}

func
StartServer() {

	//Listen
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Could not GRPC start server")
	}

	s := grpc.NewServer()
	fun.RegisterFunServer(s, &funServer{})

	if err := s.Serve(li); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
