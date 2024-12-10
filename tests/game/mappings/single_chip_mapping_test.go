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

func mockResultValueAppendFirst_1D() []string {
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

func mockResultValueAppendLast_1D() []string {
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

func mockIsPlaceOnBoardFalseFirst_14() []bool {
	return []bool{false, true, true, true, true, true, true, true, true, true, true, true, true, true}
}

func mockIsPlaceOnBoardFalseLast_14() []bool {
	return []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false}
}

func TestOneDirectionMapping(t *testing.T) {
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
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseFirst_14(),
			expectedValue:          mockResultValueAppendFirst_1D(),
		},
		{
			name:                   "Correct Bottom Connected Mapping",
			board:                  boardOneDirectionTest,
			position:               [2]int{8, 1},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseFirst_14(),
			expectedValue:          mockResultValueAppendFirst_1D(),
		},
		{
			name:                   "Correct Left Connected Mapping",
			board:                  boardOneDirectionTest,
			position:               [2]int{15, 8},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseLast_14(),
			expectedValue:          mockResultValueAppendLast_1D(),
		},
		{
			name:                   "Correct Top Connected Mapping",
			board:                  boardOneDirectionTest,
			position:               [2]int{8, 15},
			value:                  string(constants.Subtraction),
			expectedIsPlaceOnBoard: mockIsPlaceOnBoardFalseLast_14(),
			expectedValue:          mockResultValueAppendLast_1D(),
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

func TwoDirectionChipMapping_1() []models.ChipForPlacing {
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
	return chips
}

func TwoDirectionChipMapping_2() []models.ChipForPlacing {
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
	return chips
}

func mockIsPlaceOnBoardFalseFirst_6() []bool {
	return []bool{false, true, true, true, true, true}
}

func mockIsPlaceOnBoardFalseLast_6() []bool {
	return []bool{true, true, true, true, true, false}
}

func TestTwoDirectionMapping(t *testing.T) {
	boardTwoDirectionTest_1 := components.NewBoard()
	chipForBoard := TwoDirectionChipMapping_1()

	for _, ch := range chipForBoard {
		boardTwoDirectionTest_1.Add(ch.Position, ch.Chip)
	}

	tests := []struct {
		name                     string
		board                    components.Board
		position                 [2]int
		value                    string
		expectedIsPlaceOnBoard_0 []bool
		expectedValue_0          []string
		expectedIsPlaceOnBoard_1 []bool
		expectedValue_1          []string
	}{
		{
			name:                     "Correct Top-Right Connected Mapping",
			board:                    boardTwoDirectionTest_1,
			position:                 [2]int{5, 5},
			value:                    string(constants.Equal),
			expectedIsPlaceOnBoard_0: mockIsPlaceOnBoardFalseFirst_6(),
			expectedValue_0: []string{
				string(constants.Equal),
				string(constants.Eighteen),
				string(constants.Equal),
				string(constants.Four),
				string(constants.Addition),
				string(constants.Fourteen),
			},
			expectedIsPlaceOnBoard_1: mockIsPlaceOnBoardFalseFirst_6(),
			expectedValue_1: []string{
				string(constants.Equal),
				string(constants.Five),
				string(constants.Equal),
				string(constants.Five),
				string(constants.Multiply),
				string(constants.One),
			},
		},
		// {
		// 	name:                     "Correct Bottom-Right Connected Mapping",
		// 	board:                    boardTwoDirectionTest_1,
		// 	position:                 [2]int{5, 11},
		// 	value:                    string(constants.Equal),
		// 	expectedIsPlaceOnBoard_0: mockIsPlaceOnBoardFalseLast_6(),
		// 	expectedValue_0: []string{
		// 		string(constants.Five),
		// 		string(constants.Equal),
		// 		string(constants.Five),
		// 		string(constants.Multiply),
		// 		string(constants.One),
		// 		string(constants.Equal),
		// 	},
		// 	expectedIsPlaceOnBoard_1: mockIsPlaceOnBoardFalseFirst_6(),
		// 	expectedValue_1: []string{
		// 		string(constants.Equal),
		// 		string(constants.Zero),
		// 		string(constants.Equal),
		// 		string(constants.Two),
		// 		string(constants.Multiply),
		// 		string(constants.Two),
		// 	},
		// },
		// {
		// 	name:                     "Correct Top-Left Connected Mapping",
		// 	board:                    boardTwoDirectionTest_1,
		// 	position:                 [2]int{11, 5},
		// 	value:                    string(constants.Equal),
		// 	expectedIsPlaceOnBoard_0: mockIsPlaceOnBoardFalseLast_6(),
		// 	expectedValue_0: []string{
		// 		string(constants.Eighteen),
		// 		string(constants.Equal),
		// 		string(constants.Four),
		// 		string(constants.Addition),
		// 		string(constants.Fourteen),
		// 		string(constants.Equal),
		// 	},
		// 	expectedIsPlaceOnBoard_1: mockIsPlaceOnBoardFalseFirst_6(),
		// 	expectedValue_1: []string{
		// 		string(constants.Equal),
		// 		string(constants.Thirteen),
		// 		string(constants.Subtraction),
		// 		string(constants.Seven),
		// 		string(constants.Equal),
		// 		string(constants.Six),
		// 	},
		// },
		// {
		// 	name:                     "Correct Bottom-Left Connected Mapping",
		// 	board:                    boardTwoDirectionTest_1,
		// 	position:                 [2]int{11, 11},
		// 	value:                    string(constants.Equal),
		// 	expectedIsPlaceOnBoard_0: mockIsPlaceOnBoardFalseLast_6(),
		// 	expectedValue_0: []string{
		// 		string(constants.Zero),
		// 		string(constants.Equal),
		// 		string(constants.Two),
		// 		string(constants.Multiply),
		// 		string(constants.Two),
		// 		string(constants.Equal),
		// 	},
		// 	expectedIsPlaceOnBoard_1: mockIsPlaceOnBoardFalseLast_6(),
		// 	expectedValue_1: []string{
		// 		string(constants.Thirteen),
		// 		string(constants.Subtraction),
		// 		string(constants.Seven),
		// 		string(constants.Equal),
		// 		string(constants.Six),
		// 		string(constants.Equal),
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var chipForPlacing []models.ChipForPlacing
			chipForPlacing = append(chipForPlacing, models.NewChipForPlacing(tt.position, components.NewChip(tt.value)))

			connectors := mappings.SingleChipConnector(tt.board, tt.position)
			results := mappings.SingleChipMapping(tt.board, chipForPlacing, connectors)

			for _, connector := range connectors {
				if connector.TotalConnecter() != 2 {
					t.Errorf("SingleChipMapping() connector = result:%v; expected %v", connector.TotalConnecter(), 2)
				}
			}

			for i, result := range results {
				if result[0].IsPlacedOnBoard != tt.expectedIsPlaceOnBoard_0[i] {
					t.Errorf("SingleChipMapping() result[%d] = result:%v; expected %v", i, result[0].IsPlacedOnBoard, tt.expectedIsPlaceOnBoard_0[i])
				}

				if result[0].ChipForCalculating.Value != tt.expectedValue_0[i] {
					t.Errorf("SingleChipMapping() result[%d] = result:%v; expected %v", i, result[0].ChipForCalculating.Value, tt.expectedValue_0[i])
				}

				if result[1].IsPlacedOnBoard != tt.expectedIsPlaceOnBoard_1[i] {
					t.Errorf("SingleChipMapping() result[%d] = result:%v; expected %v", i, result[1].IsPlacedOnBoard, tt.expectedIsPlaceOnBoard_1[i])
				}

				if result[1].ChipForCalculating.Value != tt.expectedValue_1[i] {
					t.Errorf("SingleChipMapping() result[%d] = result:%v; expected %v", i, result[1].ChipForCalculating.Value, tt.expectedValue_1[i])
				}
			}

		})
	}
}
