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

// func SingleChipMapping(board components.Board, SingleChipConnectorSet [][2]int) []models.ChipForCalculating {
// 	chipForCalculating := []models.ChipForCalculating{}

// 	return chipForCalculating
// }
