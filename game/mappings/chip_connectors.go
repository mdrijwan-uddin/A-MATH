package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/models"
)

func SingleChipConnector(board components.Board, coordinates [2]int) []models.ChipConnector {
	var connectors []models.ChipConnector

	singleConnector := models.NewChipConnector(coordinates)
	singleConnector.CheckAllDirectionConnector(board)

	connectors = append(connectors, singleConnector)
	return connectors
}

func StraightConnector(board components.Board, coordinates [][2]int, isVertical bool) []models.ChipConnector {
	var connectors []models.ChipConnector

	if isVertical {

		firstChipConnector := models.NewChipConnector(coordinates[0])
		firstChipConnector.CheckTopConnector(board)
		firstChipConnector.CheckHorizontalConnector(board)
		connectors = append(connectors, firstChipConnector)

		for i := 1; i < len(coordinates)-1; i++ {
			connector := models.NewChipConnector(coordinates[i])
			connector.CheckHorizontalConnector(board)
			connectors = append(connectors, connector)
		}

		lastChipConnector := models.NewChipConnector(coordinates[len(coordinates)-1])
		lastChipConnector.CheckHorizontalConnector(board)
		lastChipConnector.CheckBottomConnector(board)
		connectors = append(connectors, lastChipConnector)
	}

	if !isVertical { //isHorizontal

		firstChipConnector := models.NewChipConnector(coordinates[0])
		firstChipConnector.CheckLeftConnector(board)
		firstChipConnector.CheckVerticalConnector(board)
		connectors = append(connectors, firstChipConnector)

		for i := 1; i < len(coordinates)-1; i++ {
			connector := models.NewChipConnector(coordinates[i])
			connector.CheckVerticalConnector(board)
			connectors = append(connectors, connector)
		}

		lastChipConnector := models.NewChipConnector(coordinates[len(coordinates)-1])
		lastChipConnector.CheckVerticalConnector(board)
		lastChipConnector.CheckRightConnector(board)
		connectors = append(connectors, lastChipConnector)
	}

	return connectors
}

func CrossConnector(board components.Board, coordinates [][2]int, isVertical bool) []models.ChipConnector {
	var connectors []models.ChipConnector
	splitConnectors := splitConnector(coordinates, isVertical)

	for _, sp := range splitConnectors {
		if len(sp) == 1 {
			connectors = append(connectors, SingleChipConnector(board, sp[0])[0])

		} else {
			length := len(sp)
			straightConnector := StraightConnector(board, sp, isVertical)

			for i := 0; i < length; i++ {
				connectors = append(connectors, straightConnector[i])

			}
		}
	}
	return connectors
}

func splitConnector(coordinates [][2]int, isVertical bool) [][][2]int {
	var currentCoord, nextCoord, fixedCoord int
	var connectors [][2]int
	var connectorsSet [][][2]int

	for i := 0; i < len(coordinates)-1; i++ {
		if isVertical {
			currentCoord, nextCoord = coordinates[i][1], coordinates[i+1][1]
			fixedCoord = coordinates[i][0]

			connectors = append(connectors, [2]int{fixedCoord, currentCoord})
			if currentCoord+1 != nextCoord { // Check for separation
				connectorsSet = append(connectorsSet, connectors)
				connectors = nil //reset connectors
			}
		} else { //isHorizontal
			currentCoord, nextCoord = coordinates[i][0], coordinates[i+1][0]
			fixedCoord = coordinates[i][1]

			connectors = append(connectors, [2]int{currentCoord, fixedCoord})
			if currentCoord+1 != nextCoord { // Check for separation
				connectorsSet = append(connectorsSet, connectors)
				connectors = nil //reset connectors
			}
		}
	}

	// Append the lastconnector
	if isVertical {
		connectors = append(connectors, [2]int{fixedCoord, nextCoord})
	} else {
		connectors = append(connectors, [2]int{nextCoord, fixedCoord})
	}
	connectorsSet = append(connectorsSet, connectors)

	return connectorsSet
}
