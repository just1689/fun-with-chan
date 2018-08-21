package state

import (
	"container/ring"
	"fmt"
	"time"
)

type Topic struct {
	Name              string
	head              *ring.Ring
	count             int
	countID           int64
	incomingChan      chan string
	completedChan     chan DoneMessage
	consumers         []consumer
	consumerInc       int
	incomingConsumers chan consumer
	hasTimeout        bool
	timeout           int
}

func NewTopic(config TopicConfig) *Topic {
	t := Topic{Name: config.Name, count: 0, countID: 0, consumerInc: 0, hasTimeout: config.TimeoutSeconds > 0, timeout: config.TimeoutSeconds}
	t.incomingChan = make(chan string, 5)
	t.completedChan = make(chan DoneMessage, 5)
	t.incomingConsumers = make(chan consumer, 5)
	t.manageIO()
	return &t
}

type TopicConfig struct {
	Name           string
	TimeoutSeconds int
}

func (t *Topic) manageIO() {
	go func() {
		for {
			select {
			case c := <-t.incomingConsumers:
				t.handleConsumer(c)
				break
			case message := <-t.completedChan:
				t.handleDone(message)
				break
			case in := <-t.incomingChan:
				t.handleIn(in)
				break
			}
		}
	}()
}

func (t *Topic) PutItem(msg string) {
	t.incomingChan <- msg
}

func (t *Topic) CompletedItem(message DoneMessage) {
	t.completedChan <- message
}
func (t *Topic) Subscribe(ID string) chan *Item {
	t.consumerInc++
	consumer := consumer{idle: true, id: ID}
	consumer.channel = make(chan *Item)
	t.consumers = append(t.consumers, consumer)
	return consumer.channel
}

func (t *Topic) handleConsumer(c consumer) {
	t.consumers = append(t.consumers, c)

}

func (t *Topic) handleIn(msg string) {

	t.count++

	if t.count == 1 {
		r := ring.New(1)
		t.head = r
		t.head.Value = newItem(t, &msg)
		return
	}

	r := ring.New(1)
	r.Value = newItem(t, &msg)
	r.Link(t.head)

	t.work()

}

func (t *Topic) canWork() bool {

	if t.count == 0 {
		return false
	}

	if t.consumers == nil {
		fmt.Println("CW: consumers nil")
		return false
	}

	anyIdle := false
	for _, c := range t.consumers {
		if c.idle == true {
			anyIdle = true
			break
		}
	}
	if !anyIdle {
		return false
	}

	return (t.head.Value.(*Item)).Busy == false

}

func (t *Topic) work() int {
	worked := 0

	if !t.canWork() {
		return worked
	}

	for _, consumer := range t.consumers {

		item := t.findFirstAvailMsg()
		if item == nil {
			return worked
		}

		if consumer.idle {
			fmt.Println("->", item.Msg, " to consumer", consumer.id)
			consumer.channel <- item
			item.Busy = true
			consumer.idle = false
			worked++
			continue

		} else {
			fmt.Println("Worker: ", consumer.id, " was idle? ", consumer.idle)
		}

	}

	return worked

}

func (t *Topic) findFirstAvailMsg() *Item {
	ok := true
	r := t.head
	count := 0
	var item *Item
	for ok {
		item = r.Value.(*Item)
		if !item.Busy {
			return item
		}
		if item.Busy && item.BookedUntil.Before(time.Now()) {
			item.Busy = false
			return item
		}

		if count > r.Len() {
			ok = false
		}
		r = r.Next()
		count++
	}
	return nil
}

func (t *Topic) handleDone(message DoneMessage) {
	r := find(t.head, message.ItemID)

	n := t.head.Next()

	removed := r.Prev().Unlink(1)

	if t.head == removed {
		t.head = n
	}

	for _, c := range t.consumers {
		if c.id == message.ConsumerID {
			c.idle = true
			break
		}
	}

	t.count--

	t.work()
}
