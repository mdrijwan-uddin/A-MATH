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

// For mapping for calculating
func NewChipForCalculating(value string, score int, chipType string) Chip {
	return Chip{value, score, chipType}
}

func setScoreandType(value string) (int, string) {
	switch value {
	case
		string(constants.Zero),
		string(constants.One),
		string(constants.Two),
		string(constants.Three):
		return 1, string(constants.OneDigitNumberType)
		// "0" "1" "2" "3" = 1 Point
	case
		string(constants.Four),
		string(constants.Five),
		string(constants.Six),
		string(constants.Seven),
		string(constants.Eight),
		string(constants.Nine):
		return 2, string(constants.OneDigitNumberType)
		// "5" "6" "7" "8" "9" = 2 Points
	case
		string(constants.Ten),
		string(constants.Twelve):
		return 3, string(constants.TwoDigitNumberType)
		// "10" "12" = 3 Points
	case
		string(constants.Eleven),
		string(constants.Fourteen),
		string(constants.Fifteen),
		string(constants.Sixteen),
		string(constants.Eighteen):
		return 4, string(constants.TwoDigitNumberType)
		// "11" "14" "15" "16" "18" = 4 Points
	case
		string(constants.Thirteen),
		string(constants.Seventeen):
		return 6, string(constants.TwoDigitNumberType)
		// "13" "17" = 6 Points
	case string(constants.Nineteen):
		return 7, string(constants.TwoDigitNumberType)
		// "19" = 7 Points
	case string(constants.Twenty):
		return 5, string(constants.TwoDigitNumberType)
		// "20" = 5 Points
	case
		string(constants.Addition),
		string(constants.Subtraction),
		string(constants.Multiply),
		string(constants.Division):
		return 2, string(constants.OperatorType)
		// "+" "-" "*" "/" = 2 Points
	case
		string(constants.Add_sub),
		string(constants.Multi_divide):
		return 1, string(constants.AlterOperatorType)
		// "+-" "*/" = 1 Point
	case string(constants.Equal):
		return 1, string(constants.EqualType)
		// "=" = 1 Point
	default:
		return 0, string(constants.BlankType)
		// "~" = 0 Point
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
