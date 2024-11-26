package models

import "A-MATH/game/components"

type ChipForCalculating struct {
	Position        [2]int
	Chip            components.Chip
	IsPlacedOnBoard bool
}

func NewChipForCalculating(position [2]int, chip components.Chip, IsPlacedOnBoard bool) ChipForCalculating {
	return ChipForCalculating{position, chip, IsPlacedOnBoard}
}
