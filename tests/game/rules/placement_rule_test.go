package rules_test

import (
	"A-MATH/game/components"
	"A-MATH/game/models"
	"A-MATH/game/rules"
	"testing"
)

func TestIsChipPlaceOnCenterSquare(t *testing.T) {
	tests := []struct {
		name        string
		coordinates [][2]int
		expected    bool
	}{
		{
			name: "Correct vertical chip placement",
			coordinates: [][2]int{
				{8, 8},
				{8, 9},
				{8, 10},
				{8, 11},
				{8, 12},
				{8, 13},
				{8, 14},
				{8, 15},
			},
			expected: true,
		},
		{
			name: "Correct horizontal chip placement",
			coordinates: [][2]int{
				{1, 8}, {2, 8}, {3, 8}, {4, 8}, {5, 8}, {6, 8}, {7, 8}, {8, 8},
			},
			expected: true,
		},
		{
			name: "Incorrect vertical chip placement",
			coordinates: [][2]int{
				{9, 8},
				{9, 9},
				{9, 10},
				{9, 11},
				{9, 12},
				{9, 13},
				{9, 14},
				{9, 15},
			},
			expected: false,
		},
		{
			name: "Incorrect horizontal chip placement",
			coordinates: [][2]int{
				{1, 9}, {2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9}, {8, 9},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.IsChipPlaceOnCenterSquare(tt.coordinates)

			if result != tt.expected {
				t.Errorf("IsChipPlaceOnCenterSquare(%v) = result:%v; expected %v", tt.coordinates, result, tt.expected)
			}

		})
	}
}

func TestIsChipPlaceOnVerticalOrHorizontal(t *testing.T) {
	tests := []struct {
		name                  string
		coordinates           [][2]int
		expectedForVertical   bool
		expectedForHorizontal bool
	}{
		{
			name: "Correct regular vertical chip placement",
			coordinates: [][2]int{
				{3, 6},
				{3, 7},
				{3, 8},
				{3, 9},
				{3, 10},
				{3, 11},
				{3, 12},
				{3, 13},
				{3, 14},
				{3, 15},
			},
			expectedForVertical:   true,
			expectedForHorizontal: false,
		},
		{
			name: "Correct regular horizontal chip placement",
			coordinates: [][2]int{
				{11, 5}, {12, 5}, {13, 5}, {14, 5},
			},
			expectedForVertical:   false,
			expectedForHorizontal: true,
		},
		{
			name: "Correct seperated vertical chip placement",
			coordinates: [][2]int{
				{13, 4},
				{13, 5},
				{13, 12},
				{13, 13},
			},
			expectedForVertical:   true,
			expectedForHorizontal: false,
		},
		{
			name: "Correct seperated horizontal chip placement",
			coordinates: [][2]int{
				{8, 2}, {9, 2}, {10, 2}, {14, 2}, {15, 2},
			},
			expectedForVertical:   false,
			expectedForHorizontal: true,
		},
		{
			name: "Correct single chip placement",
			coordinates: [][2]int{
				{3, 5},
			},
			expectedForVertical:   true,
			expectedForHorizontal: true,
		},
		{
			name: "Incorrect chip placement (case 1)",
			coordinates: [][2]int{
				{2, 7}, {5, 7}, {2, 12}, {5, 12},
			},
			expectedForVertical:   false,
			expectedForHorizontal: false,
		},
		{
			name: "Incorrect chip placement (case 2)",
			coordinates: [][2]int{
				{9, 5},
				{9, 6},
				{9, 7},
				{10, 8},
			},
			expectedForVertical:   false,
			expectedForHorizontal: false,
		},
		{
			name: "Incorrect chip placement (case 3)",
			coordinates: [][2]int{
				{4, 4}, {5, 5}, {6, 6}, {7, 7},
			},
			expectedForVertical:   false,
			expectedForHorizontal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultForVertical, resultForHorizontal := rules.IsChipPlaceOnVerticalOrHorizontal(tt.coordinates)

			if resultForVertical != tt.expectedForVertical {
				t.Errorf("IsChipPlaceOnVertical(%v) = result:%v; expected %v", tt.coordinates, resultForVertical, tt.expectedForVertical)
			}

			if resultForHorizontal != tt.expectedForHorizontal {
				t.Errorf("IsChipPlaceOnHorizontal(%v) = result:%v; expected %v", tt.coordinates, resultForHorizontal, tt.expectedForHorizontal)
			}

		})
	}
}

func TestIsChipsPlacedOnOccupiedSquare(t *testing.T) {
	var boardForTest components.Board
	var chipForPlacing []models.ChipForPlacing
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("14")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("-")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("7")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 8}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{12, 8}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 8}, components.NewChip("8")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{14, 8}, components.NewChip("+-")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{15, 8}, components.NewChip("11")))

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 11}, components.NewChip("+-")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 12}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 13}, components.NewChip("0")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 14}, components.NewChip("+")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 15}, components.NewChip("3")))

	for _, ch := range chipForPlacing {
		boardForTest.Add(ch.Position, ch.Chip)
	}

	tests := []struct {
		name        string
		board       components.Board
		coordinates [][2]int
		expected    bool
	}{
		{
			name:  "Correct vertical chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{12, 2},
				{12, 3},
				{12, 4},
				{12, 5},
				{12, 6},
				{12, 7},
			},
			expected: false,
		},
		{
			name:  "Correct horizontal chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{9, 15}, {10, 15}, {11, 14}, {12, 15}, {13, 15},
			},
			expected: false,
		},
		{
			name:  "Incorrect vertical chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{15, 5},
				{15, 6},
				{15, 7},
				{15, 8},
			},
			expected: true,
		},
		{
			name:  "Incorrect horizontal chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{7, 11}, {8, 11}, {9, 11}, {10, 11}, {11, 11},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := rules.IsChipsPlacedOnOccupiedSquare(tt.board, tt.coordinates)

			if result != tt.expected {
				t.Errorf("IsChipsPlacedOnOccupiedSquare(%v) = result:%v; expected %v", tt.coordinates, result, tt.expected)
			}

		})
	}
}

