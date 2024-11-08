package game

import (
	. "A-MATH/constants"
	"A-MATH/err"
	"slices"
	"strconv"
)

type Chip struct {
	Value    string
	Score    int
	ChipType string
}

func NewChip(value string) (Chip, error) {
	score, e := setChipScore(value)
	if e != nil {
		return Chip{}, e
	}

	chipType := setChipType(value)
	return Chip{value, score, chipType}, nil
}

func setChipScore(values string) (int, error) {
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
	zeroPointSet := []string{string(Blank)}

	if slices.Contains(onePointSet, values) {
		return 1, nil
	} else if slices.Contains(twoPointSet, values) {
		return 2, nil
	} else if slices.Contains(threePointSet, values) {
		return 3, nil
	} else if slices.Contains(fourPointSet, values) {
		return 4, nil
	} else if slices.Contains(fivePointSet, values) {
		return 5, nil
	} else if slices.Contains(SixPointSet, values) {
		return 6, nil
	} else if slices.Contains(sevenPointSet, values) {
		return 7, nil
	} else if slices.Contains(zeroPointSet, values) {
		return 0, nil
	} else {
		return -1, err.New(int(BadRequest), string(InvalidInputForChip))
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

func (c Chip) String() string {
	return "Value: " + c.Value +
		"\tScore: " + strconv.Itoa(c.Score) +
		"\tType: " + c.ChipType
}
