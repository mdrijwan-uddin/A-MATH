package game

import (
	. "A-MATH/model" //might be change to another folder later and remove the dot
	"slices"
	"strconv"
)

type Chip struct {
	Value    string
	Score    int
	ChipType string
}

func NewChip(value string) Chip {
	score := setChipScore(value)
	chipType := setChipType(value)
	return Chip{value, score, chipType}
}

func setChipScore(values string) int {
	onePointSet := []string{string(Zero), string(One), string(Two), string(Three), string(Add_sub),
		string(Multi_divide), string(Equal)}
	twoPointSet := []string{string(Four), string(Five), string(Six), string(Seven), string(Eight),
		string(Nine), string(Addition), string(Subtraction), string(Multiply), string(Division)}
	threePointSet := []string{string(Ten), string(Twelve)}
	fourPointSet := []string{string(Eleven), string(Fourteen), string(Fifteen), string(Sixteen),
		string(Eighteen)}
	fivePointSet := []string{string(Twenty)}
	SixPointSet := []string{string(Thirteen), string(Seventeen)}
	sevenPointSet := []string{string(Nineteen)}

	if slices.Contains(onePointSet, values) {
		return int(OnePoint)
	} else if slices.Contains(twoPointSet, values) {
		return int(TwoPoint)
	} else if slices.Contains(threePointSet, values) {
		return int(ThreePoint)
	} else if slices.Contains(fourPointSet, values) {
		return int(FourPoint)
	} else if slices.Contains(fivePointSet, values) {
		return int(FivePoint)
	} else if slices.Contains(SixPointSet, values) {
		return int(SixPoint)
	} else if slices.Contains(sevenPointSet, values) {
		return int(SevenPoint)
	} else {
		return int(ZeroPoint)
	}
}

func setChipType(values string) string {
	oneDigitSet := []string{string(Zero), string(One), string(Two), string(Three), string(Four),
		string(Five), string(Six), string(Seven), string(Eight), string(Nine)}
	twoDigitSet := []string{string(Ten), string(Eleven), string(Twelve), string(Thirteen),
		string(Fourteen), string(Fifteen), string(Sixteen), string(Seventeen), string(Eighteen),
		string(Nineteen), string(Twenty)}
	operatorSet := []string{string(Addition), string(Subtraction), string(Multiply),
		string(Division), string(Add_sub), string(Multi_divide)}
	equalSet := []string{string(Equal)}

	if slices.Contains(oneDigitSet, values) {
		return string(OneDigitNumberType)
	} else if slices.Contains(twoDigitSet, values) {
		return string(TwoDigitNumberType)
	} else if slices.Contains(operatorSet, values) {
		return string(OperatorType)
	} else if slices.Contains(equalSet, values) {
		return string(EqualType)
	} else {
		return string(BlankType)
	}
}

func (chip Chip) String() string {
	return "Value: " + chip.Value +
		"\tScore: " + strconv.Itoa(chip.Score) +
		"\tType: " + chip.ChipType
}
