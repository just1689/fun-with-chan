package state

type Consumer struct {
	Channel chan *Item
	Idle    bool
}
