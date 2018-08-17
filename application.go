package main

import (
	"fmt"
	"github.com/just1689/fun-with-chan/state"
	"time"
)

func main() {

	fmt.Println("Starting")

	topic := state.NewTopic("WORK")


	createConsumer(topic, "100")
	createConsumer(topic, "200")




	go func() {

		for i := 1; i <= 10; i++ {
			msg := fmt.Sprint(i)
			//fmt.Println("Writing: ", msg)
			topic.PutItem(msg)
			time.Sleep(10 * time.Millisecond)
		}
	}()


	time.Sleep(2 * time.Second)
}

func createConsumer(topic *state.Topic, ID string) {
	c := topic.Subscribe(ID)
	go func() {
		for item := range c {
			fmt.Println("Message from ", ID, " says: ", item.Msg)
			topic.CompletedItem(state.DoneMessage{ConsumerID: ID, ItemID: item.ID})

		}
	}()

}
