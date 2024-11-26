package components

import (
	"A-MATH/game/constants"
	"fmt"

	"strings"

	"github.com/fatih/color"
)

type Board struct {
	Squares [constants.BoardRange][constants.BoardRange]square
}

func NewBoard() Board {
	var boardsSquare [constants.BoardRange][constants.BoardRange]square

	xAxisSet := []string{string(constants.A), string(constants.B), string(constants.C), string(constants.D), string(constants.E), string(constants.F),
		string(constants.G), string(constants.H), string(constants.I), string(constants.J), string(constants.K), string(constants.L), string(constants.M),
		string(constants.N), string(constants.O)}

	yAxisSet := []string{string(constants.One), string(constants.Two), string(constants.Three), string(constants.Four), string(constants.Five),
		string(constants.Six), string(constants.Seven), string(constants.Eight), string(constants.Nine), string(constants.Ten), string(constants.Eleven),
		string(constants.Twelve), string(constants.Thirteen), string(constants.Fourteen), string(constants.Fifteen)}

	for posX := 0; posX < constants.BoardRange; posX++ {
		for posY := 0; posY < constants.BoardRange; posY++ {
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
	for posX := 0; posX < constants.BoardRange; posX++ {
		for posY := 0; posY < constants.BoardRange; posY++ {
			if !b.Squares[posY][posX].ChipPlaceOn.IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (b Board) String() string {
	var sb strings.Builder

	// Define color mapping for square types
	colorMap := map[string]func(a ...interface{}) string{
		"R": color.New(color.FgRed).SprintFunc(),
		"Y": color.New(color.FgYellow).SprintFunc(),
		"B": color.New(color.FgBlue).SprintFunc(),
		"O": color.RGB(255, 128, 0).SprintFunc(),
		"C": color.New(color.FgCyan).SprintFunc(),
	}

	// Function to write colored or default squares
	writeSquare := func(content string) {
		if colorFunc, exists := colorMap[content]; exists {
			sb.WriteString(colorFunc("[$"))
			sb.WriteString(colorFunc(content))
			sb.WriteString(colorFunc("] "))
		} else {
			sb.WriteString("[")
			if len(content) < 2 {
				sb.WriteString(" ")
			}
			sb.WriteString(content)
			sb.WriteString("] ")
		}
	}

	// Build the board rows
	for j := 0; j < constants.BoardRange; j++ {
		for i := 0; i < constants.BoardRange; i++ {
			if b.Squares[j][i].HasChipPlacedOn() {
				writeSquare(b.Squares[j][i].ChipPlaceOn.Value)
			} else if b.Squares[j][i].SquareType != string(constants.NormalSquare) {
				writeSquare(string(b.Squares[j][i].SquareType[:1]))
			} else {
				writeSquare("  ")
			}
		}
		// Append row number
		sb.WriteString(fmt.Sprintf("%02d\n", j+1))
	}

	// Build the column labels
	sb.WriteString("  ")
	for i := 0; i < constants.BoardRange; i++ {
		sb.WriteString(fmt.Sprintf("%-5s", string(rune(i+'A'))))
	}
	sb.WriteString("\n")

	return sb.String()
}
