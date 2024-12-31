package mappings

import (
	"A-MATH/err"
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"A-MATH/game/rules"
)

func Management(board components.Board, chips []models.ChipForPlacing) ([][]models.ChipForCalculating, error) {
	// If no connector is returned, invalid placement
	if len(chips) == 0 {
		return nil, err.New(int(constants.BadRequest), string(constants.InvalidChipPlacement))
	}

	// Gather coordinates
	var coords [][2]int
	for _, chip := range chips {
		coords = append(coords, chip.Coordinate)
	}

	// Validate placement
	isVertical, isStraightLine, err := rules.ValidateChipPlacement(board, coords)
	if err != nil {
		return nil, err
	}

	// Determine connector
	var connector []models.ChipConnector
	switch {
	case len(chips) == 1:
		connector = SingleChipConnector(board, coords[0])
	case isStraightLine:
		connector = StraightConnector(board, coords, isVertical)
	default:
		connector = CrossConnector(board, coords, isVertical)
	}

	// Handle mapping
	mappedChips := chipMappingHandler(board, chips, connector, isVertical, isStraightLine)
	return mappedChips, nil
}

// chipMappingHandler determines how to map chips based on the board state.
func chipMappingHandler(
	board components.Board,
	chips []models.ChipForPlacing,
	connector []models.ChipConnector,
	isVertical bool,
	isStraightLine bool,
) [][]models.ChipForCalculating {

	// If board is empty, do first-turn mapping
	if board.IsEmpty() {
		return FirstTurnMapping(board, chips)
	}

	// Single chip or multiple chips
	if len(chips) == 1 {
		return SingleChipMapping(board, chips, connector)
	}

	return StraightMapping(board, chips, connector, isVertical, isStraightLine)
}
