package state

import "time"

type Item struct {
	ID          int64
	Msg         string
	Busy        bool
	hasTimeout  bool
	BookedUntil time.Time
}

func NewItem(t *Topic, msg *string) *Item {
	item := &Item{ID: t.CountID, Msg: *msg}
	if t.hasTimeout {
		item.hasTimeout = true
		item.BookedUntil = time.Now().Add(time.Duration(t.timeout) * time.Second)
	}
	return item
}
