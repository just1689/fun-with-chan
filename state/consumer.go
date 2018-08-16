package state

type Consumer struct {
	ID      int
	Channel chan *Item
	Idle    bool
}
