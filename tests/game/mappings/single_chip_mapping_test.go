package mappings_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"testing"
)

func TestOneDirectionMapping(t *testing.T) {
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
			board:                  mockSingleDirectionChipMapping(),
			position:               [2]int{1, 8},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseFirst14(),
			expectedValue:          mockResultValueAppendFirst1D(),
		},
		{
			name:                   "(1D) Bottom Connected Mapping",
			board:                  mockSingleDirectionChipMapping(),
			position:               [2]int{8, 1},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseFirst14(),
			expectedValue:          mockResultValueAppendFirst1D(),
		},
		{
			name:                   "(1D) Left Connected Mapping",
			board:                  mockSingleDirectionChipMapping(),
			position:               [2]int{15, 8},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseLast14(),
			expectedValue:          mockResultValueAppendLast1D(),
		},
		{
			name:                   "(1D) Top Connected Mapping",
			board:                  mockSingleDirectionChipMapping(),
			position:               [2]int{8, 15},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseLast14(),
			expectedValue:          mockResultValueAppendLast1D(),
		},
		{
			name:     "(2D) Top-Right Connected Mapping",
			board:    mockTwoDirectionChipMapping1(),
			position: [2]int{5, 5},
			value:    string(constants.Equal),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseFirst6(),
				mockIsPlaceOnBoardFalseFirst6(),
			},
			expectedValue: [][]string{
				{
					string(constants.Equal),
					string(constants.Eighteen),
					string(constants.Equal),
					string(constants.Four),
					string(constants.Addition),
					string(constants.Fourteen),
				},
				{
					string(constants.Equal),
					string(constants.Five),
					string(constants.Equal),
					string(constants.Five),
					string(constants.Multiply),
					string(constants.One),
				},
			},
		},
		{
			name:     "(2D) Bottom-Right Connected Mapping",
			board:    mockTwoDirectionChipMapping1(),
			position: [2]int{5, 11},
			value:    string(constants.Equal),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseLast6(),
				mockIsPlaceOnBoardFalseFirst6(),
			},
			expectedValue: [][]string{
				{
					string(constants.Five),
					string(constants.Equal),
					string(constants.Five),
					string(constants.Multiply),
					string(constants.One),
					string(constants.Equal),
				},
				{
					string(constants.Equal),
					string(constants.Zero),
					string(constants.Equal),
					string(constants.Two),
					string(constants.Multiply),
					string(constants.Two),
				},
			},
		},
		{
			name:     "(2D) Top-Left Connected Mapping",
			board:    mockTwoDirectionChipMapping1(),
			position: [2]int{11, 5},
			value:    string(constants.Equal),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseLast6(),
				mockIsPlaceOnBoardFalseFirst6(),
			},
			expectedValue: [][]string{
				{
					string(constants.Eighteen),
					string(constants.Equal),
					string(constants.Four),
					string(constants.Addition),
					string(constants.Fourteen),
					string(constants.Equal),
				},
				{
					string(constants.Equal),
					string(constants.Thirteen),
					string(constants.Subtraction),
					string(constants.Seven),
					string(constants.Equal),
					string(constants.Six),
				},
			},
		},
		{
			name:     "(2D) Bottom-Left Connected Mapping",
			board:    mockTwoDirectionChipMapping1(),
			position: [2]int{11, 11},
			value:    string(constants.Equal),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseLast6(),
				mockIsPlaceOnBoardFalseLast6(),
			},
			expectedValue: [][]string{
				{
					string(constants.Zero),
					string(constants.Equal),
					string(constants.Two),
					string(constants.Multiply),
					string(constants.Two),
					string(constants.Equal),
				},
				{
					string(constants.Thirteen),
					string(constants.Subtraction),
					string(constants.Seven),
					string(constants.Equal),
					string(constants.Six),
					string(constants.Equal),
				},
			},
		},
		{
			name:     "(2D) Left-Right Horizontal Mapping",
			board:    mockTwoDirectionChipMapping2(),
			position: [2]int{4, 9},
			value:    string(constants.Equal),
			expectedIsPlaceOnBoard: [][]bool{
				{true, true, true, true, true, true, true, false, true},
			},
			expectedValue: [][]string{
				{
					string(constants.Four),
					string(constants.Multiply),
					string(constants.Five),
					string(constants.Equal),
					string(constants.Ten),
					string(constants.Addition),
					string(constants.Ten),
					string(constants.Equal),
					string(constants.Twenty),
				},
			},
		},
		{
			name:     "(2D) Top-Bottom Vertical Mapping",
			board:    mockTwoDirectionChipMapping2(),
			position: [2]int{5, 10},
			value:    string(constants.Equal),
			expectedIsPlaceOnBoard: [][]bool{
				{true, true, true, true, false, true, true, true, true, true, true},
			},
			expectedValue: [][]string{
				{
					string(constants.Two),
					string(constants.Zero),
					string(constants.Equal),
					string(constants.Twenty),
					string(constants.Equal),
					string(constants.Fifteen),
					string(constants.Addition),
					string(constants.Five),
					string(constants.Equal),
					string(constants.Two),
					string(constants.Zero),
				},
			},
		},
		{
			name:     "(3D) Top-Bottom Vertical & Right Connected Mapping",
			board:    mockThreeDirectionChipMapping(),
			position: [2]int{4, 8},
			value:    string(constants.Nineteen),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseMiddle9(),
				mockIsPlaceOnBoardFalseFirst8(),
			},
			expectedValue: [][]string{
				{
					string(constants.Addition),
					string(constants.Equal),
					string(constants.Addition),
					string(constants.Four),
					string(constants.Nineteen),
					string(constants.Thirteen),
					string(constants.Equal),
					string(constants.Fifteen),
					string(constants.Seven),
				},
				{
					string(constants.Nineteen),
					string(constants.Fourteen),
					string(constants.Two),
					string(constants.Eight),
					string(constants.Two),
					string(constants.One),
					string(constants.Twelve),
					string(constants.Zero),
				},
			},
		},
		{
			name:     "(3D) Left-Right Horizontal & Bottom Connected Mapping",
			board:    mockThreeDirectionChipMapping(),
			position: [2]int{8, 4},
			value:    string(constants.Nineteen),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseMiddle9(),
				mockIsPlaceOnBoardFalseFirst8(),
			},
			expectedValue: [][]string{
				{
					string(constants.Addition),
					string(constants.Equal),
					string(constants.Eight),
					string(constants.Nine),
					string(constants.Nineteen),
					string(constants.Seventeen),
					string(constants.One),
					string(constants.Seven),
					string(constants.Seven),
				},
				{
					string(constants.Nineteen),
					string(constants.Subtraction),
					string(constants.Three),
					string(constants.One),
					string(constants.Two),
					string(constants.Nine),
					string(constants.Four),
					string(constants.Division),
				},
			},
		},
		{
			name:     "(3D) Top-Bottom Vertical & Left Connected Mapping",
			board:    mockThreeDirectionChipMapping(),
			position: [2]int{12, 8},
			value:    string(constants.Nineteen),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseMiddle9(),
				mockIsPlaceOnBoardFalseLast8(),
			},
			expectedValue: [][]string{
				{
					string(constants.Seven),
					string(constants.Two),
					string(constants.Seven),
					string(constants.Three),
					string(constants.Nineteen),
					string(constants.Zero),
					string(constants.Division),
					string(constants.Equal),
					string(constants.Four),
				},
				{
					string(constants.Fourteen),
					string(constants.Two),
					string(constants.Eight),
					string(constants.Two),
					string(constants.One),
					string(constants.Twelve),
					string(constants.Zero),
					string(constants.Nineteen),
				},
			},
		},
		{
			name:     "(3D) Left-Right Horizontal & Top Connected Mapping",
			board:    mockThreeDirectionChipMapping(),
			position: [2]int{8, 12},
			value:    string(constants.Nineteen),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseMiddle9(),
				mockIsPlaceOnBoardFalseLast8(),
			},
			expectedValue: [][]string{
				{
					string(constants.Seven),
					string(constants.Equal),
					string(constants.Eight),
					string(constants.One),
					string(constants.Nineteen),
					string(constants.Multiply),
					string(constants.Division),
					string(constants.Six),
					string(constants.Four),
				},
				{
					string(constants.Subtraction),
					string(constants.Three),
					string(constants.One),
					string(constants.Two),
					string(constants.Nine),
					string(constants.Four),
					string(constants.Division),
					string(constants.Nineteen),
				},
			},
		},
		{
			name:     "(4D) Left-Right Horizontal & Top-Bottom Vertical Connected Mapping",
			board:    mockFourDirectionChipMapping(),
			position: [2]int{5, 8},
			value:    string(constants.Eight),
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardFalseMiddle7(),
				mockIsPlaceOnBoardFalseMiddle7(),
			},
			expectedValue: [][]string{
				{
					string(constants.Thirteen),
					string(constants.Equal),
					string(constants.Seven),
					string(constants.Eight),
					string(constants.Subtraction),
					string(constants.Zero),
					string(constants.Nineteen),
				},
				{
					string(constants.Five),
					string(constants.Multiply),
					string(constants.Three),
					string(constants.Eight),
					string(constants.Zero),
					string(constants.Two),
					string(constants.Two),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var chipForPlacing []models.ChipForPlacing
			chipForPlacing = append(chipForPlacing, models.NewChipForPlacing(tt.position, components.NewChip(tt.value)))

			connector := mappings.SingleChipConnector(tt.board, tt.position)
			results := mappings.SingleChipMapping(tt.board, chipForPlacing, connector)

			testSingleChipMapChecking(t, results, tt.expectedIsPlaceOnBoard, tt.expectedValue)

		})
	}
}

