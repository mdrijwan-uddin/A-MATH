package models

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"strconv"
	"strings"
)

type ChipConnector struct {
	ChipCoordinate  [2]int
	LeftConnected   bool
	TopConnected    bool
	RightConnected  bool
	BottomConnected bool
}

func NewChipConnector(position [2]int) ChipConnector {
	return ChipConnector{position, false, false, false, false}
}

func (c ChipConnector) TotalConnecter() int {
	var counter int
	if c.LeftConnected {
		counter++
	}
	if c.TopConnected {
		counter++
	}
	if c.RightConnected {
		counter++
	}
	if c.BottomConnected {
		counter++
	}
	return counter
}

func (c *ChipConnector) CheckVerticalConnector(board components.Board) {
	c.CheckTopConnector(board)
	c.CheckBottomConnector(board)
}

func (c *ChipConnector) CheckHorizontalConnector(board components.Board) {
	c.CheckLeftConnector(board)
	c.CheckRightConnector(board)
}

func (c *ChipConnector) CheckAllDirectionConnector(board components.Board) {
	c.CheckLeftConnector(board)
	c.CheckTopConnector(board)
	c.CheckRightConnector(board)
	c.CheckBottomConnector(board)
}

func (c *ChipConnector) CheckLeftConnector(board components.Board) {
	coord := c.LeftCoordinate()
	c.LeftConnected = c.setConnector(board, coord[0], coord[1])
}

func (c *ChipConnector) CheckTopConnector(board components.Board) {
	coord := c.TopCoordinate()
	c.TopConnected = c.setConnector(board, coord[0], coord[1])
}

func (c *ChipConnector) CheckRightConnector(board components.Board) {
	coord := c.RightCoordinate()
	c.RightConnected = c.setConnector(board, coord[0], coord[1])
}

func (c *ChipConnector) CheckBottomConnector(board components.Board) {
	coord := c.BottomCoordinate()
	c.BottomConnected = c.setConnector(board, coord[0], coord[1])
}

func (c ChipConnector) LeftCoordinate() [2]int {
	return [2]int{c.ChipCoordinate[0] - 1, c.ChipCoordinate[1]}
}

func (c ChipConnector) TopCoordinate() [2]int {
	return [2]int{c.ChipCoordinate[0], c.ChipCoordinate[1] - 1}
}

func (c ChipConnector) RightCoordinate() [2]int {
	return [2]int{c.ChipCoordinate[0] + 1, c.ChipCoordinate[1]}
}

func (c ChipConnector) BottomCoordinate() [2]int {
	return [2]int{c.ChipCoordinate[0], c.ChipCoordinate[1] + 1}
}

func (c ChipConnector) TemporarySeperateConnector() (ChipConnector, ChipConnector) {
	if c.TotalConnecter() == 2 {
		return c.splitTwoConnections()
	} else if c.TotalConnecter() == 3 {
		return c.splitThreeConnections()
	} else if c.TotalConnecter() == 4 {
		return c.splitFourConnections()
	}
	// Default case: disconnected connectors
	return c.disconnectedConnectors()
}

// Helper to handle two connections
func (c ChipConnector) splitTwoConnections() (ChipConnector, ChipConnector) {
	// Identify two connections and split them
	if c.LeftConnected && c.TopConnected {
		return ChipConnector{c.ChipCoordinate, true, false, false, false},
			ChipConnector{c.ChipCoordinate, false, true, false, false}
	}
	if c.LeftConnected && c.RightConnected {
		return ChipConnector{c.ChipCoordinate, true, false, false, false},
			ChipConnector{c.ChipCoordinate, false, false, true, false}
	}
	if c.LeftConnected && c.BottomConnected {
		return ChipConnector{c.ChipCoordinate, true, false, false, false},
			ChipConnector{c.ChipCoordinate, false, false, false, true}
	}
	if c.TopConnected && c.RightConnected {
		return ChipConnector{c.ChipCoordinate, false, true, false, false},
			ChipConnector{c.ChipCoordinate, false, false, true, false}
	}
	if c.TopConnected && c.BottomConnected {
		return ChipConnector{c.ChipCoordinate, false, true, false, false},
			ChipConnector{c.ChipCoordinate, false, false, false, true}
	}
	if c.RightConnected && c.BottomConnected {
		return ChipConnector{c.ChipCoordinate, false, false, true, false},
			ChipConnector{c.ChipCoordinate, false, false, false, true}
	}
	return c.disconnectedConnectors()
}

// Helper to handle three connections
func (c ChipConnector) splitThreeConnections() (ChipConnector, ChipConnector) {
	switch {
	case !c.LeftConnected:
		return ChipConnector{c.ChipCoordinate, false, true, false, true},
			ChipConnector{c.ChipCoordinate, false, false, true, false}
	case !c.TopConnected:
		return ChipConnector{c.ChipCoordinate, true, false, true, false},
			ChipConnector{c.ChipCoordinate, false, false, false, true}
	case !c.RightConnected:
		return ChipConnector{c.ChipCoordinate, false, true, false, true},
			ChipConnector{c.ChipCoordinate, true, false, false, false}
	case !c.BottomConnected:
		return ChipConnector{c.ChipCoordinate, true, false, true, false},
			ChipConnector{c.ChipCoordinate, false, true, false, false}
	default:
		return c.disconnectedConnectors()
	}
}

// Helper to handle four connections
func (c ChipConnector) splitFourConnections() (ChipConnector, ChipConnector) {
	return ChipConnector{c.ChipCoordinate, true, false, true, false},
		ChipConnector{c.ChipCoordinate, false, true, false, true}
}

// Default helper to return disconnected connectors
func (c ChipConnector) disconnectedConnectors() (ChipConnector, ChipConnector) {
	return ChipConnector{c.ChipCoordinate, false, false, false, false},
		ChipConnector{c.ChipCoordinate, false, false, false, false}
}

func (c ChipConnector) setConnector(board components.Board, x, y int) bool {
	isAvailable := func(x, y int) bool {
		return 0 < x && x <= constants.BoardRange && 0 < y && y <= constants.BoardRange
	}

	if isAvailable(x, y) && board.GetSquare([2]int{x, y}).HasChipPlacedOn() {
		return true
	}
	return false
}

func (c ChipConnector) String() string {
	var sb strings.Builder
	sb.WriteString("{Position: [")
	sb.WriteString(strconv.Itoa(c.ChipCoordinate[0]))
	sb.WriteString(", ")
	sb.WriteString(strconv.Itoa(c.ChipCoordinate[1]))
	sb.WriteString("]| LeftConnected: ")
	sb.WriteString(strconv.FormatBool(c.LeftConnected))
	sb.WriteString("| TopConnected: ")
	sb.WriteString(strconv.FormatBool(c.TopConnected))
	sb.WriteString("| RightConnected: ")
	sb.WriteString(strconv.FormatBool(c.RightConnected))
	sb.WriteString("| BottomConnected: ")
	sb.WriteString(strconv.FormatBool(c.BottomConnected))
	sb.WriteString("}\n")
	return sb.String()
}
