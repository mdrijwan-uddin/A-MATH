package mappings

import "A-MATH/game/components"

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
