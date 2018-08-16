package state

import "container/ring"

func find(head *ring.Ring, ID int64) *ring.Ring {
	var r *ring.Ring
	r = head
	item := r.Value.(*Item)
	headID := item.ID
	found := item.ID == ID
	if found {
		return r
	}

	for !found {
		r = r.Next()
		if r == nil {
			return nil
		}
		item := r.Value.(*Item)
		found = item.ID == ID
		if found {
			return r
		}
		if item.ID == headID {
			return nil
		}
	}
	return nil

}
