package actions

import (
	"A-MATH/game/components"
)

type ChipForPlacing struct {
	position [2]int
	chip     components.Chip
}

func NewChipForPlacing(position [2]int, chip components.Chip) ChipForPlacing {
	return ChipForPlacing{position, chip}
}

func IsChipPlaceOnCenterSquare(c []ChipForPlacing) bool {
	for _, chip := range c {
		if chip.position == [2]int{8, 8} {
			return true
		}
	}
	return false
}

func IsChipPlaceOnOneLineCorrectly(c []ChipForPlacing) bool {
	if len(c) == 1 {
		return true
	}

	return isChipsOnVertical(c) || isChipsOnHorizontal(c)
}

func isChipsOnVertical(c []ChipForPlacing) bool {
	firstColumn := c[0].position[0]
	for i := 1; i < len(c); i++ {
		if firstColumn != c[i].position[0] {
			return false
		}
	}
	return true
}

func isChipsOnHorizontal(c []ChipForPlacing) bool {
	firstRow := c[0].position[1]
	for i := 1; i < len(c); i++ {
		if firstRow != c[i].position[1] {
			return false
		}
	}
	return true
}
