package state

import "time"

type Item struct {
	ID        int
	Msg       string
	Busy      bool
	BusyUntil time.Time
}
