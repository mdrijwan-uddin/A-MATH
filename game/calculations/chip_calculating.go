package calculations

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"fmt"
)

type NumberForCalculation struct {
	Integer  int
	Fraction Fraction
}

func ChipCalculationManagement(chipsForCalculationSet [][]models.ChipForCalculating) {
	// var index int
	var valueForCalSet [][]string
	var valueForCal []string

	for _, chipsForCalculation := range chipsForCalculationSet {
		for _, chips := range chipsForCalculation {
			valueForCal = append(valueForCal, chips.ChipForCalculating.Value)
		}
		valueForCalSet = append(valueForCalSet, valueForCal)
	}

	fmt.Println(valueForCalSet)
}

func ChipCalculationSeperation(chipsForCalculation []models.ChipForCalculating) {
	var index = 0
	var valueForCalSet [][]components.Chip
	var valueForCal []components.Chip

	for i, chips := range chipsForCalculation {
		currentChip := chips.ChipForCalculating.Value
		if currentChip == "=" {
			for j := index; j < i; j++ {
				valueForCal = append(valueForCal, chipsForCalculation[j].ChipForCalculating)
			}
			valueForCalSet = append(valueForCalSet, valueForCal)
			index = i + 1
			valueForCal = []components.Chip{} // Empty chip as placeholder
		}
	}

	for j := index; j < len(chipsForCalculation); j++ {
		valueForCal = append(valueForCal, chipsForCalculation[j].ChipForCalculating)
	}
	valueForCalSet = append(valueForCalSet, valueForCal)

	fmt.Println(valueForCalSet)
}

// func calculating(valueForCalSet [][]string) {
// 	var n = 0
// 	var num [][]NumberForCalculation
// 	var operator [][]string
// 	var isFirstNumberNegative = false

// 	for _, valueForCal := range valueForCalSet {
// 		for i := 0; i < len(valueForCal); i++ {
// 			if i == 0 && valueForCal[i] == string(constants.Subtraction) {
// 				isFirstNumberNegative = true
// 			}

// 			if

// 		}
// 	}
// }

func mathematicOperating(firstNum, secondNum NumberForCalculation, sign string) NumberForCalculation {

	switch sign {
	case string(constants.Addition):
		return additionOperating(firstNum, secondNum)
	case string(constants.Subtraction):
		return subtractionOperating(firstNum, secondNum)
	case string(constants.Multiply):
		return multiplicationOperating(firstNum, secondNum)
	case string(constants.Division):
		return divisionOperating(firstNum, secondNum)
	}

	return NumberForCalculation{}
}

func additionOperating(first, second NumberForCalculation) NumberForCalculation {
	resultNum := 0
	resultFraction := Fraction{}

	handleFraction := func(fraction Fraction) {
		if fraction.Denominator == 1 {
			resultNum = fraction.Numerator
		} else {
			resultFraction = fraction
		}
	}

	// Both inputs have fractions
	if !first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		first.Fraction.AddFractionBy(second.Fraction)
		handleFraction(first.Fraction)
	}

	// First has a fraction, second is an integer
	if !first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		first.Fraction.AddBy(second.Integer)
		handleFraction(first.Fraction)
	}

	// First is an integer, second has a fraction
	if first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		second.Fraction.AddBy(first.Integer)
		handleFraction(second.Fraction)
	}

	// Both inputs are integers
	if first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		resultNum = first.Integer + second.Integer
	}

	return NumberForCalculation{resultNum, resultFraction}
}

func subtractionOperating(first, second NumberForCalculation) NumberForCalculation {
	resultNum := 0
	resultFraction := Fraction{}

	handleFraction := func(fraction Fraction) {
		if fraction.Denominator == 1 {
			resultNum = fraction.Numerator
		} else {
			resultFraction = fraction
		}
	}

	// Both inputs have fractions
	if !first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		first.Fraction.SubtractFractionBy(second.Fraction)
		handleFraction(first.Fraction)
	}

	// First has a fraction, second is an integer
	if !first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		first.Fraction.SubtractBy(second.Integer)
		handleFraction(first.Fraction)
	}

	// First is an integer, second has a fraction
	if first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		second.Fraction.IntegerSubtractFraction(first.Integer)
		handleFraction(second.Fraction)
	}

	// Both inputs are integers
	if first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		resultNum = first.Integer - second.Integer
	}

	return NumberForCalculation{resultNum, resultFraction}
}

func multiplicationOperating(first, second NumberForCalculation) NumberForCalculation {
	resultNum := 0
	resultFraction := Fraction{}

	handleFraction := func(fraction Fraction) {
		if fraction.Denominator == 1 {
			resultNum = fraction.Numerator
		} else {
			resultFraction = fraction
		}
	}

	// Both inputs have fractions
	if !first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		first.Fraction.MultiplyFractionBy(second.Fraction)
		handleFraction(first.Fraction)
	}

	// First has a fraction, second is an integer
	if !first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		first.Fraction.MultiplyBy(second.Integer)
		handleFraction(first.Fraction)
	}

	// First is an integer, second has a fraction
	if first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		second.Fraction.MultiplyBy(first.Integer)
		handleFraction(second.Fraction)
	}

	// Both inputs are integers
	if first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		resultNum = first.Integer * second.Integer
	}

	return NumberForCalculation{resultNum, resultFraction}
}

func divisionOperating(first, second NumberForCalculation) NumberForCalculation {
	resultNum := 0
	resultFraction := Fraction{}

	handleFraction := func(fraction Fraction) {
		if fraction.Denominator == 1 {
			resultNum = fraction.Numerator
		} else {
			resultFraction = fraction
		}
	}

	// Both inputs have fractions
	if !first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		first.Fraction.DivideFractionBy(second.Fraction)
		handleFraction(first.Fraction)
	}

	// First has a fraction, second is an integer
	if !first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		first.Fraction.DivideBy(second.Integer)
		handleFraction(first.Fraction)
	}

	// First is an integer, second has a fraction
	if first.Fraction.IsEmpty() && !second.Fraction.IsEmpty() {
		second.Fraction.IntegerDivideFraction(first.Integer)
		handleFraction(second.Fraction)
	}

	// Both inputs are integers
	if first.Fraction.IsEmpty() && second.Fraction.IsEmpty() {
		newFraction := NewFraction(first.Integer, second.Integer)
		handleFraction(newFraction)
	}

	return NumberForCalculation{resultNum, resultFraction}
}
