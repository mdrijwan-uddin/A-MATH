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
	isStraightLine bool,
) [][]models.ChipForCalculating {

	chipForCalculatingTable := [][]models.ChipForCalculating{}

	if isVertical { // Handle Vertical mapping (Up and Down connections)
		chipForCalculatingTable = append(chipForCalculatingTable, mapVertical(board, chipForPlacing, straightConnector, isStraightLine)...)
	} else { // Handle Horizontal mapping (Left and Right connections)
		chipForCalculatingTable = append(chipForCalculatingTable, mapHorizontal(board, chipForPlacing, straightConnector, isStraightLine)...)
	}

	// Handle Mapping for edge connectors
	chipForCalculatingTable = append(chipForCalculatingTable, handleEdgeConnectors(board, chipForPlacing, straightConnector, isVertical)...)

	return chipForCalculatingTable
}

func mapVertical(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	straightConnector []models.ChipConnector,
	isStraightLine bool,
) [][]models.ChipForCalculating {

	var chipForCalculatingSet []models.ChipForCalculating
	firstIndex := 0
	lastIndex := len(straightConnector) - 1

	if straightConnector[firstIndex].TopConnected {
		directionsMapping(board, straightConnector[firstIndex], &chipForCalculatingSet, string(constants.Up))
	}

	if isStraightLine {
		straightLineMapping(board, chipForPlacing, false, &chipForCalculatingSet)
	} else {
		crossMappingSeperateChipForPlacing(board, chipForPlacing, straightConnector, true, &chipForCalculatingSet)
	}

	if straightConnector[lastIndex].BottomConnected {
		directionsMapping(board, straightConnector[lastIndex], &chipForCalculatingSet, string(constants.Down))
	}

	return [][]models.ChipForCalculating{chipForCalculatingSet}
}

func mapHorizontal(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	straightConnector []models.ChipConnector,
	isStraightLine bool,
) [][]models.ChipForCalculating {

	var chipForCalculatingSet []models.ChipForCalculating
	firstIndex := 0
	lastIndex := len(straightConnector) - 1

	if straightConnector[firstIndex].LeftConnected {
		directionsMapping(board, straightConnector[firstIndex], &chipForCalculatingSet, string(constants.Left))
	}

	if isStraightLine {
		straightLineMapping(board, chipForPlacing, false, &chipForCalculatingSet)
	} else {
		crossMappingSeperateChipForPlacing(board, chipForPlacing, straightConnector, false, &chipForCalculatingSet)
	}

	if straightConnector[lastIndex].RightConnected {
		directionsMapping(board, straightConnector[lastIndex], &chipForCalculatingSet, string(constants.Right))
	}

	return [][]models.ChipForCalculating{chipForCalculatingSet}
}

func handleEdgeConnectors(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	straightConnector []models.ChipConnector,
	isVertical bool,
) [][]models.ChipForCalculating {

	var chipForCalculatingTable [][]models.ChipForCalculating

	for i, connector := range straightConnector {
		// Check if there is any edge connection for the current connector
		edgeConnector := connector.EdgeConnector(isVertical)

		if edgeConnector.TotalConnecter() > 0 {
			// Mapping for individual chips at the edge connectors
			newChipForPlacing := []models.ChipForPlacing{chipForPlacing[i]}
			newConnector := []models.ChipConnector{straightConnector[i]}
			// Perform single chip mapping for edge connectors
			chipForCalculatingTable = append(chipForCalculatingTable,
				SingleChipMapping(board, newChipForPlacing, newConnector)...)
		}
	}

	// Return the final chip calculation table after all mappings
	return chipForCalculatingTable
}

func crossMappingSeperateChipForPlacing(
	board components.Board,
	chipForPlacing []models.ChipForPlacing,
	connector []models.ChipConnector,
	isVertical bool,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {
	start := 0

	// Iterate through the chips, excluding the last one for comparison
	for j := 0; j < len(chipForPlacing)-1; j++ {
		var placeChipsRange int
		var delta [2]int
		// Check for valid vertical connection between chips
		if isVertical && connector[j].BottomConnected && connector[j+1].TopConnected {
			placeChipsRange = chipForPlacing[j+1].Position[1] - chipForPlacing[j].Position[1] - 1
			delta = [2]int{0, 1}
		}

		// Check for valid horizontal connection between chips
		if !isVertical && connector[j].RightConnected && connector[j+1].LeftConnected {
			placeChipsRange = chipForPlacing[j+1].Position[0] - chipForPlacing[j].Position[0] - 1
			delta = [2]int{1, 0}
		}

		// If no valid connection, continue to the next iteration
		if placeChipsRange <= 0 {
			continue
		}

		// Map the starting placed chips to the board
		straightLineMapping(board, chipForPlacing[start:j+1], false, chipForCalculatingSet)
		start = j + 1

		// Collect chips along the calculated line
		var chipsOnBoard []models.ChipForPlacing

		for i := 0; i < placeChipsRange; i++ {
			currentPosition := chipForPlacing[j].Position
			needPosition := [2]int{
				currentPosition[0] + (1+i)*delta[0],
				currentPosition[1] + (1+i)*delta[1],
			}

			// Get the square from the board based on the calculated position
			square := board.GetSquare(needPosition)

			// Add the chip for placement along this line
			chipsOnBoard = append(chipsOnBoard, models.ChipForPlacing{
				Position:     square.Position,
				Chip:         square.ChipPlaceOn,
				SelectedChip: components.Chip{}, // Empty chip placeholder
			})
		}

		// Map the existing placed chips to the board
		straightLineMapping(board, chipsOnBoard, true, chipForCalculatingSet)
	}

	// Map the remaining chips after the last placed one
	straightLineMapping(board, chipForPlacing[start:], false, chipForCalculatingSet)
}
