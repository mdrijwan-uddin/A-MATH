package mappings_test

import (
	"A-MATH/game/components"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"A-MATH/game/rules"
	"testing"
)

func TestSingleChipConnector(t *testing.T) {
	tests := []struct {
		name              string
		board             components.Board
		expectedConnector models.ChipConnector
	}{
		{
			name:  "(1D) Right Connected",
			board: mockSingleDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{1, 8},
				LeftConnected:   false,
				TopConnected:    false,
				RightConnected:  true,
				BottomConnected: false,
			},
		},
		{
			name:  "(1D) Bottom Connected",
			board: mockSingleDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{8, 1},
				LeftConnected:   false,
				TopConnected:    false,
				RightConnected:  false,
				BottomConnected: true,
			},
		},
		{
			name:  "(1D) Left Connected",
			board: mockSingleDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{15, 8},
				LeftConnected:   true,
				TopConnected:    false,
				RightConnected:  false,
				BottomConnected: false,
			},
		},
		{
			name:  "(1D) Top Connected",
			board: mockSingleDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{8, 15},
				LeftConnected:   false,
				TopConnected:    true,
				RightConnected:  false,
				BottomConnected: false,
			},
		},
		{
			name:  "(2D) Bottom-Right Connected",
			board: mockTwoDirectionChipMapping1(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{5, 5},
				LeftConnected:   false,
				TopConnected:    false,
				RightConnected:  true,
				BottomConnected: true,
			},
		},
		{
			name:  "(2D) Top-Right Connected",
			board: mockTwoDirectionChipMapping1(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{5, 11},
				LeftConnected:   false,
				TopConnected:    true,
				RightConnected:  true,
				BottomConnected: false,
			},
		},
		{
			name:  "(2D) Bottom-Left Connected",
			board: mockTwoDirectionChipMapping1(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{11, 5},
				LeftConnected:   true,
				TopConnected:    false,
				RightConnected:  false,
				BottomConnected: true,
			},
		},
		{
			name:  "(2D) Top-Left Connected",
			board: mockTwoDirectionChipMapping1(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{11, 11},
				LeftConnected:   true,
				TopConnected:    true,
				RightConnected:  false,
				BottomConnected: false,
			},
		},
		{
			name:  "(2D) Top-Bottom Vertical",
			board: mockTwoDirectionChipMapping2(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{4, 9},
				LeftConnected:   false,
				TopConnected:    true,
				RightConnected:  false,
				BottomConnected: true,
			},
		},
		{
			name:  "(2D) Left-Right Horizontal",
			board: mockTwoDirectionChipMapping2(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{5, 10},
				LeftConnected:   true,
				TopConnected:    false,
				RightConnected:  true,
				BottomConnected: false,
			},
		},
		{
			name:  "(3D) Top-Bottom Vertical & Right Connected Mapping",
			board: mockThreeDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{4, 8},
				LeftConnected:   false,
				TopConnected:    true,
				RightConnected:  true,
				BottomConnected: true,
			},
		},
		{
			name:  "(3D) Left-Right Horizontal & Bottom Connected Mapping",
			board: mockThreeDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{8, 4},
				LeftConnected:   true,
				TopConnected:    false,
				RightConnected:  true,
				BottomConnected: true,
			},
		},
		{
			name:  "(3D) Top-Bottom Vertical & Left Connected Mapping",
			board: mockThreeDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{12, 8},
				LeftConnected:   true,
				TopConnected:    true,
				RightConnected:  false,
				BottomConnected: true,
			},
		},
		{
			name:  "(3D) Left-Right Horizontal & Top Connected Mapping",
			board: mockThreeDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{8, 12},
				LeftConnected:   true,
				TopConnected:    true,
				RightConnected:  true,
				BottomConnected: false,
			},
		},
		{
			name:  "(4D) Left-Right Horizontal & Top-Bottom Vertical Connected Mapping",
			board: mockFourDirectionChipMapping(),
			expectedConnector: models.ChipConnector{
				ChipCoordinate:  [2]int{5, 8},
				LeftConnected:   true,
				TopConnected:    true,
				RightConnected:  true,
				BottomConnected: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connector := mappings.SingleChipConnector(tt.board, tt.expectedConnector.ChipCoordinate)
			if len(connector) != 1 {
				t.Errorf("Single chip connector should have lenght = 1")
			}

			if connector[0] != tt.expectedConnector {
				t.Errorf("SingleChipConnector with[%d, %d] got\n result:%v;\n expected %v",
					tt.expectedConnector.ChipCoordinate[0],
					tt.expectedConnector.ChipCoordinate[1],
					connector[0],
					tt.expectedConnector,
				)
			}
		})
	}
}

