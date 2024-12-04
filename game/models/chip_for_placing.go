package models

import (
	"A-MATH/game/components"
	"strconv"
	"strings"
)

type ChipForPlacing struct {
	Position     [2]int
	Chip         components.Chip
	SelectedChip components.Chip
}

func NewChipForPlacing(position [2]int, chip components.Chip) ChipForPlacing {
	return ChipForPlacing{position, chip, components.Chip{}}
}

func NewChipForPlacingWithAlternative(position [2]int, chip components.Chip, selectedChip components.Chip) ChipForPlacing {
	return ChipForPlacing{position, chip, selectedChip}
}

func (c ChipForPlacing) String() string {
	var sb strings.Builder
	sb.WriteString("Position: [")
	sb.WriteString(strconv.Itoa(c.Position[0]))
	sb.WriteString(", ")
	sb.WriteString(strconv.Itoa(c.Position[1]))
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
