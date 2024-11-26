package rules

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
)

func IsChipPlaceCorrectly(board components.Board, position [][2]int) {
	if !IsChipPlaceOnCenterSquare(position) && board.IsEmpty() {
		return //add error
	}

	if IsChipsPlacedOnOccupiedSquare(board, position) {
		return //add error
	}

	isVertical, isHorizontal := IsChipPlaceOnVerticalOrHorizontal(position)
	if !isVertical && !isHorizontal {
		return //add error
	}
}

func IsChipPlaceOnCenterSquare(position [][2]int) bool {
	for _, pos := range position {
		if pos == [2]int{8, 8} {
			return true
		}
	}
	return false
}

func IsChipPlaceOnVerticalOrHorizontal(position [][2]int) (bool, bool) {
	if len(position) == 1 {
		return true, true
	}

	isVertical := true
	isHorizontal := true
	firstRow, firstColumn := position[0][0], position[0][1]

	for _, pos := range position[1:] {
		if pos[0] != firstRow {
			isVertical = false
		}

		if pos[1] != firstColumn {
			isHorizontal = false
		}

		// If neither vertical nor horizontal alignment is possible, exit early
		if !isVertical && !isHorizontal {
			return false, false
		}
	}

	return isVertical, isHorizontal
}

func IsChipsPlacedOnOccupiedSquare(board components.Board, position [][2]int) bool {
	for _, pos := range position {
		if board.Squares[pos[1]-1][pos[0]-1].HasChipPlacedOn() {
			return true
		}
	}
	return false
}

func ChipPlacementConnector(board components.Board, position [][2]int, isVertical, isHorizontal bool) {
	if len(position) == 0 {
		return
	}
	if len(position) == 1 {
		return
	}

	edgeConnectorSet := EdgeConnector(board, position, isVertical, isHorizontal)
	crossConnectorSet := CrossConnector(board, position, isVertical, isHorizontal)

	if len(edgeConnectorSet) == 0 {
		return
	}
	if len(crossConnectorSet) == 0 {
		return
	}

}

func SingleChipConnector(board components.Board, position [2]int) [][2]int {
	var connector [][2]int
	isAvailable := func(pos int) bool {
		return 0 < pos && pos <= constants.BoardRange
	}

	addConnector := func(x, y int) {
		if isAvailable(x) && isAvailable(y) && board.Squares[y-1][x-1].HasChipPlacedOn() {
			connector = append(connector, [2]int{x, y})
		}
	}

	x, y := position[0], position[1]
	addConnector(x-1, y)
	addConnector(x+1, y)
	addConnector(x, y-1)
	addConnector(x, y+1)
	return connector
}

func EdgeConnector(board components.Board, position [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connector [][2]int

	isAvailable := func(pos int) bool {
		return 0 < pos && pos <= constants.BoardRange
	}

	addConnector := func(x, y int) {
		if isAvailable(x) && isAvailable(y) && board.Squares[y-1][x-1].HasChipPlacedOn() {
			connector = append(connector, [2]int{x, y})
		}
	}

	if isVertical {
		addConnector(position[0][0], position[0][1]-1)
		for _, pos := range position {
			addConnector(pos[0]-1, pos[1])
			addConnector(pos[0]+1, pos[1])
		}
		addConnector(position[len(position)-1][0], position[len(position)-1][1]+1)
	}

	if isHorizontal {
		addConnector(position[0][0]-1, position[0][1])
		for _, pos := range position {
			addConnector(pos[0], pos[1]-1)
			addConnector(pos[0], pos[1]+1)
		}
		addConnector(position[len(position)-1][0]+1, position[len(position)-1][1])
	}

	return connector
}

func CrossConnector(board components.Board, position [][2]int, isVertical, isHorizontal bool) [][2]int {
	var connector [][2]int
	if isVertical {
		for i := 0; i < len(position)-1; i++ {
			if position[i][1]+1 != position[i+1][1] {
				connector = append(connector, [2]int{position[i][0], position[i][1] + 1})
			}
		}
	}

	if isHorizontal {
		for i := 0; i < len(position)-1; i++ {
			if position[i][0]+1 != position[i+1][0] {
				connector = append(connector, [2]int{position[i][0] + 1, position[i][1]})
			}
		}
	}

	return connector
}