func testSingleChipMapChecking(t *testing.T, singleChipMapping [][]models.ChipForCalculating, expectedIsPlaceOnBoard [][]bool, expectedValue [][]string) {
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

func mockSingleDirectionChipMapping() components.Board {
	board := components.NewBoard()

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

	for _, ch := range chips {
		board.Add(ch.Position, ch.Chip)
	}
	return board
}

func mockTwoDirectionChipMapping1() components.Board {
	board := components.NewBoard()

	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{5, 8}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 8}, components.NewChip("7")))

	chips = append(chips, models.NewChipForPlacing([2]int{8, 5}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("9")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 11}, components.NewChip("2")))

	chips = append(chips, models.NewChipForPlacing([2]int{5, 6}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 7}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 9}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 10}, components.NewChip("1")))

	chips = append(chips, models.NewChipForPlacing([2]int{11, 6}, components.NewChip("13")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 7}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 9}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 10}, components.NewChip("6")))

	chips = append(chips, models.NewChipForPlacing([2]int{6, 5}, components.NewChip("18")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 5}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 5}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 5}, components.NewChip("14")))

	chips = append(chips, models.NewChipForPlacing([2]int{6, 11}, components.NewChip("0")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 11}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 11}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 11}, components.NewChip("2")))

	for _, ch := range chips {
		board.Add(ch.Position, ch.Chip)
	}
	return board
}

