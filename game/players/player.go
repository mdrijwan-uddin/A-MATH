package players

import (
	"strconv"
	"strings"
)

type Player struct {
	ID   int
	Name string
}

func NewPlayer(ID int, Name string) Player {
	return Player{ID, Name}
}

func (p Player) String() string {
	var sb strings.Builder
	sb.WriteString("{Player ID: ")
	sb.WriteString(strconv.Itoa(p.ID))
	sb.WriteString(", Name: ")
	sb.WriteString(p.Name)
	sb.WriteString("}")
	return sb.String()
}
