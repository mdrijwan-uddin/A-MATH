package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"fmt"
)

func SingleChipMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	singleChipConnectorSet [][2]int,
) [][]models.ChipForCalculating {

	chipForCalculatingSet := [][]models.ChipForCalculating{}

	// Determine the indexes of relevant connectors
	indexes := connectorsIndex(singleChipConnectorSet)

	// Handle single index case
	if len(indexes) == 1 {
		chipForCalculating := singleChipOneDirectionMapping(board, chipForPlacing, singleChipConnectorSet, indexes)
		chipForCalculatingSet = append(chipForCalculatingSet, chipForCalculating)
	}
	if len(indexes) == 2 {
		if indexes[0] == 0 && indexes[1] == 2 { //[[X X] [0 0] [X X] [0 0]]
			chipForCalculating := []models.ChipForCalculating{}
			leftConnecterMapping(board, singleChipConnectorSet[2], &chipForCalculating)
			straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
			rightConnecterMapping(board, singleChipConnectorSet[0], &chipForCalculating)
			chipForCalculatingSet = append(chipForCalculatingSet, chipForCalculating)
		}
		if indexes[0] == 1 && indexes[1] == 3 { //[[0 0] [X X] [0 0] [X X]]
			chipForCalculating := []models.ChipForCalculating{}
			upperConnecterMapping(board, singleChipConnectorSet[3], &chipForCalculating)
			straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
			lowerConnecterMapping(board, singleChipConnectorSet[1], &chipForCalculating)
			chipForCalculatingSet = append(chipForCalculatingSet, chipForCalculating)
		}

	}

	return chipForCalculatingSet
}

func singleChipOneDirectionMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	singleChipConnectorSet [][2]int,
	indexes []int,
) []models.ChipForCalculating {
	chipForCalculating := []models.ChipForCalculating{}
	if indexes[0] == 0 { //[[X X] [0 0] [0 0] [0 0]]
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
		directionsMapping(board, singleChipConnectorSet, &chipForCalculating, string(constants.Right))
	}
	if indexes[0] == 1 { //[[0 0] [X X] [0 0] [0 0]]
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
		directionsMapping(board, singleChipConnectorSet, &chipForCalculating, string(constants.Down))
	}
	if indexes[0] == 2 { //[[0 0] [0 0] [X X] [0 0]]
		directionsMapping(board, singleChipConnectorSet, &chipForCalculating, string(constants.Left))
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
	}
	if indexes[0] == 3 { //[[0 0] [0 0] [0 0] [X X]]
		directionsMapping(board, singleChipConnectorSet, &chipForCalculating, string(constants.Up))
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
	}
	return chipForCalculating
}

func directionsMapping(
	board components.Board,
	chipConnectorSet [][2]int,
	chipForCalculatingSet *[]models.ChipForCalculating,
	directions string,
) {
	var chipsOnBoard []models.ChipForPlacing

	// Define direction vectors for movement
	var delta [4]int
	var chipConnector [2]int
	switch directions {
	case string(constants.Right):
		delta = [4]int{1, 0, 0, 0} // Move left
		chipConnector = chipConnectorSet[0]
	case string(constants.Down):
		delta = [4]int{0, 1, 0, 0} // Move left
		chipConnector = chipConnectorSet[1]
	case string(constants.Left):
		delta = [4]int{0, 0, 1, 0} // Move left
		chipConnector = chipConnectorSet[2]
	case string(constants.Up):
		delta = [4]int{0, 0, 0, 1} // Move up
		chipConnector = chipConnectorSet[3]
	}

	currentPosition := chipConnector
	var counter = 0

	fmt.Println("currentPosition: ")
	fmt.Println(currentPosition)

	// count total place-on-board chip in specified direction
	for counter <= constants.BoardRange {
		position := [2]int{
			currentPosition[0] + counter*delta[0] - counter*delta[2],
			currentPosition[1] + counter*delta[1] - counter*delta[3],
		}

		if !board.IsValidSquare(position) {
			break
		}

		square := board.GetSquare(position)
		if !board.IsValidSquare(position) || square.ChipPlaceOn.IsEmpty() {
			break
		}
		counter++
	}

	fmt.Println("counter: ")
	fmt.Println(counter)

	// Backtrack to the starting point of the line
	currentPosition = [2]int{
		chipConnector[0] - (counter-1)*delta[2],
		chipConnector[1] - (counter-1)*delta[3],
	}

	fmt.Println("currentPosition: ")
	fmt.Println(currentPosition)

	// Collect chips along the line
	for i := 0; i < counter; i++ {
		position := [2]int{
			currentPosition[0] + i*delta[0] + i*delta[2],
			currentPosition[1] + i*delta[1] + i*delta[3],
		}
		fmt.Println("position: ")
		fmt.Println(position)
		square := board.GetSquare(position)

		chipsOnBoard = append(chipsOnBoard, models.ChipForPlacing{
			Position:     square.Position,
			Chip:         square.ChipPlaceOn,
			SelectedChip: components.Chip{}, // Empty chip as placeholder
		})
	}

	// Map the detected chips in a straight line
	straightLineMapping(board, chipsOnBoard, true, chipForCalculatingSet)
}

