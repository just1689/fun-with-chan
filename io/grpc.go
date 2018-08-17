package io

import (
	"github.com/just1689/fun-with-chan/fun"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"log"
)

type FunServer struct {
}

func (f FunServer) Sub(context.Context, *fun.SimpleMessage) (*fun.Void, error) {
	return nil, nil
}
func (f FunServer) Put(context.Context, *fun.PutMessage) (*fun.Void, error) {
	return nil, nil
}
func (f FunServer) Done(context.Context, *fun.SimpleMessage) (*fun.Void, error) {
	return nil, nil
}

func StartServer() {
	server := grpc.NewServer()
	var servers FunServer
	fun.RegisterFunServer(server, servers)
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Could not GRPC start server")
	}
	server.Serve(li)

}
