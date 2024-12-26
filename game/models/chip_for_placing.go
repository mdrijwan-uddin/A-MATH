package models

import (
	"A-MATH/game/components"
	"strconv"
	"strings"
)

type ChipForPlacing struct {
	Coordinate   [2]int
	Chip         components.Chip
	SelectedChip components.Chip
}

func NewChipForPlacing(coord [2]int, chip components.Chip) ChipForPlacing {
	return ChipForPlacing{coord, chip, components.Chip{}}
}

func NewChipForPlacingWithAlternative(coord [2]int, chip components.Chip, selectedChip components.Chip) ChipForPlacing {
	return ChipForPlacing{coord, chip, selectedChip}
}

func (c ChipForPlacing) String() string {
	var sb strings.Builder
	sb.WriteString("Coordinate: [")
	sb.WriteString(strconv.Itoa(c.Coordinate[0]))
	sb.WriteString(", ")
	sb.WriteString(strconv.Itoa(c.Coordinate[1]))
	sb.WriteString("],\tChip: {")
	sb.WriteString(c.Chip.String())
	sb.WriteString("}")
	if !c.SelectedChip.IsEmpty() {
		sb.WriteString(",\tSelectedChip: {")
		sb.WriteString(c.SelectedChip.String())
		sb.WriteString("}")
	}

	sb.WriteString("\n")
	return sb.String()
}
