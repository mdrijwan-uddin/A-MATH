package models

import "A-MATH/game/components"

type ChipForPlacing struct {
	Position [2]int
	Chip     components.Chip
}

func NewChipForPlacing(position [2]int, chip components.Chip) ChipForPlacing {
	return ChipForPlacing{position, chip}
}
