package game

import (
	"fmt"
	"slices"

	"A-MATH/constants"
	"A-MATH/utils"
)

type Square struct {
	Position    [2]int
	SquareType  string
	ChipPlaceOn Chip
}

func NewSquare(pos string) (Square, error) {
	position, e := utils.ValidateSquarePosition(pos)
	if e != nil {
		return Square{}, e
	}

	squareType := setSquareType(position)
	return Square{position, squareType, Chip{}}, nil
}

func setSquareType(pos [2]int) string {
	minimalizePosition := func(pos int) int {
		if pos == 8 || pos == 15 {
			return 1
		} else if pos > 8 {
			return pos - ((pos - 8) * 2)
		}
		return pos
	}

	if pos[0] != 8 && pos[1] != 8 {
		pos[0] = minimalizePosition(pos[0])
		pos[1] = minimalizePosition(pos[1])
	}

	var CenterSquareSet [][2]int
	CenterSquareSet = append(CenterSquareSet, [2]int{8, 8})

	var RedSquareSet [][2]int
	RedSquareSet = append(RedSquareSet, [2]int{1, 1})

	var YellowSquareSet [][2]int
	YellowSquareSet = append(YellowSquareSet, [2]int{2, 2})
	YellowSquareSet = append(YellowSquareSet, [2]int{3, 3})
	YellowSquareSet = append(YellowSquareSet, [2]int{4, 4})

	var BlueSquareSet [][2]int
	BlueSquareSet = append(BlueSquareSet, [2]int{5, 5})
	BlueSquareSet = append(BlueSquareSet, [2]int{2, 6})
	BlueSquareSet = append(BlueSquareSet, [2]int{6, 2})
	BlueSquareSet = append(BlueSquareSet, [2]int{6, 6})

	var OrangeSquareSet [][2]int
	OrangeSquareSet = append(OrangeSquareSet, [2]int{4, 1})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{1, 4})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{3, 7})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{7, 3})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{7, 7})

	if slices.Contains(CenterSquareSet, pos) {
		return string(constants.CenterSquare)
	} else if slices.Contains(RedSquareSet, pos) {
		return string(constants.RedSquare)
	} else if slices.Contains(YellowSquareSet, pos) {
		return string(constants.YellowSquare)
	} else if slices.Contains(BlueSquareSet, pos) {
		return string(constants.BlueSquare)
	} else if slices.Contains(OrangeSquareSet, pos) {
		return string(constants.OrangeSquare)
	} else {
		return string(constants.NormalSquare)
	}
}

func (s Square) String() string {
	str := "Position: [" + fmt.Sprint(s.Position[0]) +
		", " + fmt.Sprint(s.Position[1]) + "]" +
		"\tType: " + s.SquareType

	if s.ChipPlaceOn.isEmpty() {
		str = str + "\tChip: Empty"
	} else {
		str = str + "\tChip: {" + s.ChipPlaceOn.String() + "}"
	}

	return str
}
