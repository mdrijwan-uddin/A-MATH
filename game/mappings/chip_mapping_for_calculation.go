package mappings

import (
	"A-MATH/game/components"
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

func straightLineMapping(board components.Board, chipForPlacing []models.ChipForPlacing, isPlacedOnBoard bool, chipForCalculatingSet *[]models.ChipForCalculating) {
	for _, ch := range chipForPlacing {
		chipForCalculating := models.ChipForCalculating{
			SquareType:      board.GetSquare(ch.Position).SquareType,
			Chip:            ch.Chip,
			IsPlacedOnBoard: isPlacedOnBoard,
		}
		*chipForCalculatingSet = append(*chipForCalculatingSet, chipForCalculating)
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

func FirstTurnMapping(board components.Board, chipForPlacing []models.ChipForPlacing) []models.ChipForCalculating {
	chipForCalculatingSet := []models.ChipForCalculating{}
	straightLineMapping(board, chipForPlacing, false, &chipForCalculatingSet)
	return chipForCalculatingSet
}

func SingleChipMapping(board components.Board, SingleChipConnectorSet [][2]int) []models.ChipForCalculating {
	chipForCalculating := []models.ChipForCalculating{}

	index := connectorsIndex(SingleChipConnectorSet)

	if len(index) == 1 {
		if index[0] == 0 {
			leftConnectedMapping(board, SingleChipConnectorSet[0], &chipForCalculating)
		}
	}

	return chipForCalculating
}

func leftConnectedMapping(board components.Board, ChipConnector [2]int, chipForCalculatingSet *[]models.ChipForCalculating) {

}

// func SingleChipMapping(board components.Board, SingleChipConnectorSet [][2]int) []models.ChipForCalculating {
// 	chipForCalculating := []models.ChipForCalculating{}

// 	return chipForCalculating
// }