func rightConnecterMapping(
	board components.Board,
	chipConnector [2]int,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {

	currentPosition, fixedCoord := chipConnector[0], chipConnector[1]
	var chipsOnBoard []models.ChipForPlacing

	for currentPosition <= 15 {
		position := [2]int{currentPosition, fixedCoord}
		square := board.GetSquare(position)

		if square.ChipPlaceOn.IsEmpty() {
			break
		}

		chipsOnBoard = append(chipsOnBoard, models.ChipForPlacing{
			Position:     square.Position,
			Chip:         square.ChipPlaceOn,
			SelectedChip: components.Chip{}, // Empty chip as placeholder
		})
		currentPosition++
	}

	// Map the detected chips in a straight line
	straightLineMapping(board, chipsOnBoard, true, chipForCalculatingSet)
}

func lowerConnecterMapping(
	board components.Board,
	chipConnector [2]int,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {

	fixedCoord, currentPosition := chipConnector[0], chipConnector[1]
	var chipsOnBoard []models.ChipForPlacing

	for currentPosition <= 15 {
		position := [2]int{fixedCoord, currentPosition}
		square := board.GetSquare(position)

		if square.ChipPlaceOn.IsEmpty() {
			break
		}

		chipsOnBoard = append(chipsOnBoard, models.ChipForPlacing{
			Position:     square.Position,
			Chip:         square.ChipPlaceOn,
			SelectedChip: components.Chip{}, // Empty chip as placeholder
		})
		currentPosition++
	}

	// Map the detected chips in a straight line
	straightLineMapping(board, chipsOnBoard, true, chipForCalculatingSet)
}

func leftConnecterMapping(
	board components.Board,
	chipConnector [2]int,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {

	var counter = 0
	currentPosition, fixedCoord := chipConnector[0], chipConnector[1]
	var chipsOnBoard []models.ChipForPlacing

	for counter <= 15 {
		position := [2]int{currentPosition - counter, fixedCoord}
		if !board.IsValidSquare(position) {
			break
		}

		square := board.GetSquare(position)
		if square.ChipPlaceOn.IsEmpty() {
			break
		}
		counter++
	}

	currentPosition = currentPosition - counter + 1

	for i := 0; i < counter; i++ {
		position := [2]int{currentPosition + i, fixedCoord}
		square := board.GetSquare(position)

		chipsOnBoard = append(chipsOnBoard, models.ChipForPlacing{
			Position:     square.Position,
			Chip:         square.ChipPlaceOn,
			SelectedChip: components.Chip{}, // Empty chip as placeholder
		})
	}

	// Map the detected chips in a straight line
	straightLineMapping(board, chipsOnBoard, true, chipForCalculatingSet)
}

func upperConnecterMapping(
	board components.Board,
	chipConnector [2]int,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {

	var counter = 0
	fixedCoord, currentPosition := chipConnector[0], chipConnector[1]
	var chipsOnBoard []models.ChipForPlacing

	for counter <= 15 {
		position := [2]int{fixedCoord, currentPosition - counter}
		square := board.GetSquare(position)
		if square.ChipPlaceOn.IsEmpty() {
			break
		}
		counter++
	}

	currentPosition = currentPosition - counter + 1

	for i := 0; i < counter; i++ {
		position := [2]int{fixedCoord, currentPosition + i}
		square := board.GetSquare(position)

		chipsOnBoard = append(chipsOnBoard, models.ChipForPlacing{
			Position:     square.Position,
			Chip:         square.ChipPlaceOn,
			SelectedChip: components.Chip{}, // Empty chip as placeholder
		})
	}

	// Map the detected chips in a straight line
	straightLineMapping(board, chipsOnBoard, true, chipForCalculatingSet)
}
