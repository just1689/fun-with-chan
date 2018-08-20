package main

import (
	"fmt"
	"github.com/just1689/fun-with-chan/io"
	"github.com/just1689/fun-with-chan/state"
	"time"
	"sync"
	"github.com/just1689/fun-with-chan/example"
)

var w sync.WaitGroup

func main() {

	fmt.Println("Starting")


	go func() {
		for i := 1; i <= 2; i++ {
			msg := fmt.Sprint(i)
			example.Topic.PutItem(msg)
		}
	}()

	createPrintConsumer(example.Topic, "A")


	io.StartServer()

	w.Add(1)
	w.Wait()

}

func createPrintConsumer(topic *state.Topic, ID string) {
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
