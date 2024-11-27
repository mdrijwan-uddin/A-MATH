package rules

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
)

func IsChipPlaceCorrectly(board components.Board, coordinate [][2]int) {
	if !IsChipPlaceOnCenterSquare(coordinate) && board.IsEmpty() {
		return //add error
	}

	if IsChipsPlacedOnOccupiedSquare(board, coordinate) {
		return //add error
	}

	isVertical, isHorizontal := IsChipPlaceOnVerticalOrHorizontal(coordinate)
	if !isVertical && !isHorizontal {
		return //add error
	}
}

func IsChipPlaceOnCenterSquare(coordinate [][2]int) bool {
	for _, co := range coordinate {
		if co == [2]int{8, 8} {
			return true
		}
	}
	return false
}

func IsChipPlaceOnVerticalOrHorizontal(coordinate [][2]int) (bool, bool) {
	if len(coordinate) == 1 {
		return true, true
	}

	isVertical := true
	isHorizontal := true
	firstRow, firstColumn := coordinate[0][0], coordinate[0][1]

	for _, co := range coordinate[1:] {
		if co[0] != firstRow {
			isVertical = false
		}

		if co[1] != firstColumn {
			isHorizontal = false
		}

		// If neither vertical nor horizontal alignment is possible, exit early
		if !isVertical && !isHorizontal {
			return false, false
		}
	}

	return isVertical, isHorizontal
}

func IsChipsPlacedOnOccupiedSquare(board components.Board, coordinate [][2]int) bool {
	for _, co := range coordinate {
		if board.GetSquare(co).HasChipPlacedOn() {
			return true
		}
	}
	return false
}

func ChipPlacementConnector(board components.Board, coordinate [][2]int, isVertical, isHorizontal bool) {
	if len(coordinate) == 0 {
		return
	}
	if len(coordinate) == 1 {
		return
	}

	edgeConnectorSet := EdgeConnector(board, coordinate, isVertical, isHorizontal)
	crossConnectorSet := CrossConnector(board, coordinate, isVertical, isHorizontal)

	if len(edgeConnectorSet) == 0 {
		return
	}
	if len(crossConnectorSet) == 0 {
		return
	}

}

func SingleChipConnector(board components.Board, coordinate [2]int) [][2]int {
	var connectors [][2]int

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	addDirectionalConnectors(board, coordinate[0], coordinate[1], directions, &connectors)

	return connectors
}

func PrefixSuffixConnector(board components.Board, coordinate [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connectors [][2]int

	if isVertical {
		addConnector(board, coordinate[0][0], coordinate[0][1]-1, &connectors)
		addConnector(board, coordinate[len(coordinate)-1][0], coordinate[len(coordinate)-1][1]+1, &connectors)
	}

	if isHorizontal {
		addConnector(board, coordinate[0][0]-1, coordinate[0][1], &connectors)
		addConnector(board, coordinate[len(coordinate)-1][0]+1, coordinate[len(coordinate)-1][1], &connectors)
	}

	return connectors
}

func EdgeConnector(board components.Board, coordinate [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connectors [][2]int

	if isVertical {
		for _, co := range coordinate {
			addDirectionalConnectors(board, co[0], co[1], [][2]int{{-1, 0}, {1, 0}}, &connectors)
		}
	}

	if isHorizontal {
		for _, co := range coordinate {
			addDirectionalConnectors(board, co[0], co[1], [][2]int{{0, -1}, {0, 1}}, &connectors)
		}
	}

	return connectors
}

func CrossConnector(board components.Board, coordinate [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connectors [][2]int

	if isVertical {
		for i := 0; i < len(coordinate)-1; i++ {
			if coordinate[i][1]+1 != coordinate[i+1][1] {
				connectors = append(connectors, [2]int{coordinate[i][0], coordinate[i][1] + 1})
			}
		}
	}

	if isHorizontal {
		for i := 0; i < len(coordinate)-1; i++ {
			if coordinate[i][0]+1 != coordinate[i+1][0] {
				connectors = append(connectors, [2]int{coordinate[i][0] + 1, coordinate[i][1]})
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
	}
}

func addDirectionalConnectors(board components.Board, x, y int, directions [][2]int, connectors *[][2]int) {
	for _, delta := range directions {
		addConnector(board, x+delta[0], y+delta[1], connectors)
	}
}