func mockTwoDirectionChipMapping2() components.Board {
	board := components.NewBoard()

	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{2, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{3, 8}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 8}, components.NewChip("10")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 8}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("17")))

	chips = append(chips, models.NewChipForPlacing([2]int{2, 6}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 7}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 9}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 10}, components.NewChip("0")))

	chips = append(chips, models.NewChipForPlacing([2]int{4, 2}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 3}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 4}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 5}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 6}, components.NewChip("10")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 7}, components.NewChip("+")))

	chips = append(chips, models.NewChipForPlacing([2]int{8, 5}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("5")))

	chips = append(chips, models.NewChipForPlacing([2]int{1, 10}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{3, 10}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 10}, components.NewChip("20")))

	chips = append(chips, models.NewChipForPlacing([2]int{6, 10}, components.NewChip("15")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 10}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 10}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 10}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 10}, components.NewChip("0")))

	for _, ch := range chips {
		board.Add(ch.Position, ch.Chip)
	}
	return board
}

func mockThreeDirectionChipMapping() components.Board {
	board := components.NewBoard()

	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{5, 8}, components.NewChip("14")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("8")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("12")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 8}, components.NewChip("0")))

	chips = append(chips, models.NewChipForPlacing([2]int{8, 5}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("3")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("9")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 11}, components.NewChip("/")))

	chips = append(chips, models.NewChipForPlacing([2]int{4, 4}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 5}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 6}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 7}, components.NewChip("4")))

	chips = append(chips, models.NewChipForPlacing([2]int{4, 9}, components.NewChip("13")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 10}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 11}, components.NewChip("15")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 12}, components.NewChip("7")))

	chips = append(chips, models.NewChipForPlacing([2]int{12, 4}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 5}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 6}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 7}, components.NewChip("3")))

	chips = append(chips, models.NewChipForPlacing([2]int{12, 9}, components.NewChip("0")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 10}, components.NewChip("/")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 11}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 12}, components.NewChip("4")))

	chips = append(chips, models.NewChipForPlacing([2]int{5, 4}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 4}, components.NewChip("8")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 4}, components.NewChip("9")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 4}, components.NewChip("17")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 4}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 4}, components.NewChip("7")))

	chips = append(chips, models.NewChipForPlacing([2]int{5, 6}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 6}, components.NewChip("6")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 6}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 6}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 6}, components.NewChip("11")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 6}, components.NewChip("0")))

	chips = append(chips, models.NewChipForPlacing([2]int{5, 10}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 10}, components.NewChip("3")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 10}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 10}, components.NewChip("9")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 10}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 10}, components.NewChip("/")))

	chips = append(chips, models.NewChipForPlacing([2]int{5, 12}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 12}, components.NewChip("8")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 12}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 12}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 12}, components.NewChip("/")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 12}, components.NewChip("6")))

	for _, ch := range chips {
		board.Add(ch.Position, ch.Chip)
	}
	return board
}

