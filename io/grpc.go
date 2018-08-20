package io

import (
	"github.com/just1689/fun-with-chan/fun"
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/just1689/fun-with-chan/example"
)

type funServer struct {
}

func (f funServer) Sub(context.Context, *fun.SimpleMessage) (*fun.Void, error) {
	return nil, nil
}
func (f funServer) Put(ctx context.Context, item *fun.PutMessage) (*fun.Void, error) {
	example.Topic.PutItem(item.Msg)
	return new(fun.Void), nil
}
func (f funServer) Done(context.Context, *fun.SimpleMessage) (*fun.Void, error) {
	return nil, nil
}
func (f funServer) Push(context.Context, *fun.Item) (*fun.Void, error) {
	return nil, nil
}

func StartServer() {

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
