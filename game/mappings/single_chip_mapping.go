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
	singleChipConnector []models.ChipConnector,
) [][]models.ChipForCalculating {

	chipForCalculatingSet := [][]models.ChipForCalculating{}
	connector := singleChipConnector[0]

	// Determine the totalConnector of relevant connectors
	totalConnector := connector.TotalConnecter()

	// Check totalConnector
	if totalConnector == 1 {
		singleChipOneDirectionMapping(board, chipForPlacing, connector, &chipForCalculatingSet)
	} else if totalConnector == 2 {
		singleChipTwoDirectionMapping(board, chipForPlacing, connector, &chipForCalculatingSet)
	} else if totalConnector == 3 {
		singleChipThreeDirectionMapping(board, chipForPlacing, connector, &chipForCalculatingSet)
	} else {
		singleChipAllDirectionMapping(board, chipForPlacing, connector, &chipForCalculatingSet)
	}

	return chipForCalculatingSet
}

func singleChipOneDirectionMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	connector models.ChipConnector,
	chipForCalculatingSet *[][]models.ChipForCalculating,
) {
	chipForCalculating := []models.ChipForCalculating{}

	if connector.RightConnected { //[[X][0][0][0][0][0][0][0][0]] horizontal
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
		directionsMapping(board, connector, &chipForCalculating, string(constants.Right))

	} else if connector.BottomConnected { //[[X][0][0][0][0][0][0][0][0]] vertical
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
		directionsMapping(board, connector, &chipForCalculating, string(constants.Down))

	} else if connector.LeftConnected { //[[0][0][0][0][0][0][0][0][X]] horizontal
		directionsMapping(board, connector, &chipForCalculating, string(constants.Left))
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)

	} else if connector.TopConnected { //[[0][0][0][0][0][0][0][0][X]] vertical
		directionsMapping(board, connector, &chipForCalculating, string(constants.Up))
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
	}

	*chipForCalculatingSet = append(*chipForCalculatingSet, chipForCalculating)
}

func singleChipTwoDirectionMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	connector models.ChipConnector,
	chipForCalculatingSet *[][]models.ChipForCalculating,
) {
	if connector.LeftConnected && connector.RightConnected {

		//[[0][0][0][0][X][0][0][0][0]] horizontal
		chipForCalculating := []models.ChipForCalculating{}

		directionsMapping(board, connector, &chipForCalculating, string(constants.Left))
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
		directionsMapping(board, connector, &chipForCalculating, string(constants.Right))

		*chipForCalculatingSet = append(*chipForCalculatingSet, chipForCalculating)

	} else if connector.TopConnected && connector.BottomConnected {

		//[[0][0][0][0][X][0][0][0][0]] vertical
		chipForCalculating := []models.ChipForCalculating{}

		directionsMapping(board, connector, &chipForCalculating, string(constants.Up))
		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
		directionsMapping(board, connector, &chipForCalculating, string(constants.Down))

		*chipForCalculatingSet = append(*chipForCalculatingSet, chipForCalculating)

	} else {

		// other
		firstConnector, secondConnector := connector.TemporarySeperateConnector()
		fmt.Println(firstConnector)
		fmt.Println(secondConnector)

		singleChipOneDirectionMapping(board, chipForPlacing, firstConnector, chipForCalculatingSet)
		singleChipOneDirectionMapping(board, chipForPlacing, secondConnector, chipForCalculatingSet)
	}
}

func singleChipThreeDirectionMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	connector models.ChipConnector,
	chipForCalculatingSet *[][]models.ChipForCalculating,
) {
	// other
	firstConnector, secondConnector := connector.TemporarySeperateConnector()

	singleChipTwoDirectionMapping(board, chipForPlacing, firstConnector, chipForCalculatingSet)
	singleChipOneDirectionMapping(board, chipForPlacing, secondConnector, chipForCalculatingSet)
}

func singleChipAllDirectionMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	connector models.ChipConnector,
	chipForCalculatingSet *[][]models.ChipForCalculating,
) {
	// other
	firstConnector, secondConnector := connector.TemporarySeperateConnector()

	singleChipTwoDirectionMapping(board, chipForPlacing, firstConnector, chipForCalculatingSet)
	singleChipTwoDirectionMapping(board, chipForPlacing, secondConnector, chipForCalculatingSet)
}

func directionsMapping(
	board components.Board,
	connector models.ChipConnector,
	chipForCalculatingSet *[]models.ChipForCalculating,
	directions string,
) {
	var chipsOnBoard []models.ChipForPlacing

	// Define direction vectors for movement
	var delta [4]int
	var chipConnector [2]int
	switch directions {
	case string(constants.Right):
		delta = [4]int{1, 0, 0, 0} // Move right
		chipConnector = connector.RightCoordinate()
	case string(constants.Down):
		delta = [4]int{0, 1, 0, 0} // Move down
		chipConnector = connector.BottomCoordinate()
	case string(constants.Left):
		delta = [4]int{0, 0, 1, 0} // Move left
		chipConnector = connector.LeftCoordinate()
	case string(constants.Up):
		delta = [4]int{0, 0, 0, 1} // Move up
		chipConnector = connector.TopCoordinate()
	}

	currentPosition := chipConnector
	var counter = 0

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

	// Backtrack to the starting point of the line
	currentPosition = [2]int{
		chipConnector[0] - (counter-1)*delta[2],
		chipConnector[1] - (counter-1)*delta[3],
	}

	// Collect chips along the line
	for i := 0; i < counter; i++ {
		position := [2]int{
			currentPosition[0] + i*delta[0] + i*delta[2],
			currentPosition[1] + i*delta[1] + i*delta[3],
		}
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
