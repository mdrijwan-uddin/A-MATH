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

	isVertical := true
	isHorizontal := true
	firstColumn, firstRow := c[0].position[0], c[0].position[1]

	for _, chip := range c[1:] {
		if chip.position[0] != firstColumn {
			isVertical = false
		}
		if chip.position[1] != firstRow {
			isHorizontal = false
		}
		// If neither vertical nor horizontal alignment is possible, exit early
		if !isVertical && !isHorizontal {
			return false
		}
	}

	return isVertical || isHorizontal
}