package components

import (
	"A-MATH/game/constants"
	"strings"
)

type Rack struct {
	Chips [constants.MaxChipInRack]Chip
}

func NewRack() Rack {
	return Rack{[constants.MaxChipInRack]Chip{}}
}

func (r *Rack) Add(c Chip) {
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

func (r *Rack) Remove(c Chip) {
	if r.IsEmpty() {
		return
	}

	emptyChip := Chip{}
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

func (r Rack) GetTotalChip() int {
	for i, ch := range r.Chips {
		if ch.IsEmpty() {
			return i
		}
	}
	return 8
}

// need to discuss
func (r Rack) IsChipFound(c Chip) bool {
	for _, ch := range r.Chips {
		if ch.IsEmpty() {
			return false
		}
		if ch == c {
			return true
		}
	}
	return false
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
	return r.Chips == [constants.MaxChipInRack]Chip{}
}

func (r Rack) String() string {
	var sb strings.Builder
	sb.WriteString("Rack: {")
	for i := range len(r.Chips) {
		sb.WriteString("[")
		sb.WriteString(r.Chips[i].Value)
		sb.WriteString("]")
	}
	sb.WriteString("}")
	return sb.String()
}
