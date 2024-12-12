package mappings_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"A-MATH/game/rules"
	"testing"
)

func TestStraightMapping(t *testing.T) {
	tests := []struct {
		name                   string
		board                  components.Board
		coordinates            [][2]int
		values                 []string
		expectedIsPlaceOnBoard [][]bool
		expectedValue          [][]string
		expectedSquareType     [][]string
	}{
		{
			name:        "(Straight) Left Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{9, 6}, {10, 6}, {11, 6}, {12, 6}, {13, 6}},
			values: []string{
				string(constants.Addition),
				string(constants.Four),
				string(constants.Equal),
				string(constants.One),
				string(constants.Zero),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueFirst6(),
			},
			expectedValue: [][]string{
				{
					string(constants.Six),
					string(constants.Addition),
					string(constants.Four),
					string(constants.Equal),
					string(constants.One),
					string(constants.Zero),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected1(),
			},
		},
		{
			name:        "(Straight) Top Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{10, 9}, {10, 10}, {10, 11}, {10, 12}, {10, 13}},
			values: []string{
				string(constants.Three),
				string(constants.Subtraction),
				string(constants.Eighteen),
				string(constants.Equal),
				string(constants.Fifteen),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueFirst6(),
			},
			expectedValue: [][]string{
				{
					string(constants.Three),
					string(constants.Three),
					string(constants.Subtraction),
					string(constants.Eighteen),
					string(constants.Equal),
					string(constants.Fifteen),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected1(),
			},
		},
		{
			name:        "(Straight) Right Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}},
			values: []string{
				string(constants.One),
				string(constants.One),
				string(constants.Subtraction),
				string(constants.Six),
				string(constants.Equal),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueLast6(),
			},
			expectedValue: [][]string{
				{
					string(constants.One),
					string(constants.One),
					string(constants.Subtraction),
					string(constants.Six),
					string(constants.Equal),
					string(constants.Five),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected2(),
			},
		},
		{
			name:        "(Straight) Bottom Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{6, 3}, {6, 4}, {6, 5}, {6, 6}, {6, 7}},
			values: []string{
				string(constants.Eleven),
				string(constants.Multiply),
				string(constants.Three),
				string(constants.Equal),
				string(constants.Three),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueLast6(),
			},
			expectedValue: [][]string{
				{
					string(constants.Eleven),
					string(constants.Multiply),
					string(constants.Three),
					string(constants.Equal),
					string(constants.Three),
					string(constants.Three),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected2(),
			},
		},
		{
			name:        "(Straight) Left Extend Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{2, 8}, {3, 8}, {4, 8}, {5, 8}},
			values: []string{
				string(constants.Twenty),
				string(constants.Subtraction),
				string(constants.Seventeen),
				string(constants.Equal),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueLast9(),
			},
			expectedValue: [][]string{
				{
					string(constants.Twenty),
					string(constants.Subtraction),
					string(constants.Seventeen),
					string(constants.Equal),
					string(constants.Three),
					string(constants.Multiply),
					string(constants.One),
					string(constants.Equal),
					string(constants.Three),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected3(),
			},
		},
		{
			name:        "(Straight) Top Extend Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{8, 2}, {8, 3}, {8, 4}, {8, 5}},
			values: []string{
				string(constants.Four),
				string(constants.Addition),
				string(constants.Two),
				string(constants.Equal),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueLast9(),
			},
			expectedValue: [][]string{
				{
					string(constants.Four),
					string(constants.Addition),
					string(constants.Two),
					string(constants.Equal),
					string(constants.Six),
					string(constants.Equal),
					string(constants.One),
					string(constants.Addition),
					string(constants.Five),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected3(),
			},
		},
		{
			name:        "(Straight) Right Extend Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{11, 8}, {12, 8}, {13, 8}, {14, 8}},
			values: []string{
				string(constants.Equal),
				string(constants.Nine),
				string(constants.Division),
				string(constants.Three),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueFirst9(),
			},
			expectedValue: [][]string{
				{
					string(constants.Three),
					string(constants.Multiply),
					string(constants.One),
					string(constants.Equal),
					string(constants.Three),
					string(constants.Equal),
					string(constants.Nine),
					string(constants.Division),
					string(constants.Three),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected4(),
			},
		},
		{
			name:        "(Straight) Bottom Extend Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{8, 11}, {8, 12}, {8, 13}, {8, 14}},
			values: []string{
				string(constants.Equal),
				string(constants.Twelve),
				string(constants.Division),
				string(constants.Two),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardTrueFirst9(),
			},
			expectedValue: [][]string{
				{
					string(constants.Six),
					string(constants.Equal),
					string(constants.One),
					string(constants.Addition),
					string(constants.Five),
					string(constants.Equal),
					string(constants.Twelve),
					string(constants.Division),
					string(constants.Two),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected4(),
			},
		},
		{
			name:        "(Straight) Left Cross Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{5, 6}, {5, 7}, {5, 8}, {5, 9}, {5, 10}},
			values: []string{
				string(constants.One),
				string(constants.Two),
				string(constants.Equal),
				string(constants.One),
				string(constants.Two),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardAllFalse5(),
				mockIsPlaceOnBoardFalseFirst6(),
			},
			expectedValue: [][]string{
				{
					string(constants.One),
					string(constants.Two),
					string(constants.Equal),
					string(constants.One),
					string(constants.Two),
				},
				{
					string(constants.Equal),
					string(constants.Three),
					string(constants.Multiply),
					string(constants.One),
					string(constants.Equal),
					string(constants.Three),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected0(),
				mockSquareTypeSingleConnected5(),
			},
		},
		{
			name:        "(Straight) Top Cross Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{6, 5}, {7, 5}, {8, 5}, {9, 5}, {10, 5}},
			values: []string{
				string(constants.One),
				string(constants.Two),
				string(constants.Equal),
				string(constants.One),
				string(constants.Two),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardAllFalse5(),
				mockIsPlaceOnBoardFalseFirst6(),
			},
			expectedValue: [][]string{
				{
					string(constants.One),
					string(constants.Two),
					string(constants.Equal),
					string(constants.One),
					string(constants.Two),
				},
				{
					string(constants.Equal),
					string(constants.Six),
					string(constants.Equal),
					string(constants.One),
					string(constants.Addition),
					string(constants.Five),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected0(),
				mockSquareTypeSingleConnected5(),
			},
		},
		{
			name:        "(Straight) Right Cross Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{11, 6}, {11, 7}, {11, 8}, {11, 9}, {11, 10}},
			values: []string{
				string(constants.One),
				string(constants.Two),
				string(constants.Equal),
				string(constants.One),
				string(constants.Two),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardAllFalse5(),
				mockIsPlaceOnBoardFalseLast6(),
			},
			expectedValue: [][]string{
				{
					string(constants.One),
					string(constants.Two),
					string(constants.Equal),
					string(constants.One),
					string(constants.Two),
				},
				{
					string(constants.Three),
					string(constants.Multiply),
					string(constants.One),
					string(constants.Equal),
					string(constants.Three),
					string(constants.Equal),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected0(),
				mockSquareTypeSingleConnected6(),
			},
		},
		{
			name:        "(Straight) Bottom Cross Single Connected Mapping",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{6, 11}, {7, 11}, {8, 11}, {9, 11}, {10, 11}},
			values: []string{
				string(constants.One),
				string(constants.Two),
				string(constants.Equal),
				string(constants.One),
				string(constants.Two),
			},
			expectedIsPlaceOnBoard: [][]bool{
				mockIsPlaceOnBoardAllFalse5(),
				mockIsPlaceOnBoardFalseLast6(),
			},
			expectedValue: [][]string{
				{
					string(constants.One),
					string(constants.Two),
					string(constants.Equal),
					string(constants.One),
					string(constants.Two),
				},
				{
					string(constants.Six),
					string(constants.Equal),
					string(constants.One),
					string(constants.Addition),
					string(constants.Five),
					string(constants.Equal),
				},
			},
			expectedSquareType: [][]string{
				mockSquareTypeSingleConnected0(),
				mockSquareTypeSingleConnected6(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			chipForPlacing := mapChipForPlacing(tt.coordinates, tt.values)
			isVertical, _ := rules.IsChipPlaceOnVerticalOrHorizontal(tt.coordinates)
			connector := mappings.StraightConnector(tt.board, tt.coordinates, isVertical)
			results := mappings.StraightMapping(tt.board, chipForPlacing, connector, isVertical)

			testStraightMapChecking(t, results, tt.expectedIsPlaceOnBoard, tt.expectedValue, tt.expectedSquareType)
		})
	}
}

func testStraightMapChecking(
	t *testing.T,
	mapping [][]models.ChipForCalculating,
	expectedIsPlaceOnBoard [][]bool,
	expectedValue [][]string,
	expectedSquareType [][]string,
) {
	for j, resultSet := range mapping {
		for i, result := range resultSet {
			if result.IsPlacedOnBoard != expectedIsPlaceOnBoard[j][i] {
				t.Errorf("SingleChipMapping().IsPlacedOnBoard result[%d] = result:%v; expected %v", i, result.IsPlacedOnBoard, expectedIsPlaceOnBoard[j][i])
			}

			if result.ChipForCalculating.Value != expectedValue[j][i] {
				t.Errorf("SingleChipMapping().Value result[%d] = result:%v; expected %v", i, result.ChipForCalculating.Value, expectedValue[j][i])
			}

			if result.SquareType != expectedSquareType[j][i] {
				t.Errorf("SingleChipMapping().SquareType result[%d] = result:%v; expected %v", i, result.SquareType, expectedSquareType[j][i])
			}
		}
	}
}

func mapChipForPlacing(coordinate [][2]int, value []string) []models.ChipForPlacing {
	var chipForPlacing []models.ChipForPlacing

	for i, val := range value {
		chipForPlacing = append(chipForPlacing,
			models.NewChipForPlacing(coordinate[i], components.NewChip(val)))
	}

	return chipForPlacing
}

func mockStraightChipMapping() components.Board {
	board := components.NewBoard()

	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("3")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("3")))

	chips = append(chips, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("6")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("5")))

	for _, ch := range chips {
		board.Add(ch.Position, ch.Chip)
	}
	return board
}

func mockIsPlaceOnBoardTrueFirst6() []bool {
	return []bool{true, false, false, false, false, false}
}

func mockIsPlaceOnBoardTrueLast6() []bool {
	return []bool{false, false, false, false, false, true}
}

func mockIsPlaceOnBoardTrueFirst9() []bool {
	return []bool{true, true, true, true, true, false, false, false, false}
}

func mockIsPlaceOnBoardTrueLast9() []bool {
	return []bool{false, false, false, false, true, true, true, true, true}
}

func mockIsPlaceOnBoardAllFalse5() []bool {
	return []bool{false, false, false, false, false}
}

func mockSquareTypeSingleConnected0() []string {
	return []string{
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
	}
}

func mockSquareTypeSingleConnected1() []string {
	return []string{
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.BlueSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
	}
}

func mockSquareTypeSingleConnected2() []string {
	return []string{
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.BlueSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
	}
}

func mockSquareTypeSingleConnected3() []string {
	return []string{
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.OrangeSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.CenterSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
	}
}

func mockSquareTypeSingleConnected4() []string {
	return []string{
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.CenterSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.OrangeSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
	}
}

func mockSquareTypeSingleConnected5() []string {
	return []string{
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.CenterSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
	}
}

func mockSquareTypeSingleConnected6() []string {
	return []string{
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.CenterSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
		string(constants.NormalSquare),
	}
}
