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
	var score int
	var times int
	var delta int

	if c.IsPlacedOnBoard {
		delta = 1
	} else {
		delta = 0
	}

	switch c.SquareType {
	case string(constants.RedSquare):
		score = c.Chip.Score
		times = (3 - (delta * 2))
	case string(constants.YellowSquare):
		score = c.Chip.Score
		times = (2 - delta)
	case string(constants.BlueSquare), string(constants.CenterSquare):
		score = c.Chip.Score * (3 - (delta * 2))
		times = 1
	case string(constants.OrangeSquare):
		score = c.Chip.Score * (2 - delta)
		times = 1
	default:
		score = c.Chip.Score
		times = 1
	}

	return score, times
}
