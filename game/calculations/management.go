package calculations

import (
	"A-MATH/err"
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"A-MATH/game/rules"
)

func Management(chipsForCalculationSet [][]models.ChipForCalculating) (int, error) {
	var totalScore int

	for _, chipsForCalculation := range chipsForCalculationSet {

		var chipSet []components.Chip
		for _, chips := range chipsForCalculation {
			chipSet = append(chipSet, components.NewChip(chips.Chip.Value))
		}

		amathError := rules.ValidateMathameticPrinciple(chipSet)
		if amathError != nil {
			return -1, amathError
		}

		if !ProcessCalculating(chipsForCalculation) {
			return -1, err.New(int(constants.BadRequest), string(constants.InvalidEquationFormed))
		}

		score := Score(chipsForCalculation)
		totalScore += score
	}
	return totalScore, nil
}
