package main

import (
	"fmt"
	"github.com/just1689/fun-with-chan/fun"
	"github.com/just1689/fun-with-chan/state"
	"google.golang.org/grpc"
	"net"
	"time"
	"golang.org/x/net/context"
)

func main() {

	fmt.Println("Starting")

	topicConfig := state.TopicConfig{Name: "Le queue", TimeoutSeconds: 1}

	topic := state.NewTopic(topicConfig)

	go func() {
		for i := 1; i <= 100; i++ {
			msg := fmt.Sprint(i)
			topic.PutItem(msg)
		}
	}()

	createConsumer(topic, "A")
	createConsumer(topic, "B")

	startServer()

}

type SubsciberServer struct {
}

func (s SubsciberServer) SaySubscribe(context.Context, *fun.SubscribeRequest) (*fun.Void, error) {
	return nil, nil
}

func startServer() {
	server := grpc.NewServer()
	var subscribers SubsciberServer
	fun.RegisterSubscribeServer(server, subscribers)
	li, _ := net.Listen("tcp", ":8000")
	server.Serve(li)

}

func createConsumer(topic *state.Topic, ID string) {
	c := topic.Subscribe(ID)
	go func() {
		for item := range c {
			if item.BookedUntil.Before(time.Now()) {
				topic.CompletedItem(state.DoneMessage{ConsumerID: ID, ItemID: item.ID})
				continue
			}
			fmt.Println("<-", item.Msg, "says consumer", ID)
			topic.CompletedItem(state.DoneMessage{ConsumerID: ID, ItemID: item.ID})
		}
	}()

}
