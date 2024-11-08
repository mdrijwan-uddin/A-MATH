package game

import (
	"fmt"

	. "A-MATH/constants"
	"A-MATH/err"
	"strconv"
)

type Square struct {
	Position    [2]int
	SquareType  string
	ChipPlaceOn Chip
}

func NewSquare(pos string) (Square, error) {
	position, e := setSquarePosition(pos)
	if e != nil {
		return Square{}, e
	}

	return Square{position, "", Chip{}}, nil
}

func setSquarePosition(pos string) ([2]int, error) {
	if len(pos) != 2 && len(pos) != 3 {
		return [2]int{}, err.New(int(BadRequest), string(InvalidInputForBoard))
	}

	xAxis := int(pos[0] - 64)
	if xAxis < 0 || xAxis > 15 {
		return [2]int{}, err.New(int(BadRequest), string(InvalidInputForBoard))
	}

	yAxis, _ := strconv.Atoi(pos[1:])
	if yAxis < 0 || yAxis > 15 {
		return [2]int{}, err.New(int(BadRequest), string(InvalidInputForBoard))
	}

	return [2]int{xAxis, yAxis}, nil
}

func setSquareType(pos [2]int)

func (s Square) String() string {
	return "Position: [" + fmt.Sprint(s.Position[0]) +
		", " + fmt.Sprint(s.Position[1]) + "]" +
		"\tType: " + s.SquareType +
		"\tChip: " + s.ChipPlaceOn.String()
}

//    AA BB CC DD EE FF GG HH II JJ KK LL MM NN OO
// 01
// 02
// 03
// 04
// 05
// 06
// 07
// 08
// 09
// 10
// 11
// 12
// 13
// 14
// 15

// red = 8

// {1,1},{8,1},{15,1},{1,8},{15,8},{1,15},{8,15},{15,15}

// red :
// {1,1}: {1,1},{8,1},{15,1},{1,8},{15,8},{1,15},{8,15},{15,15}
// -------------------------------------------------------------

// yellow = 12 {2,2},{14,2},{3,3},{13,3},{4,4},{12,4},{4,12},{12,12},{3,13},{13,13},{2,14},{14,14}

// yellow :
// {2,2}: {14,2},{2,14},{14,14}
// {3,3}: {13,3},{3,13},{13,13}
// {4,4}: {12,4},{4,12},{12,12}
// -------------------------------------------------------------

// blue = 16 {6,2},{10,2},
// {5,5},{11,5},{2,6},{6,6},{10,6},{14,6}, {2,10},{6,10},{10,10},{14,10},{5,11},{11,11},{6,14},{10,14}

// blue :
// {5,5}: {11,5},{5,11},{11,11}
// {6,2}: {10,2},{6,14},{10,14}
// {2,6}: {14,6},{2,10},{14,10}
// {6,6}: {10,6},{6,10},{10,10}
// -------------------------------------------------------------

// orange = 24 {4,1},{12,1},{7,3},{9,3},{1,4},{8,4},{15,4},
// {3,7},{7,7},{9,7},{13,7},{4,8},{12,8},{3,9},{7,9},{9,9},{13,9},
// {1,12},{8,12},{15,12},{7,13},{9,13},{4,15},{12,15}

// orange :
// {4,1}: {12,1},{4,8},{12,8},{4,15},{12,15}
// {1,4}: {8,4},{15,4},{1,12},{8,12},{15,12}
// {3,7}: {13,7},{3,9},{13,9}
// {7,3}: {9,3},{7,13},{9,13}
// {7,7}: {9,7},{7,9},{9,9}
