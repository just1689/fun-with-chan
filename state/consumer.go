package state

type Consumer struct {
	ID      string
	Channel chan *Item
	Idle    bool
}
