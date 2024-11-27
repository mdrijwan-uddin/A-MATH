package components

import (
	"A-MATH/game/constants"

	"strconv"
	"strings"
)

type Chip struct {
	Value    string
	Score    int
	ChipType string
}

func NewChip(value string) Chip {
	score, chipType := setScoreandType(value)
	return Chip{value, score, chipType}
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

func (c Chip) IsEmpty() bool {
	return c == Chip{}
}

func (c Chip) String() string {
	var sb strings.Builder
	sb.WriteString("Value: ")
	sb.WriteString(c.Value)
	sb.WriteString(", Score: ")
	sb.WriteString(strconv.Itoa(c.Score))
	sb.WriteString(", Type: ")
	sb.WriteString(c.ChipType)
	return sb.String()
}
