package mappings_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"testing"
)

func SingleDirectionChipMapping() []models.ChipForPlacing {
	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{2, 8}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{3, 8}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 8}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 8}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("14")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 8}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 8}, components.NewChip("19")))
	chips = append(chips, models.NewChipForPlacing([2]int{13, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{14, 8}, components.NewChip("3")))

	chips = append(chips, models.NewChipForPlacing([2]int{8, 2}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 3}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 4}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 5}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 11}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 12}, components.NewChip("19")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 13}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 14}, components.NewChip("3")))
	return chips
}

func mockResultValueAppendFirst() []string {
	return []string{
		string(constants.Subtraction),
		string(constants.Subtraction),
		string(constants.Seven),
		string(constants.Addition),
		string(constants.Four),
		string(constants.Equal),
		string(constants.One),
		string(constants.Fourteen),
		string(constants.Subtraction),
		string(constants.Five),
		string(constants.Multiply),
		string(constants.Nineteen),
		string(constants.Equal),
		string(constants.Three),
	}
}

func mockResultValueAppendLast() []string {
	return []string{
		string(constants.Subtraction),
		string(constants.Seven),
		string(constants.Addition),
		string(constants.Four),
		string(constants.Equal),
		string(constants.One),
		string(constants.Fourteen),
		string(constants.Subtraction),
		string(constants.Five),
		string(constants.Multiply),
		string(constants.Nineteen),
		string(constants.Equal),
		string(constants.Three),
		string(constants.Subtraction),
	}
}

func mockIsPlaceOnBoardTrueFirst() []bool {
	return []bool{false, true, true, true, true, true, true, true, true, true, true, true, true, true}
}

func mockIsPlaceOnBoardTrueLast() []bool {
	return []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false}
}

func TestSingleChipMapping(t *testing.T) {
	boardOneDirectionTest := components.NewBoard()
	chipForBoard := SingleDirectionChipMapping()

	for _, ch := range chipForBoard {
		boardOneDirectionTest.Add(ch.Position, ch.Chip)
	}

	tests := []struct {
		name                   string
		board                  components.Board
		position               [2]int
		value                  string
		expectedIsPlaceOnBoard []bool
		expectedValue          []string
	}{
		{
			name:                   "Correct Right Connected Mapping",
			board:                  boardOneDirectionTest,
			position:               [2]int{1, 8},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardTrueFirst(),
			expectedValue:          mockResultValueAppendFirst(),
		},
		{
			name:                   "Correct Down Connected Mapping",
			board:                  boardOneDirectionTest,
			position:               [2]int{8, 1},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardTrueFirst(),
			expectedValue:          mockResultValueAppendFirst(),
		},
		{
			name:                   "Correct Left Connected Mapping",
			board:                  boardOneDirectionTest,
			position:               [2]int{15, 8},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardTrueLast(),
			expectedValue:          mockResultValueAppendLast(),
		},
		{
			name:                   "Correct Up Connected Mapping",
			board:                  boardOneDirectionTest,
			position:               [2]int{8, 15},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardTrueLast(),
			expectedValue:          mockResultValueAppendLast(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var chipForPlacing []models.ChipForPlacing
			chipForPlacing = append(chipForPlacing, models.NewChipForPlacing(tt.position, components.NewChip(tt.value)))

			connector := mappings.SingleChipConnector(tt.board, tt.position)
			results := mappings.SingleChipMapping(tt.board, chipForPlacing, connector)

			for i, result := range results {
				if result[0].IsPlacedOnBoard != tt.expectedIsPlaceOnBoard[i] {
					t.Errorf("SingleChipMapping() result = result:%v; expected %v", result[0].IsPlacedOnBoard, tt.expectedIsPlaceOnBoard[i])
				}

				if result[0].ChipForCalculating.Value != tt.expectedValue[i] {
					t.Errorf("SingleChipMapping() result = result:%v; expected %v", result[0].ChipForCalculating.Value, tt.expectedValue[i])
				}
			}

		})
	}
}