func TestIsChipPlaceingOnStraightLineOrSeperated(t *testing.T) {
	var boardForTest components.Board
	var chipForPlacing []models.ChipForPlacing
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("14")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("-")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("7")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 8}, components.NewChip("+-")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{12, 8}, components.NewChip("11")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 8}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{14, 8}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{15, 8}, components.NewChip("~")))

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 2}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 3}, components.NewChip("4")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 4}, components.NewChip("+")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 5}, components.NewChip("9")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 6}, components.NewChip("*")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 7}, components.NewChip("3")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 9}, components.NewChip("8")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 10}, components.NewChip("1")))

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 1}, components.NewChip("9")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 2}, components.NewChip("/")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 3}, components.NewChip("3")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 4}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 5}, components.NewChip("*/")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("0")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("-")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("*/")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("0")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 11}, components.NewChip("/")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 12}, components.NewChip("17")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 13}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 14}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 15}, components.NewChip("4")))

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{4, 4}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{5, 4}, components.NewChip("*")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{6, 4}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{7, 4}, components.NewChip("=")))

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{9, 14}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{10, 14}, components.NewChip("0")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 14}, components.NewChip("/")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{12, 14}, components.NewChip("10")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 14}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{14, 14}, components.NewChip("15")))

	for _, ch := range chipForPlacing {
		boardForTest.Add(ch.Position, ch.Chip)
	}

	tests := []struct {
		name                  string
		board                 components.Board
		coordinates           [][2]int
		expectedToBeStraight  bool
		expectedToBeSeperated bool
	}{
		{
			name:  "Correct straight vertical chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{10, 8},
				{10, 9},
				{10, 10},
				{10, 11},
				{10, 12},
				{10, 13},
			},
			expectedToBeStraight:  true,
			expectedToBeSeperated: false,
		},
		{
			name:  "Correct straight horizontal chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{9, 10}, {10, 10}, {11, 10}, {12, 10}, {13, 10},
			},
			expectedToBeStraight:  true,
			expectedToBeSeperated: false,
		},
		{
			name:  "Correct cross vertical chip placement (case 1)",
			board: boardForTest,
			coordinates: [][2]int{
				{15, 6},
				{15, 7},
				{15, 9},
				{15, 10},
			},
			expectedToBeStraight:  false,
			expectedToBeSeperated: true,
		},
		{
			name:  "Correct cross vertical chip placement (case 2)",
			board: boardForTest,
			coordinates: [][2]int{
				{13, 1},
				{13, 11},
				{13, 12},
				{13, 13},
			},
			expectedToBeStraight:  false,
			expectedToBeSeperated: true,
		},
		{
			name:  "Correct cross horizontal chip placement (case 1)",
			board: boardForTest,
			coordinates: [][2]int{
				{5, 12}, {6, 12}, {7, 12}, {9, 12}, {10, 12},
			},
			expectedToBeStraight:  false,
			expectedToBeSeperated: true,
		},
		{
			name:  "Correct cross horizontal chip placement (case 2)",
			board: boardForTest,
			coordinates: [][2]int{
				{2, 4}, {3, 4}, {9, 4}, {10, 4},
			},
			expectedToBeStraight:  false,
			expectedToBeSeperated: true,
		},
		{
			name:  "Incorrect cross vertical chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{13, 1},
				{13, 13},
				{13, 15},
			},
			expectedToBeStraight:  false,
			expectedToBeSeperated: false,
		},
		{
			name:  "Incorrect cross horizontal chip placement",
			board: boardForTest,
			coordinates: [][2]int{
				{2, 4}, {3, 4}, {10, 4}, {11, 4}, {12, 4}, {14, 4},
			},
			expectedToBeStraight:  false,
			expectedToBeSeperated: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isVertical, _ := rules.IsChipPlaceOnVerticalOrHorizontal(tt.coordinates)
			resultForStraightLine, resultForSeperated := rules.IsChipPlacingOnStraightLineOrSeparated(tt.board, tt.coordinates, isVertical)

			if resultForStraightLine != tt.expectedToBeStraight {
				t.Errorf("IsChipPlaceingOnStraightLine(%v) = result:%v; expected %v", tt.coordinates, resultForStraightLine, tt.expectedToBeStraight)
			}

			if resultForSeperated != tt.expectedToBeSeperated {
				t.Errorf("IsChipPlaceingOnSeperated(%v) = result:%v; expected %v", tt.coordinates, resultForSeperated, tt.expectedToBeSeperated)
			}

		})
	}
}
