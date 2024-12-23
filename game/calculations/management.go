package calculations

import (
	"A-MATH/err"
	"A-MATH/game/constants"
	"A-MATH/game/models"
)

func Management(chipsForCalculationSet [][]models.ChipForCalculating) (int, error) {
	var totalScore int

	for _, chipsForCalculation := range chipsForCalculationSet {
		if EquationSeperation(chipsForCalculation) {
			score := Score(chipsForCalculation)
			totalScore += score
		} else {
			return -1, err.New(int(constants.BadRequest), string(constants.InvalidEquationFormed))
		}
	}
	return totalScore, nil
}
