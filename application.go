package main

import (
	"fmt"
	"github.com/just1689/fun-with-chan/state"
	"time"
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

	createConsumer(topic, "A", 11)
	createConsumer(topic, "B", 80)

	time.Sleep(10 * time.Second)
}

func createConsumer(topic *state.Topic, ID string, sleep int) {
	c := topic.Subscribe(ID)
	go func() {
		for item := range c {
			//time.Sleep(time.Duration(sleep) * time.Millisecond)
			if item.BookedUntil.Before(time.Now()) {
				topic.CompletedItem(state.DoneMessage{ConsumerID: ID, ItemID: item.ID})
				continue
			}
			fmt.Println("<-", item.Msg, "says consumer", ID)
			topic.CompletedItem(state.DoneMessage{ConsumerID: ID, ItemID: item.ID})
		}
	}()

}