func TestStraightConnector(t *testing.T) {
	tests := []struct {
		name               string
		board              components.Board
		coordinates        [][2]int
		expectedConnectors []models.ChipConnector
	}{
		{
			name:        "(Straight) Left Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{9, 6}, {10, 6}, {11, 6}, {12, 6}, {13, 6}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{9, 6},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{12, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{13, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Top Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{10, 9}, {10, 10}, {10, 11}, {10, 12}, {10, 13}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{10, 9},
					LeftConnected:   false,
					TopConnected:    true,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 11},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 12},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 13},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Right Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{3, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{4, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{5, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{6, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{7, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Bottom Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{6, 3}, {6, 4}, {6, 5}, {6, 6}, {6, 7}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{6, 3},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{6, 4},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{6, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{6, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{6, 7},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: true,
				},
			},
		},
		{
			name:        "(Straight) Left Extend Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{2, 8}, {3, 8}, {4, 8}, {5, 8}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{2, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{3, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{4, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{5, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Top Extend Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{8, 2}, {8, 3}, {8, 4}, {8, 5}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{8, 2},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 3},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 4},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: true,
				},
			},
		},
		{
			name:        "(Straight) Right Extend Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{11, 8}, {12, 8}, {13, 8}, {14, 8}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{11, 8},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{12, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{13, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{14, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Bottom Extend Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{8, 11}, {8, 12}, {8, 13}, {8, 14}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{8, 11},
					LeftConnected:   false,
					TopConnected:    true,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 12},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 13},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 14},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Left Cross Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{5, 6}, {5, 7}, {5, 8}, {5, 9}, {5, 10}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{5, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{5, 7},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{5, 8},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{5, 9},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{5, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Top Cross Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{6, 5}, {7, 5}, {8, 5}, {9, 5}, {10, 5}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{6, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{7, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: true,
				},
				{
					ChipCoordinate:  [2]int{9, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Right Cross Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{11, 6}, {11, 7}, {11, 8}, {11, 9}, {11, 10}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{11, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 7},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 8},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 9},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Straight) Bottom Cross Single Connected",
			board:       mockStraightChipMapping(),
			coordinates: [][2]int{{6, 11}, {7, 11}, {8, 11}, {9, 11}, {10, 11}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{6, 11},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{7, 11},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{8, 11},
					LeftConnected:   false,
					TopConnected:    true,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{9, 11},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 11},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isVertical, _ := rules.IsChipPlaceOnVerticalOrHorizontal(tt.coordinates)
			connectors := mappings.StraightConnector(tt.board, tt.coordinates, isVertical)

			if len(connectors) != len(tt.expectedConnectors) {
				t.Errorf("Straight connector should have lenght equal to: %d", len(connectors))
			}

			for i := 0; i < len(connectors); i++ {
				if connectors[i] != tt.expectedConnectors[i] {
					t.Errorf("Straight Connector[%d] with[%d, %d] got\n result:%v;\n expected %v",
						i,
						tt.expectedConnectors[i].ChipCoordinate[0],
						tt.expectedConnectors[i].ChipCoordinate[1],
						connectors[i],
						tt.expectedConnectors[i],
					)
				}
			}
		})
	}
}

func TestCrossConnector(t *testing.T) {
	tests := []struct {
		name               string
		board              components.Board
		coordinates        [][2]int
		expectedConnectors []models.ChipConnector
	}{
		{
			name:        "(Cross) Vertical Connected Mapping",
			board:       mockCrossChipMapping1(),
			coordinates: [][2]int{{11, 5}, {11, 6}, {11, 7}, {11, 9}, {11, 10}, {11, 11}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{11, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 6},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 7},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: true,
				},
				{
					ChipCoordinate:  [2]int{11, 9},
					LeftConnected:   false,
					TopConnected:    true,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 11},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Cross) Horizontal Connected Mapping",
			board:       mockCrossChipMapping1(),
			coordinates: [][2]int{{5, 5}, {6, 5}, {7, 5}, {9, 5}, {10, 5}, {11, 5}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{5, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{6, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{7, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{9, 5},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{10, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Cross) Extends Vertical Connected Mapping",
			board:       mockCrossChipMapping1(),
			coordinates: [][2]int{{13, 3}, {13, 4}, {13, 12}, {13, 13}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{13, 3},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{13, 4},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: true,
				},
				{
					ChipCoordinate:  [2]int{13, 12},
					LeftConnected:   false,
					TopConnected:    true,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{13, 13},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Cross) Extends Horizontal Connected Mapping",
			board:       mockCrossChipMapping1(),
			coordinates: [][2]int{{3, 3}, {4, 3}, {12, 3}, {13, 3}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{3, 3},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{4, 3},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{12, 3},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{13, 3},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
		{
			name:        "(Cross) Multiple Vertical Cross Mapping",
			board:       mockCrossChipMapping2(),
			coordinates: [][2]int{{11, 7}, {11, 9}, {11, 10}, {11, 11}, {11, 13}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{11, 7},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: true,
				},
				{
					ChipCoordinate:  [2]int{11, 9},
					LeftConnected:   false,
					TopConnected:    true,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 10},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{11, 11},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: true,
				},
				{
					ChipCoordinate:  [2]int{11, 13},
					LeftConnected:   false,
					TopConnected:    true,
					RightConnected:  false,
					BottomConnected: true,
				},
			},
		},
		{
			name:        "(Cross) Multiple Horizontal Cross Mapping",
			board:       mockCrossChipMapping2(),
			coordinates: [][2]int{{1, 5}, {3, 5}, {5, 5}, {7, 5}},
			expectedConnectors: []models.ChipConnector{
				{
					ChipCoordinate:  [2]int{1, 5},
					LeftConnected:   false,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{3, 5},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{5, 5},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  true,
					BottomConnected: false,
				},
				{
					ChipCoordinate:  [2]int{7, 5},
					LeftConnected:   true,
					TopConnected:    false,
					RightConnected:  false,
					BottomConnected: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isVertical, _ := rules.IsChipPlaceOnVerticalOrHorizontal(tt.coordinates)
			connectors := mappings.CrossConnector(tt.board, tt.coordinates, isVertical)

			if len(connectors) != len(tt.expectedConnectors) {
				t.Errorf("Cross connector should have lenght equal to: %d", len(connectors))
			}

			for i := 0; i < len(connectors); i++ {
				if connectors[i] != tt.expectedConnectors[i] {
					t.Errorf("Cross Connector[%d] with[%d, %d] got\n result:%v;\n expected %v",
						i,
						tt.expectedConnectors[i].ChipCoordinate[0],
						tt.expectedConnectors[i].ChipCoordinate[1],
						connectors[i],
						tt.expectedConnectors[i],
					)
				}
			}
		})
	}
}
