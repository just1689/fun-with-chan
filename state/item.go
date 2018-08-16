package state

import "time"

type Item struct {
	ID        int64
	Msg       string
	Busy      bool
	BusyUntil time.Time
}
