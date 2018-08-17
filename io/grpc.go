package io

import (
	"github.com/just1689/fun-with-chan/fun"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"log"
)

type SubsciberServer struct {
}

func (s SubsciberServer) SaySubscribe(ctx context.Context, f *fun.SubscribeRequest) (*fun.Void, error) {
	go func() {

	}()
	return nil, nil
}

func StartServer() {
	server := grpc.NewServer()
	var subscribers SubsciberServer
	fun.RegisterSubscribeServer(server, subscribers)
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Could not GRPC start server")
	}
	server.Serve(li)

}
