package components_test

import (
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
