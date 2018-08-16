package state

type Consumer struct {
	Channel  chan *Item
	BusyWith *Item
}
