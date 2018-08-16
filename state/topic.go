package state

import (
	"container/ring"
	"fmt"
)

type Topic struct {
	Name              string
	Head              *ring.Ring
	Count             int
	CountID           int64
	Incoming          chan string
	Completed         chan DoneMessage
	Consumer          []Consumer
	consumerInc       int
	IncomingConsumers chan Consumer
}

func NewTopic(name string) *Topic {
	t := Topic{Name: name, Count: 0, CountID: 0, consumerInc: 0}
	t.Incoming = make(chan string, 5)
	t.Completed = make(chan DoneMessage, 5)
	t.IncomingConsumers = make(chan Consumer, 5)
	t.manageIO()
	return &t
}

func (t *Topic) manageIO() {
	go func() {
		for {
			select {
			case c := <-t.IncomingConsumers:
				t.handleConsumer(c)
				break
			case message := <-t.Completed:
				t.handleDone(message)
				break
			case in := <-t.Incoming:
				t.handleIn(in)
				break
			}
		}
	}()
}

func (t *Topic) PutItem(msg string) {
	t.Incoming <- msg
}

func (t *Topic) CompletedItem(message DoneMessage) {
	t.Completed <- message
}
func (t *Topic) Subscribe(ID string) chan *Item {
	t.consumerInc++
	consumer := Consumer{Idle: true, ID: ID}
	consumer.Channel = make(chan *Item)
	t.Consumer = append(t.Consumer, consumer)
	return consumer.Channel
}

func (t *Topic) handleConsumer(c Consumer) {
	t.Consumer = append(t.Consumer, c)

}

func (t *Topic) handleIn(msg string) {

	t.Count++

	if t.Count == 1 {
		r := ring.New(1)
		t.Head = r
		t.Head.Value = &Item{ID: t.CountID, Msg: msg, Busy: false}
		return
	}

	r := ring.New(1)
	r.Value = &Item{ID: t.CountID, Msg: msg, Busy: false}
	r.Link(t.Head)

	t.work(0)

}

func (t *Topic) canWork() bool {

	if t.Count == 0 {
		return false
	}

	if t.Consumer == nil {
		fmt.Println("NW: consumers")
		return false
	}

	anyIdle := false
	for _, c := range t.Consumer {
		if c.Idle == true {
			anyIdle = true
			break
		}
	}
	if !anyIdle {
		return false
	}

	return (t.Head.Value.(*Item)).Busy == false

}
func (t *Topic) work(last int) {
	if !t.canWork() {
		return
	}

	worked := 0

	item := t.Head.Value.(*Item)

	for _, consumer := range t.Consumer {

		if consumer.Idle {
			consumer.Channel <- item
			item.Busy = true
			consumer.Idle = false
			worked++
			break
		}

	}

}

func (t *Topic) handleDone(message DoneMessage) {
	r := find(t.Head, message.ItemID)

	n := t.Head.Next()

	removed := r.Prev().Unlink(1)

	if t.Head == removed {
		t.Head = n
	}

	for _, c := range t.Consumer {
		if c.ID == message.ConsumerID {
			c.Idle = true
			break
		}
	}

	t.Count--

	t.work(0)
}
