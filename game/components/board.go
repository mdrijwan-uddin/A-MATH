package components

import (
	"A-MATH/game/constants"
	"A-MATH/game/utils"
	"fmt"

	"strings"

	"github.com/fatih/color"
)

type Board struct {
	squares [constants.BoardRange][constants.BoardRange]Square
}

func NewBoard() Board {
	var boardsSquare [constants.BoardRange][constants.BoardRange]Square

	xAxisSet := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	}

	yAxisSet := []string{
		string(constants.One), string(constants.Two), string(constants.Three), string(constants.Four),
		string(constants.Five), string(constants.Six), string(constants.Seven), string(constants.Eight),
		string(constants.Nine), string(constants.Ten), string(constants.Eleven), string(constants.Twelve),
		string(constants.Thirteen), string(constants.Fourteen), string(constants.Fifteen),
	}

	for x := 0; x < constants.BoardRange; x++ {
		for y := 0; y < constants.BoardRange; y++ {
			position := xAxisSet[x] + yAxisSet[y]
			coordinate, _ := utils.ValidateSquarePosition(position)
			boardsSquare[y][x] = NewSquare(coordinate)
		}
	}

	return Board{boardsSquare}
}

// input: {X, Y} | output: Square{X, Y}
func (b *Board) GetSquare(coordinate [2]int) *Square {
	return &b.squares[coordinate[1]-1][coordinate[0]-1]
}

// check if {0 < X,Y <= 15} or not
func (b *Board) IsValidSquare(coordinate [2]int) bool {
	return 0 < coordinate[0] && coordinate[0] <= 15 &&
		0 < coordinate[1] && coordinate[1] <= 15
}

func (b *Board) Add(coordinate [2]int, c Chip) {
	selectedSquare := b.GetSquare(coordinate)
	if selectedSquare.ChipPlaceOn.IsEmpty() {
		selectedSquare.ChipPlaceOn = c
	}
}

func (b *Board) Remove(coordinate [2]int) Chip {
	c := Chip{}
	var selectedChip Chip
	selectedSquare := b.GetSquare(coordinate)
	if !selectedSquare.ChipPlaceOn.IsEmpty() {
		selectedChip = selectedSquare.ChipPlaceOn
		selectedSquare.ChipPlaceOn = c
	}
	return selectedChip
}

// has no chip place on board
func (b *Board) IsEmpty() bool {
	for y := 0; y < constants.BoardRange; y++ {
		for x := 0; x < constants.BoardRange; x++ {
			coordinate := [2]int{x + 1, y + 1}
			if !b.GetSquare(coordinate).ChipPlaceOn.IsEmpty() {
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
	for y := 0; y < constants.BoardRange; y++ {
		for x := 0; x < constants.BoardRange; x++ {
			targetedSquare := b.GetSquare([2]int{x + 1, y + 1})

			if targetedSquare.HasChipPlacedOn() {
				writeSquare(targetedSquare.ChipPlaceOn.Value)
			} else if targetedSquare.SquareType != string(constants.NormalSquare) {
				writeSquare(string(targetedSquare.SquareType[:1]))
			} else {
				writeSquare("  ")
			}
		}
		// Append row number
		sb.WriteString(fmt.Sprintf("%02d\n", y+1))
	}

	// Build the column labels
	sb.WriteString("  ")
	for i := 0; i < constants.BoardRange; i++ {
		sb.WriteString(fmt.Sprintf("%-5s", string(rune(i+'A'))))
	}
	sb.WriteString("\n")

	return sb.String()
}
