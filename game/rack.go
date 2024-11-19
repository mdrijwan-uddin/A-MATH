package game

import (
	"strconv"
)

const maxChipInRack = 8

type rack struct {
	OwnerId uint
	Chips   [maxChipInRack]chip
}

func NewRack(id uint) rack {
	return rack{id, [maxChipInRack]chip{}}
}

func (r *rack) Add(c chip) {
	if r.IsFull() {
		return
	}

	for i, ch := range r.Chips {
		if ch.IsEmpty() {
			r.Chips[i] = c
			break
		}
	}
}

func (r *rack) Remove(c chip) {
	if r.IsEmpty() {
		return
	}

	emptyChip := chip{}
	index := -1

	for i, ch := range r.Chips {
		if ch == c {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	for j := index; j < len(r.Chips)-1; j++ {
		r.Chips[j] = r.Chips[j+1]
	}

	r.Chips[len(r.Chips)-1] = emptyChip
}

func (r rack) IsFull() bool {
	for _, ch := range r.Chips {
		if ch.IsEmpty() {
			return false
		}
	}
	return true
}

func (r rack) IsEmpty() bool {
	return r.Chips == [maxChipInRack]chip{}
}

func (r rack) String() string {
	var str string
	for i := range len(r.Chips) {
		str += "[" + r.Chips[i].Value + "]"
	}

	return "Rack{" + strconv.Itoa(int(r.OwnerId)) + "} : " + str
}
