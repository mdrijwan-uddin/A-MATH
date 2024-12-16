package rules

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
)

func ValidateMathameticPrinciple(chipSet []components.Chip) {

	if !haveEqualsInEquation(chipSet) {
		return // add error
	}

	if !isChipsLenghtReasonably(chipSet) {
		return // add error
	}

	if isOperationPlaceOnUnusualLocation(chipSet) {
		return // add error
	}

	if areOperationNextToEachOther(chipSet) {
		return // add error
	}

	if areTwoDigitNumbersStackedWithTheOtherNumber(chipSet) {
		return // add error
	}

	if !areOneDigitNumbersReasonablyStacked(chipSet) {
		return // add error
	}

	if isZeroAtTheBeginWithNumberFormed(chipSet) {
		return // add error
	}

	if isDivideByZero(chipSet) {
		return // add error
	}
}

// lenght of calculating chip should be between 3 to 15
func isChipsLenghtReasonably(chips []components.Chip) bool {
	return 3 <= len(chips) && len(chips) <= constants.BoardRange
}

// Need atlease 1 Equal Chip in the equations
func haveEqualsInEquation(chips []components.Chip) bool {
	for _, chip := range chips {
		if chip.Value == string(constants.Equal) {
			return true
		}
	}
	return false
}

// Operation-type Chip in the equations shouldn't be place at the beginning and the end of the equation (except Subtraction-sign Chip)
func isOperationPlaceOnUnusualLocation(chips []components.Chip) bool {
	firstChip := chips[0]
	lastChip := chips[len(chips)-1]

	return (isOperationType(firstChip) && firstChip.Value != string(constants.Subtraction)) ||
		isOperationType(lastChip)
}

// Operation-type Chip in the equations shouldn't be place at next to the each other (except Equal Chip and Subtraction-sign Chip)
func areOperationNextToEachOther(chips []components.Chip) bool {
	for i := 0; i < len(chips)-1; i++ {
		currentChip := chips[i]
		nextChip := chips[i+1]

		if isOperationType(currentChip) && isOperationType(nextChip) {
			if currentChip.Value == string(constants.Equal) || nextChip.Value == string(constants.Subtraction) {
				continue
			}
			return true
		}
	}
	return false
}

// Helper function to check if the chip is of EqualType or OperatorType
func isOperationType(chip components.Chip) bool {
	return chip.ChipType == string(constants.EqualType) || chip.ChipType == string(constants.OperatorType)
}

// Two-Digit-Number-type Chip in the equations shouldn't be place at next to the other number
func areTwoDigitNumbersStackedWithTheOtherNumber(chips []components.Chip) bool {
	for i := 0; i < len(chips)-1; i++ {
		currentChip := chips[i]
		nextChip := chips[i+1]

		if currentChip.ChipType == string(constants.TwoDigitNumberType) &&
			nextChip.ChipType == string(constants.TwoDigitNumberType) {
			return true
		}

		if currentChip.ChipType == string(constants.OneDigitNumberType) &&
			nextChip.ChipType == string(constants.TwoDigitNumberType) {
			return true
		}

		if currentChip.ChipType == string(constants.TwoDigitNumberType) &&
			nextChip.ChipType == string(constants.OneDigitNumberType) {
			return true
		}
	}
	return false
}

// One-Digit-Number-type Chip in the equations are allow to be place at next to the each other for maximum 3 stacks (Become three digit number)
func areOneDigitNumbersReasonablyStacked(chips []components.Chip) bool {
	counter := 0
	for _, chip := range chips {
		if chip.ChipType == string(constants.OneDigitNumberType) {
			counter++
		} else {
			counter = 0
		}

		if counter > 3 {
			return false
		}
	}
	return true
}

// Zero Chip in the equations aren't allow to be place at next to the each other for maximum 3 stacks (Become three digit number)
func isZeroAtTheBeginWithNumberFormed(chips []components.Chip) bool {

	for i := 0; i < len(chips)-1; i++ {
		previousChip := chips[i-1]
		currentChip := chips[i]
		nextChip := chips[i+1]

		if currentChip.Value == string(constants.Zero) &&
			nextChip.ChipType == string(constants.OneDigitNumberType) {
			if i == 0 {
				return true
			}

			if previousChip.ChipType != string(constants.OneDigitNumberType) {
				return true
			}
		}
	}

	return false
}

// The Equation with divide by zero is not allowed
func isDivideByZero(chips []components.Chip) bool {

	for i := 0; i < len(chips)-1; i++ {
		currentChip := chips[i]
		nextChip := chips[i+1]

		if currentChip.Value == string(constants.Division) &&
			nextChip.Value == string(constants.Zero) {
			return true
		}
	}

	return false
}
