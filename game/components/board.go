package components

import (
	"A-MATH/game/constants"

	"strconv"
	"strings"
)

type Board struct {
	Squares [15][15]square
}

func NewBoard() Board {
	var boardsSquare [15][15]square

	xAxisSet := []string{string(constants.A), string(constants.B), string(constants.C), string(constants.D), string(constants.E), string(constants.F),
		string(constants.G), string(constants.H), string(constants.I), string(constants.J), string(constants.K), string(constants.L), string(constants.M),
		string(constants.N), string(constants.O)}

	yAxisSet := []string{string(constants.One), string(constants.Two), string(constants.Three), string(constants.Four), string(constants.Five),
		string(constants.Six), string(constants.Seven), string(constants.Eight), string(constants.Nine), string(constants.Ten), string(constants.Eleven),
		string(constants.Twelve), string(constants.Thirteen), string(constants.Fourteen), string(constants.Fifteen)}

	for posX := 0; posX < 15; posX++ {
		for posY := 0; posY < 15; posY++ {
			position := xAxisSet[posX] + yAxisSet[posY]
			boardsSquare[posY][posX], _ = NewSquare(position)
		}
	}

	return Board{boardsSquare}
}

func (b *Board) Add(pos [2]int, c Chip) {
	posX, posY := pos[1]-1, pos[0]-1

	if b.Squares[posY][posX].ChipPlaceOn.IsEmpty() {
		b.Squares[posY][posX].ChipPlaceOn = c
	}
}

func (b *Board) Remove(pos [2]int) {
	posX, posY := pos[1]-1, pos[0]-1

	c := Chip{}
	if !b.Squares[posY][posX].ChipPlaceOn.IsEmpty() {
		b.Squares[posY][posX].ChipPlaceOn = c
	}
}

func (b *Board) IsEmpty() bool {
	for posX := 0; posX < 15; posX++ {
		for posY := 0; posY < 15; posY++ {
			if !b.Squares[posY][posX].ChipPlaceOn.IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (b Board) String() string {
	var sb strings.Builder

	// Build the board rows
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			sb.WriteString("[")
			if b.Squares[j][i].IsChipPlacedOn() {
				val := b.Squares[j][i].ChipPlaceOn.Value
				if len(val) < 2 {
					sb.WriteString(" ")
				}
				sb.WriteString(val)
			} else if b.Squares[j][i].SquareType != string(constants.NormalSquare) {
				sb.WriteString("$")
				sb.WriteString(string(b.Squares[j][i].SquareType[:1]))
			} else {
				sb.WriteString("  ")
			}
			sb.WriteString("] ")
		}
		if i < 9 {
			sb.WriteString("0")
		}
		sb.WriteString(strconv.Itoa(i + 1))
		sb.WriteString("\n")
	}

	// Build the column numbers
	sb.WriteString("  ")
	for i := 0; i < 15; i++ {
		sb.WriteString(string(rune(i + 'A')))
		sb.WriteString("    ")
	}
	sb.WriteString("\n")

	return sb.String()
}
