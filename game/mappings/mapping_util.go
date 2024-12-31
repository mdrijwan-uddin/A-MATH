package mappings

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
)

func straightLineMapping(
	board components.Board,
	chipsForPlacing []models.ChipForPlacing,
	isPlacedOnBoard bool,
	chipForCalculatingSet *[]models.ChipForCalculating,
) {
	for _, chips := range chipsForPlacing {
		// Determine the chip to use (normal or alternative)
		var chip components.Chip
		if isPlacedOnBoard {
			chip = chips.Chip
		} else {
			chip = mapChipForCalculating(chips)
		}

		// Create a new ChipForCalculating and add it to the set
		*chipForCalculatingSet = append(*chipForCalculatingSet, models.ChipForCalculating{
			SquareType:      board.GetSquare(chips.Coordinate).SquareType,
			Chip:            chip,
			IsPlacedOnBoard: isPlacedOnBoard,
		})
	}
}

func mapChipForCalculating(c models.ChipForPlacing) components.Chip {
	chip := c.Chip
	if chip.ChipType == string(constants.AlterOperatorType) || chip.ChipType == string(constants.BlankType) {
		chip = components.NewChipForCalculating(
			c.SelectedChip.Value,
			chip.Score,
			c.SelectedChip.ChipType,
		)
	}

	return chip
}

func directionsMapping(
	board components.Board,
	connector models.ChipConnector,
	chipForCalculatingSet *[]models.ChipForCalculating,
	directions string,
) {
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
	var chipsOnBoard []models.ChipForPlacing
	for i := 0; i < counter; i++ {
		position := [2]int{
			currentPosition[0] + i*delta[0] + i*delta[2],
			currentPosition[1] + i*delta[1] + i*delta[3],
		}
		square := board.GetSquare(position)

		chipsOnBoard = append(chipsOnBoard, models.ChipForPlacing{
			Coordinate:   square.Coordinate,
			Chip:         square.ChipPlaceOn,
			SelectedChip: components.Chip{}, // Empty chip as placeholder
		})
	}

	// Map the detected chips in a straight line
	straightLineMapping(board, chipsOnBoard, true, chipForCalculatingSet)
}
