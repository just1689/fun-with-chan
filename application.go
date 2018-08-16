package main

import (
	"fmt"
)

var incoming = make(chan string)
var outgoing = make(chan chan string)
var head *Item

func main() {
	fmt.Println("Starting")
	go manageQueue()

}

func manageQueue() {

	for {
		select {
		case i := <-incoming:
			handleIn(i)
			break
		case c := <-outgoing:
			handleOut(c)
		}
	}

}

type Item struct {
	Msg     string
	Next    *Item
	Prev    *Item
	Busy    bool
	Timeout int
}

func handleIn(i string) {

	if head == nil {
		head = &Item{Msg: i, Busy: false}
		return
	}

	if head.Prev == nil {
		//There is one item but only one item
		newItem := Item{Msg: i, Busy: false, Next: head, Prev: head}
		head.Next = &newItem
		head.Prev = &newItem
		return
	}

	//There are already two or more items. Append
	newItem := Item{Msg: i, Busy: false, Next: head, Prev: head.Prev}
	head.Prev.Next = &newItem
	head.Prev = &newItem

}

func handleOut(c chan string) {

}