func mockFourDirectionChipMapping() components.Board {
	board := components.NewBoard()

	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("19")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("3")))

	chips = append(chips, models.NewChipForPlacing([2]int{2, 6}, components.NewChip("6")))
	chips = append(chips, models.NewChipForPlacing([2]int{3, 6}, components.NewChip("20")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 6}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 6}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 6}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 6}, components.NewChip("12")))

	chips = append(chips, models.NewChipForPlacing([2]int{2, 10}, components.NewChip("9")))
	chips = append(chips, models.NewChipForPlacing([2]int{3, 10}, components.NewChip("3")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 10}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 10}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 10}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 10}, components.NewChip("/")))

	chips = append(chips, models.NewChipForPlacing([2]int{2, 7}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 8}, components.NewChip("13")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 9}, components.NewChip("*")))

	chips = append(chips, models.NewChipForPlacing([2]int{3, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 8}, components.NewChip("7")))

	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("0")))

	chips = append(chips, models.NewChipForPlacing([2]int{5, 5}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 7}, components.NewChip("3")))

	chips = append(chips, models.NewChipForPlacing([2]int{5, 9}, components.NewChip("0")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 11}, components.NewChip("2")))

	for _, ch := range chips {
		board.Add(ch.Position, ch.Chip)
	}
	return board
}

func mockIsPlaceOnBoardFalseFirst14() [][]bool {
	return [][]bool{{false, true, true, true, true, true, true, true, true, true, true, true, true, true}}
}

func mockIsPlaceOnBoardFalseLast14() [][]bool {
	return [][]bool{{true, true, true, true, true, true, true, true, true, true, true, true, true, false}}
}

func mockIsPlaceOnBoardFalseFirst6() []bool {
	return []bool{false, true, true, true, true, true}
}

func mockIsPlaceOnBoardFalseLast6() []bool {
	return []bool{true, true, true, true, true, false}
}

func mockIsPlaceOnBoardFalseMiddle9() []bool {
	return []bool{true, true, true, true, false, true, true, true, true}
}

func mockIsPlaceOnBoardFalseFirst8() []bool {
	return []bool{false, true, true, true, true, true, true, true}
}

func mockIsPlaceOnBoardFalseLast8() []bool {
	return []bool{true, true, true, true, true, true, true, false}
}

func mockIsPlaceOnBoardFalseMiddle7() []bool {
	return []bool{true, true, true, false, true, true, true}
}

func mockResultValueAppendFirst1D() [][]string {
	return [][]string{
		{
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
		},
	}
}

func mockResultValueAppendLast1D() [][]string {
	return [][]string{
		{
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
		},
	}
}
