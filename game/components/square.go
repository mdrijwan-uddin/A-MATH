package components

import (
	"A-MATH/game/constants"
	"A-MATH/game/utils"

	"slices"
	"strconv"
	"strings"
)

type square struct {
	Position    [2]int
	SquareType  string
	ChipPlaceOn Chip
}

func NewSquare(pos string) (square, error) {
	position, e := utils.ValidateSquarePosition(pos)
	if e != nil {
		return square{}, e
	}

	squareType := setSquareType(position)
	return square{position, squareType, Chip{}}, nil
}

func setSquareType(pos [2]int) string {
	minimalizePosition := func(pos int) int {
		if pos == 8 || pos == 15 {
			return 1
		}
		if pos > 8 {
			return 16 - pos
		}
		return pos
	}

	if pos[0] == 8 && pos[1] == 8 {
		return string(constants.CenterSquare)
	}

	for i := 0; i < 2; i++ {
		pos[i] = minimalizePosition(pos[i])
	}

	var RedSquareSet [][2]int
	RedSquareSet = append(RedSquareSet, [2]int{1, 1})

	var YellowSquareSet [][2]int
	YellowSquareSet = append(YellowSquareSet, [2]int{2, 2})
	YellowSquareSet = append(YellowSquareSet, [2]int{3, 3})
	YellowSquareSet = append(YellowSquareSet, [2]int{4, 4})

	var BlueSquareSet [][2]int
	BlueSquareSet = append(BlueSquareSet, [2]int{2, 6})
	BlueSquareSet = append(BlueSquareSet, [2]int{5, 5})
	BlueSquareSet = append(BlueSquareSet, [2]int{6, 2})
	BlueSquareSet = append(BlueSquareSet, [2]int{6, 6})

	var OrangeSquareSet [][2]int
	OrangeSquareSet = append(OrangeSquareSet, [2]int{1, 4})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{3, 7})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{4, 1})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{7, 3})
	OrangeSquareSet = append(OrangeSquareSet, [2]int{7, 7})

	if slices.Contains(RedSquareSet, pos) {
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

func (s square) IsChipPlacedOn() bool {
	return !s.ChipPlaceOn.IsEmpty()
}

func (s square) String() string {
	var sb strings.Builder
	sb.WriteString("Position: [")
	sb.WriteString(strconv.Itoa(s.Position[0]))
	sb.WriteString(", ")
	sb.WriteString(strconv.Itoa(s.Position[1]))
	sb.WriteString("]|\tType: ")
	sb.WriteString(s.SquareType)
	sb.WriteString("\t|Chip: ")
	if s.ChipPlaceOn.IsEmpty() {
		sb.WriteString("Empty")
	} else {
		sb.WriteString("{")
		sb.WriteString(s.ChipPlaceOn.String())
		sb.WriteString("}")
	}
	return sb.String()
}
