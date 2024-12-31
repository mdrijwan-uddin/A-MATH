package calculations

import (
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"fmt"
)

func Score(chipsForCalculation []models.ChipForCalculating) int {
	var currentScore = 0
	var multiplier = 1

	for _, chips := range chipsForCalculation {
		score, times := identifySquareType(chips)
		fmt.Println(times)
		currentScore += score
		multiplier *= times
	}

	return currentScore * multiplier
}

func identifySquareType(c models.ChipForCalculating) (int, int) {
	if c.IsPlacedOnBoard {
		return c.Chip.Score, 1
	}

	switch c.SquareType {
	case string(constants.RedSquare):
		return c.Chip.Score, 3
	case string(constants.YellowSquare):
		return c.Chip.Score, 2
	case string(constants.BlueSquare), string(constants.CenterSquare):
		return c.Chip.Score * 3, 1
	case string(constants.OrangeSquare):
		return c.Chip.Score * 2, 1
	default:
		return c.Chip.Score, 1
	}
}
