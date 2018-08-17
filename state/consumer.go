package state

type consumer struct {
	id      string
	channel chan *Item
	idle    bool
}
