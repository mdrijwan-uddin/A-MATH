package game

import (
	"A-MATH/constants"
	"A-MATH/utils"

	"strconv"
)

type chip struct {
	Value    string
	Score    int
	ChipType string
}

func NewChip(value string) (chip, error) {
	e := utils.ValidateChip(value)

	if e != nil {
		return chip{}, e
	}

	score, chipType := setScoreandType(value)
	return chip{value, score, chipType}, nil
}

func setScoreandType(value string) (int, string) {
	switch value {
	case string(constants.Zero), string(constants.One), string(constants.Two), string(constants.Three):
		return 1, string(constants.OneDigitNumberType)
	case string(constants.Four), string(constants.Five), string(constants.Six), string(constants.Seven), string(constants.Eight), string(constants.Nine):
		return 2, string(constants.OneDigitNumberType)
	case string(constants.Ten), string(constants.Twelve):
		return 3, string(constants.TwoDigitNumberType)
	case string(constants.Eleven), string(constants.Fourteen), string(constants.Fifteen), string(constants.Sixteen), string(constants.Eighteen):
		return 4, string(constants.TwoDigitNumberType)
	case string(constants.Thirteen), string(constants.Seventeen):
		return 6, string(constants.TwoDigitNumberType)
	case string(constants.Nineteen):
		return 7, string(constants.TwoDigitNumberType)
	case string(constants.Twenty):
		return 5, string(constants.TwoDigitNumberType)
	case string(constants.Addition), string(constants.Subtraction), string(constants.Multiply), string(constants.Division):
		return 2, string(constants.OperatorType)
	case string(constants.Add_sub), string(constants.Multi_divide):
		return 1, string(constants.AlterOperatorType)
	case string(constants.Equal):
		return 1, string(constants.EqualType)
	default:
		return 0, string(constants.BlankType)
	}
}

func (c chip) IsEmpty() bool {
	return c == chip{}
}

func (c chip) String() string {
	return "Value: " + c.Value + ", " +
		"Score: " + strconv.Itoa(c.Score) + ", " +
		"Type: " + c.ChipType
}
