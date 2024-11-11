package utils

import (
	"A-MATH/constants"
	"slices"

	"A-MATH/err"
	"strconv"
)

var ChipSet = []string{string(constants.Zero), string(constants.One), string(constants.Two), string(constants.Three), string(constants.Four),
	string(constants.Five), string(constants.Six), string(constants.Seven), string(constants.Eight), string(constants.Nine), string(constants.Ten),
	string(constants.Eleven), string(constants.Twelve), string(constants.Thirteen), string(constants.Fourteen), string(constants.Fifteen),
	string(constants.Sixteen), string(constants.Seventeen), string(constants.Eighteen), string(constants.Nineteen), string(constants.Twenty),
	string(constants.Addition), string(constants.Subtraction), string(constants.Multiply), string(constants.Division),
	string(constants.Add_sub), string(constants.Multi_divide), string(constants.Equal), string(constants.Blank)}

func ValidateSquarePosition(pos string) ([2]int, error) {
	if len(pos) != 2 && len(pos) != 3 {
		return [2]int{}, err.New(int(constants.BadRequest), string(constants.InvalidInputForBoard))
	}

	xAxis := int(pos[0] - 64)
	if xAxis < 0 || xAxis > 15 {
		return [2]int{}, err.New(int(constants.BadRequest), string(constants.InvalidInputForBoard))
	}

	yAxis, e := strconv.Atoi(pos[1:])
	if e != nil {
		return [2]int{}, err.New(int(constants.BadRequest), string(constants.InvalidInputForBoard))
	}
	if yAxis < 0 || yAxis > 15 {
		return [2]int{}, err.New(int(constants.BadRequest), string(constants.InvalidInputForBoard))
	}

	return [2]int{xAxis, yAxis}, nil
}

func ValidateChip(values string) error {
	if slices.Contains(ChipSet, values) {
		return err.New(int(constants.BadRequest), string(constants.InvalidInputForChip))
	} else {
		return nil
	}
}
