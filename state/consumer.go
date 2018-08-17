package state

type Consumer struct {
	id      string
	channel chan *Item
	idle    bool
}
