package models

import (
	"A-MATH/game/components"
	"strings"
)

type ChipForCalculating struct {
	SquareType      string
	Chip            components.Chip
	IsPlacedOnBoard bool
}

func NewChipForCalculating(squareType string, chip components.Chip, IsPlacedOnBoard bool) ChipForCalculating {
	return ChipForCalculating{squareType, chip, IsPlacedOnBoard}
}

func (c ChipForCalculating) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString("Type: ")
	sb.WriteString(c.SquareType)
	sb.WriteString(",\tChip: {")
	sb.WriteString(c.Chip.String())
	sb.WriteString("},\t")
	if c.IsPlacedOnBoard {
		sb.WriteString(" (No need to place on board)")
	} else {
		sb.WriteString(" (Need to place on board)")
	}
	return sb.String()
}
