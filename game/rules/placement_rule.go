package rules

import (
	"A-MATH/game/components"
)

// not finish
func ValidateChipPlacement(board components.Board, coordinates [][2]int) {
	if len(coordinates) == 0 {
		return //add error
	}

	if board.IsEmpty() {
		if len(coordinates) < 3 {
			return //add error
		}

		if !IsChipPlaceOnCenterSquare(coordinates) {
			return //add error
		}
	} else {
		if IsChipsPlacedOnOccupiedSquare(board, coordinates) {
			return //add error
		}

	}

	isVertical, isHorizontal := IsChipPlaceOnVerticalOrHorizontal(coordinates)
	if !isVertical && !isHorizontal {
		return //add error
	}

	isStraightLine, isSeperated := IsChipPlacingOnStraightLineOrSeparated(board, coordinates, isVertical)
	if !isStraightLine && !isSeperated {
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

func IsChipPlacingOnStraightLineOrSeparated(board components.Board, coordinates [][2]int, isVertical bool) (bool, bool) {
	var start, end, fixedCoord int
	var isStraightLine, isCorrectSeparation = true, false

	for i := 0; i < len(coordinates)-1; i++ {
		if isVertical {
			start, end = coordinates[i][1], coordinates[i+1][1]
			fixedCoord = coordinates[i][0]
		}

		if !isVertical { //isHorizontal
			start, end = coordinates[i][0], coordinates[i+1][0]
			fixedCoord = coordinates[i][1]
		}

		if start+1 != end { // Check for separation
			isStraightLine = false
			isCorrectSeparation = checkSeparation(board, start, end, fixedCoord, isVertical)
			if !isCorrectSeparation {
				return isStraightLine, isCorrectSeparation
			}
		}
	}
	return isStraightLine, isCorrectSeparation
}

func checkSeparation(board components.Board, start, end, fixedCoord int, isVertical bool) bool {
	for i := start + 1; i < end; i++ {
		pos := [2]int{}
		if isVertical {
			pos = [2]int{fixedCoord, i}
		}
		if !isVertical { //isHorizontal
			pos = [2]int{i, fixedCoord}
		}

		if !board.GetSquare(pos).HasChipPlacedOn() {
			return false
		}
	}
	return true
}
