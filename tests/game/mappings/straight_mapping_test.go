package mappings_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"testing"
)

func TestStraightMapping(t *testing.T) {
	tests := []struct {
		name                   string
		board                  components.Board
		position               [2]int
		value                  string
		expectedIsPlaceOnBoard [][]bool
		expectedValue          [][]string
	}{
		{
			name:                   "(1D) Right Connected Mapping",
			board:                  SingleDirectionChipMapping(),
			position:               [2]int{1, 8},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseFirst14(),
			expectedValue:          mockResultValueAppendFirst1D(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var chipForPlacing []models.ChipForPlacing
			chipForPlacing = append(chipForPlacing, models.NewChipForPlacing(tt.position, components.NewChip(tt.value)))

			connector := mappings.SingleChipConnector(tt.board, tt.position)
			results := mappings.SingleChipMapping(tt.board, chipForPlacing, connector)

			testStraightMapChecking(t, results, tt.expectedIsPlaceOnBoard, tt.expectedValue)

		})
	}
}

func testStraightMapChecking(t *testing.T, singleChipMapping [][]models.ChipForCalculating, expectedIsPlaceOnBoard [][]bool, expectedValue [][]string) {
	for j, resultSet := range singleChipMapping {
		for i, result := range resultSet {
			if result.IsPlacedOnBoard != expectedIsPlaceOnBoard[j][i] {
				t.Errorf("SingleChipMapping().IsPlacedOnBoard result[%d] = result:%v; expected %v", i, result.IsPlacedOnBoard, expectedIsPlaceOnBoard[j][i])
			}

			if result.ChipForCalculating.Value != expectedValue[j][i] {
				t.Errorf("SingleChipMapping().Value result[%d] = result:%v; expected %v", i, result.ChipForCalculating.Value, expectedValue[j][i])
			}
		}
	}
}
