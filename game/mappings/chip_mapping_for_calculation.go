package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
)

// not finish
func ChipPlacementManagement(board components.Board, coordinates [][2]int, isVertical, isStraightLine bool) {
	if len(coordinates) == 1 {
		SingleChipConnectorSet := SingleChipConnector(board, coordinates[0])
		if len(SingleChipConnectorSet) == 0 {
			return // add error
		}
	}

	edgeConnectorSet := EdgeConnector(board, coordinates, isVertical)
	crossConnectorSet := CrossConnector(board, coordinates, isVertical)

	if len(edgeConnectorSet) == 0 {
		return
	}
	if len(crossConnectorSet) == 0 {
		return
	}

}

func straightLineMapping(
	board components.Board,
	chipsForPlacing []models.ChipForPlacing,
	isPlacedOnBoard bool,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {
	for _, chipForPlacing := range chipsForPlacing {
		// Determine the chip to use (normal or alternative)
		chip := chipForPlacing.Chip
		if chip.ChipType == string(constants.AlterOperatorType) || chip.ChipType == string(constants.BlankType) {
			chip = components.NewChipForCalculating(
				chipForPlacing.SelectedChip.Value,
				chip.Score,
				chipForPlacing.SelectedChip.ChipType,
			)
		}

		// Create a new ChipForCalculating and add it to the set
		*chipForCalculatingSet = append(*chipForCalculatingSet, models.ChipForCalculating{
			SquareType:         board.GetSquare(chipForPlacing.Position).SquareType,
			ChipForCalculating: chip,
			IsPlacedOnBoard:    isPlacedOnBoard,
		})
	}
}

func connectorsIndex(connectors [][2]int) []int {
	var counter []int
	for i := 0; i < len(connectors); i++ {
		if connectors[i] != [2]int{} {
			counter = append(counter, i)
		}
	}
	return counter
}

// Done
func FirstTurnMapping(board components.Board, chipForPlacing []models.ChipForPlacing) []models.ChipForCalculating {
	chipForCalculatingSet := []models.ChipForCalculating{}
	straightLineMapping(board, chipForPlacing, false, &chipForCalculatingSet)
	return chipForCalculatingSet
}

func SingleChipMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	singleChipConnectorSet [][2]int,
) []models.ChipForCalculating {
	chipForCalculating := []models.ChipForCalculating{}

	// Determine the indexes of relevant connectors
	indexes := connectorsIndex(singleChipConnectorSet)

	// Handle single index case
	if len(indexes) == 1 {
		if indexes[0] == 0 {
			leftConnectedMapping(board, chipForPlacing, singleChipConnectorSet[0], &chipForCalculating)
		}
	}

	return chipForCalculating
}

// Done
func leftConnectedMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	chipConnector [2]int,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {
	// Map existing chips in a straight line
	straightLineMapping(board, chipForPlacing, false, chipForCalculatingSet)

	currentPosition, fixedCoord := chipConnector[0], chipConnector[1]

	var chipsForAdding []models.ChipForPlacing

	// Traverse left and detect chips
	for currentPosition <= 15 {
		position := [2]int{currentPosition, fixedCoord}
		square := board.GetSquare(position)

		if square.ChipPlaceOn.IsEmpty() {
			break
		}

		chipsForAdding = append(chipsForAdding, models.ChipForPlacing{
			Position:     square.Position,
			Chip:         square.ChipPlaceOn,
			SelectedChip: components.Chip{}, // Empty chip as placeholder
		})

		currentPosition++
	}

	// Map the detected chips in a straight line
	straightLineMapping(board, chipsForAdding, true, chipForCalculatingSet)
}

// func SingleChipMapping(board components.Board, SingleChipConnectorSet [][2]int) []models.ChipForCalculating {
// 	chipForCalculating := []models.ChipForCalculating{}

// 	return chipForCalculating
// }
