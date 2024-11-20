package components

import (
	"strings"
)

const maxChipInRack = 8

type Rack struct {
	Chips [maxChipInRack]chip
}

func NewRack() Rack {
	return Rack{[maxChipInRack]chip{}}
}

func (r *Rack) Add(c chip) {
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

func (r *Rack) Remove(c chip) {
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

func (r Rack) IsFull() bool {
	for _, ch := range r.Chips {
		if ch.IsEmpty() {
			return false
		}
	}
	return true
}

func (r Rack) IsEmpty() bool {
	return r.Chips == [maxChipInRack]chip{}
}

func (r Rack) String() string {
	var sb strings.Builder
	sb.WriteString("Rack: {")
	for i := range len(r.Chips) {
		sb.WriteString("[")
		sb.WriteString(r.Chips[i].Value)
		sb.WriteString("]")
	}
	return sb.String()
}
