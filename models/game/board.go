package game

import (
	"A-MATH/constants"
)

type Board struct {
	Squares [15][15]Square
}

func NewBoard() Board {
	var boardsSquare [15][15]Square

	xAxisSet := []string{string(constants.A), string(constants.B), string(constants.C), string(constants.D), string(constants.E), string(constants.F),
		string(constants.G), string(constants.H), string(constants.I), string(constants.J), string(constants.K), string(constants.L), string(constants.M),
		string(constants.N), string(constants.O)}

	yAxisSet := []string{string(constants.One), string(constants.Two), string(constants.Three), string(constants.Four), string(constants.Five),
		string(constants.Six), string(constants.Seven), string(constants.Eight), string(constants.Nine), string(constants.Ten), string(constants.Eleven),
		string(constants.Twelve), string(constants.Thirteen), string(constants.Fourteen), string(constants.Fifteen)}

	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			position := xAxisSet[i] + yAxisSet[j]
			boardsSquare[j][i], _ = NewSquare(position)
		}
	}

	return Board{boardsSquare}
}

func (b Board) String() string {
	var str string
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			str += string(b.Squares[j][i].String()) + "\n"
		}
		str += "-----------------------------------------------------------\n"
	}
	return str
}
