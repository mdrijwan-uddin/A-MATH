package game

import "strconv"

type Rack struct {
	ownerId uint
	chips   [8]Chip
}

func NewRack(id uint) Rack {
	return Rack{id, [8]Chip{}}
}

func (r Rack) String() string {
	var str string
	for i := range len(r.chips) {
		str += "[" + r.chips[i].Value + "]"
	}

	return "Rack{" + strconv.Itoa(int(r.ownerId)) + "} : " + str
}
