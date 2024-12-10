package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
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

		firstConnector, secondConnector := connector.TemporarySeperateConnector()

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

	firstConnector, secondConnector := connector.TemporarySeperateConnector()

	singleChipTwoDirectionMapping(board, chipForPlacing, firstConnector, chipForCalculatingSet)
	singleChipTwoDirectionMapping(board, chipForPlacing, secondConnector, chipForCalculatingSet)
}
