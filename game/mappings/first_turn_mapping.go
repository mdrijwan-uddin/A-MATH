package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/models"
)

func FirstTurnMapping(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
) [][]models.ChipForCalculating {

	chipForCalculatingSet := [][]models.ChipForCalculating{}
	chipForCalculating := []models.ChipForCalculating{}
	straightLineMapping(board, chipForPlacing, false, &chipForCalculating)
	chipForCalculatingSet = append(chipForCalculatingSet, chipForCalculating)
	return chipForCalculatingSet
}
