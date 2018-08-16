package main

import (
	"fmt"
	"github.com/just1689/fun-with-chan/state"
	"time"
)

func main() {
	fmt.Println("Starting")

	topic := state.NewTopic("WORK")

	go func() {
		c := topic.Subscribe()
		for s := range c {
			fmt.Println("Reading: ", s.Msg)
			topic.CompletedItem(s.ID)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			msg := fmt.Sprint(i)
			fmt.Println("Writing: ", msg)
			topic.PutItem(msg)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)

}
