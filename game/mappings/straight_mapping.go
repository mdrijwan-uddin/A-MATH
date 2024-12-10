package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
)

func StraightMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	straightConnector []models.ChipConnector,
	isVertical bool,
) [][]models.ChipForCalculating {

	chipForCalculatingSet := [][]models.ChipForCalculating{}
	chipForCalculating := []models.ChipForCalculating{}

	firstIndex := 0
	lastIndex := len(straightConnector) - 1

	if isVertical {
		if straightConnector[firstIndex].TopConnected {
			directionsMapping(board, straightConnector[firstIndex], &chipForCalculating, string(constants.Up))
		}

		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)

		if straightConnector[lastIndex].BottomConnected {
			directionsMapping(board, straightConnector[lastIndex], &chipForCalculating, string(constants.Down))
		}

	} else { // isHorizontal
		if straightConnector[firstIndex].LeftConnected {
			directionsMapping(board, straightConnector[firstIndex], &chipForCalculating, string(constants.Left))
		}

		straightLineMapping(board, chipForPlacing, false, &chipForCalculating)

		if straightConnector[lastIndex].RightConnected {
			directionsMapping(board, straightConnector[lastIndex], &chipForCalculating, string(constants.Right))
		}

	}
	chipForCalculatingSet = append(chipForCalculatingSet, chipForCalculating)
	chipForCalculating = nil //reset

	for i := 0; i < lastIndex; i++ {
		if straightConnector[i].TotalConnecter() > 0 {
			newChipForPlacing := []models.ChipForPlacing{chipForPlacing[i]}
			newConnector := []models.ChipConnector{straightConnector[i]}

			newChipForCalculating := SingleChipMapping(board, newChipForPlacing, newConnector)
			chipForCalculatingSet = append(chipForCalculatingSet, newChipForCalculating...)
		}
	}

	return chipForCalculatingSet
}