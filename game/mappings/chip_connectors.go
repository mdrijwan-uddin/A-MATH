package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
)

func SingleChipConnector(board components.Board, coordinates [2]int) [][2]int {
	var connectors [][2]int
	directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	addDirectionalConnectors(board, coordinates[0], coordinates[1], directions, &connectors)
	return connectors
}

func PrefixSuffixConnector(board components.Board, coordinates [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connectors [][2]int

	if isVertical {
		addConnector(board, coordinates[0][0], coordinates[0][1]-1, &connectors)
		addConnector(board, coordinates[len(coordinates)-1][0], coordinates[len(coordinates)-1][1]+1, &connectors)
	}

	if isHorizontal {
		addConnector(board, coordinates[0][0]-1, coordinates[0][1], &connectors)
		addConnector(board, coordinates[len(coordinates)-1][0]+1, coordinates[len(coordinates)-1][1], &connectors)
	}

	return connectors
}

func EdgeConnector(board components.Board, coordinates [][2]int, isVertical bool) [][2]int {
	var connectors [][2]int

	if isVertical {
		for _, co := range coordinates {
			addDirectionalConnectors(board, co[0], co[1], [][2]int{{-1, 0}, {1, 0}}, &connectors)
		}
	}

	if !isVertical { //isHorizontal
		for _, co := range coordinates {
			addDirectionalConnectors(board, co[0], co[1], [][2]int{{0, -1}, {0, 1}}, &connectors)
		}
	}

	return connectors
}

// might not be used
func CrossConnector(board components.Board, coordinates [][2]int, isVertical bool) [][2]int {
	var start, end, fixedCoord int
	var connectors [][2]int

	for i := 0; i < len(coordinates)-1; i++ {
		if isVertical {
			start, end = coordinates[i][1], coordinates[i+1][1]
			fixedCoord = coordinates[i][0]
			if start+1 != end { // Check for separation
				connectors = append(connectors, [2]int{fixedCoord, start + 1})
			}
		}

		if !isVertical { //isHorizontal
			start, end = coordinates[i][0], coordinates[i+1][0]
			fixedCoord = coordinates[i][1]
			if start+1 != end { // Check for separation
				connectors = append(connectors, [2]int{start + 1, fixedCoord})
			}
		}
	}

	return connectors
}

func addConnector(board components.Board, x, y int, connectors *[][2]int) {
	isAvailable := func(x, y int) bool {
		return 0 < x && x <= constants.BoardRange && 0 < y && y <= constants.BoardRange
	}

	if isAvailable(x, y) && board.GetSquare([2]int{x, y}).HasChipPlacedOn() {
		*connectors = append(*connectors, [2]int{x, y})
	} else {
		*connectors = append(*connectors, [2]int{})
	}
}

func addDirectionalConnectors(board components.Board, x, y int, directions [][2]int, connectors *[][2]int) {
	for _, delta := range directions {
		addConnector(board, x+delta[0], y+delta[1], connectors)
	}
}
