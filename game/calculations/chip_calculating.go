package calculations

import (
	"A-MATH/err"
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"A-MATH/game/utils"
	"fmt"
	"strconv"
)

type numberForCalculation struct {
	Integer  int
	Fraction Fraction
}

func Management(chipsForCalculationSet [][]models.ChipForCalculating) (int, error) {
	var totalScore int

	for _, chipsForCalculation := range chipsForCalculationSet {
		if chipSeperation(chipsForCalculation) {
			score := scoring(chipsForCalculation)
			totalScore += score
		} else {
			return -1, err.New(int(constants.BadRequest), string(constants.InvalidEquationFormed))
		}
	}
	return totalScore, nil
}

func chipSeperation(chipsForCalculation []models.ChipForCalculating) bool {
	var (
		index           = 0
		valueForCalSet  [][]components.Chip
		currentCalGroup []components.Chip
	)

	for i, chips := range chipsForCalculation {
		currentChip := chips.ChipForCalculating.Value

		if currentChip == string(constants.Equal) {
			for j := index; j < i; j++ {
				currentCalGroup = append(currentCalGroup, chipsForCalculation[j].ChipForCalculating)
			}
			valueForCalSet = append(valueForCalSet, currentCalGroup)
			index = i + 1
			currentCalGroup = []components.Chip{} // Reset placeholder
		}
	}

	for j := index; j < len(chipsForCalculation); j++ {
		currentCalGroup = append(currentCalGroup, chipsForCalculation[j].ChipForCalculating)
	}
	valueForCalSet = append(valueForCalSet, currentCalGroup)

	// Process the collected groups of chips
	result := processCalculating(valueForCalSet)

	// Debug output
	fmt.Println("Result:", processCalculating(valueForCalSet))

	return result
}

func processCalculating(chipSets [][]components.Chip) bool {
	allNumbers, allOperators := parseChipSets(chipSets)

	// Debug output
	fmt.Println("Parsed Numbers:", allNumbers)
	fmt.Println("Parsed Operators:", allOperators)

	for i := 0; i < len(allOperators); i++ {
		// Process multiplication and division first
		allNumbers[i], allOperators[i] = processOperators(allNumbers[i], allOperators[i],
			[]string{
				string(constants.Multiply),
				string(constants.Division),
			})

		// Process addition and subtraction
		allNumbers[i], allOperators[i] = processOperators(allNumbers[i], allOperators[i],
			[]string{
				string(constants.Addition),
				string(constants.Subtraction),
			})
	}

	// Debug output
	fmt.Println("Parsed Numbers:", allNumbers)
	fmt.Println("Parsed Operators:", allOperators)

	return processComparition(allNumbers)
}

func processComparition(allNumbers [][]numberForCalculation) bool {
	for i := 0; i < len(allNumbers)-1; i++ {
		if allNumbers[i][0] != allNumbers[i+1][0] {
			return false
		}
	}
	return true
}

func processOperators(numbers []numberForCalculation, operators []string, targetOperators []string) ([]numberForCalculation, []string) {
	j := 0
	for j < len(operators) {

		if len(operators) <= 0 {
			break
		}

		if contains(targetOperators, operators[j]) {
			numbers[j] = mathematicOperating(numbers[j], operators[j], numbers[j+1])

			numberTemp, _ := utils.RemoveSlideElement(numbers, j+1)
			numbers = numberTemp

			OperatorTemp, _ := utils.RemoveSlideElement(operators, j)
			operators = OperatorTemp

		} else {
			j++
		}

	}
	return numbers, operators
}

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func parseChipSets(chipSets [][]components.Chip) ([][]numberForCalculation, [][]string) {
	var allNumbers [][]numberForCalculation
	var allOperators [][]string

	for _, chipSet := range chipSets {
		numbersForCurrentSet, operatorsForCurrentSet := parseSingleChipSet(chipSet)
		allNumbers = append(allNumbers, numbersForCurrentSet)
		allOperators = append(allOperators, operatorsForCurrentSet)
	}

	return allNumbers, allOperators
}

func parseSingleChipSet(chipSet []components.Chip) ([]numberForCalculation, []string) {
	var numbersForCurrentSet []numberForCalculation
	var operatorsForCurrentSet []string
	currentNumber := 0
	isNegative := false

	for i, chip := range chipSet {
		if i == 0 && chip.Value == string(constants.Subtraction) {
			isNegative = true
		}

		currentNumber = processChip(chip, currentNumber, &isNegative, &numbersForCurrentSet, &operatorsForCurrentSet)
	}

	// Add the last number in the current set
	numbersForCurrentSet = append(numbersForCurrentSet, numberForCalculation{currentNumber, Fraction{}})

	return numbersForCurrentSet, operatorsForCurrentSet
}

func processChip(
	chip components.Chip,
	currentNumber int,
	isNegative *bool,
	numbersForCurrentSet *[]numberForCalculation,
	operatorsForCurrentSet *[]string,
) int {

	switch chip.ChipType {
	case string(constants.OneDigitNumberType):
		digitValue, _ := strconv.Atoi(chip.Value)
		currentNumber = (currentNumber * 10) + digitValue

	case string(constants.TwoDigitNumberType):
		currentNumber, _ = strconv.Atoi(chip.Value)

	case string(constants.OperatorType):
		if *isNegative {
			currentNumber *= -1
			*isNegative = false
		}
		*numbersForCurrentSet = append(*numbersForCurrentSet, numberForCalculation{currentNumber, Fraction{}})
		*operatorsForCurrentSet = append(*operatorsForCurrentSet, chip.Value)
		currentNumber = 0
	}

	return currentNumber
}

func mathematicOperating(firstNum numberForCalculation, sign string, secondNum numberForCalculation) numberForCalculation {

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

	return numberForCalculation{}
}

func additionOperating(first, second numberForCalculation) numberForCalculation {
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

	return numberForCalculation{resultNum, resultFraction}
}

func subtractionOperating(first, second numberForCalculation) numberForCalculation {
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

	return numberForCalculation{resultNum, resultFraction}
}

func multiplicationOperating(first, second numberForCalculation) numberForCalculation {
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

	return numberForCalculation{resultNum, resultFraction}
}

func divisionOperating(first, second numberForCalculation) numberForCalculation {
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

	return numberForCalculation{resultNum, resultFraction}
}