package rules

import (
	"A-MATH/err"
	"A-MATH/game/components"
	"A-MATH/game/constants"
)

func IsChipPlaceCorrectly(board components.Board, coordinates [][2]int) {
	if !IsChipPlaceOnCenterSquare(coordinates) && board.IsEmpty() {
		return //add error
	}

	if IsChipsPlacedOnOccupiedSquare(board, coordinates) {
		return //add error
	}

	isVertical, isHorizontal := IsChipPlaceOnVerticalOrHorizontal(coordinates)
	if !isVertical && !isHorizontal {
		return //add error
	}
}

func IsChipPlaceOnCenterSquare(coordinates [][2]int) bool {
	for _, co := range coordinates {
		if co == [2]int{8, 8} {
			return true
		}
	}
	return false
}

func IsChipPlaceOnVerticalOrHorizontal(coordinates [][2]int) (bool, bool) {
	if len(coordinates) == 1 {
		return true, true
	}

	isVertical := true
	isHorizontal := true
	firstRow, firstColumn := coordinates[0][0], coordinates[0][1]

	for _, co := range coordinates[1:] {
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

func IsChipsPlacedOnOccupiedSquare(board components.Board, coordinates [][2]int) bool {
	for _, co := range coordinates {
		if board.GetSquare(co).HasChipPlacedOn() {
			return true
		}
	}
	return false
}

func IsChipPlacementSeperated(board components.Board, coordinates [][2]int, isVertical, isHorizontal bool) (bool, error) {
	if len(coordinates) < 2 {
		return false, nil // Cannot be separated if less than two chips
	}

	checkGap := func(start, end, fixedCoord int, isVertical bool) error {
		for i := start + 1; i < end; i++ {
			pos := [2]int{}
			if isVertical {
				pos = [2]int{fixedCoord, i}
			} else {
				pos = [2]int{i, fixedCoord}
			}

			if !board.GetSquare(pos).HasChipPlacedOn() {
				return err.New(int(constants.BadRequest), string(constants.InvalidChipPlacement))
			}
		}
		return nil
	}

	var isSeparationDetected = false

	if isVertical {
		for i := 0; i < len(coordinates)-1; i++ {
			start, end := coordinates[i][1], coordinates[i+1][1]
			if start+1 != end { // Check for separation
				isSeparationDetected = true
				if err := checkGap(start, end, coordinates[i][0], true); err != nil {
					return false, err
				}
			}
		}
		return isSeparationDetected, nil
	}

	if isHorizontal {
		for i := 0; i < len(coordinates)-1; i++ {
			start, end := coordinates[i][0], coordinates[i+1][0]
			if start+1 != end { // Check for separation
				isSeparationDetected = true
				if err := checkGap(start, end, coordinates[i][1], false); err != nil {
					return false, err
				}
			}
		}
		return isSeparationDetected, nil
	}

	return false, nil
}

func ChipPlacementConnector(board components.Board, coordinates [][2]int, isVertical, isHorizontal bool) {
	if len(coordinates) == 0 {
		return
	}
	if len(coordinates) == 1 {
		return
	}

	edgeConnectorSet := EdgeConnector(board, coordinates, isVertical, isHorizontal)
	crossConnectorSet := CrossConnector(board, coordinates, isVertical, isHorizontal)

	if len(edgeConnectorSet) == 0 {
		return
	}
	if len(crossConnectorSet) == 0 {
		return
	}

}

func SingleChipConnector(board components.Board, coordinates [2]int) [][2]int {
	var connectors [][2]int

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
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

func EdgeConnector(board components.Board, coordinates [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connectors [][2]int

	if isVertical {
		for _, co := range coordinates {
			addDirectionalConnectors(board, co[0], co[1], [][2]int{{-1, 0}, {1, 0}}, &connectors)
		}
	}

	if isHorizontal {
		for _, co := range coordinates {
			addDirectionalConnectors(board, co[0], co[1], [][2]int{{0, -1}, {0, 1}}, &connectors)
		}
	}

	return connectors
}

func CrossConnector(board components.Board, coordinates [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connectors [][2]int

	if isVertical {
		for i := 0; i < len(coordinates)-1; i++ {
			if coordinates[i][1]+1 != coordinates[i+1][1] {
				connectors = append(connectors, [2]int{coordinates[i][0], coordinates[i][1] + 1})
			}
		}
	}

	if isHorizontal {
		for i := 0; i < len(coordinates)-1; i++ {
			if coordinates[i][0]+1 != coordinates[i+1][0] {
				connectors = append(connectors, [2]int{coordinates[i][0] + 1, coordinates[i][1]})
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
