package service

import (
	"A-MATH/game/components"
	"A-MATH/game/utils"
)

func CreateChip(value string) (components.Chip, error) {
	e := utils.ValidateChip(value)
	if e != nil {
		return components.Chip{}, e
	}
	return components.NewChip(value), nil
}

func CreateSquare(co string) (components.Square, error) {
	coordinate, e := utils.ValidateSquarePosition(co)
	if e != nil {
		return components.Square{}, e
	}
	return components.NewSquare(coordinate), nil
}
