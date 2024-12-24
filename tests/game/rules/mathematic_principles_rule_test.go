package rules_test

import (
	"A-MATH/game/components"
	"A-MATH/game/rules"
	"testing"
)

func TestHaveEqualsInEquation(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct equation (have equal)[1]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("*"),
				components.NewChip("4"),
			},
			expected: true,
		},
		{
			name: "Correct equation (have equal)[2]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("="),
				components.NewChip("2"),
				components.NewChip("="),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Incorrect equation (doesn't have equal)",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("4"),
				components.NewChip("+"),
				components.NewChip("6"),
				components.NewChip("-"),
				components.NewChip("4"),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.HaveEqualsInEquation(tt.chipSet)

			if result != tt.expected {
				t.Errorf("HaveEqualsInEquation() = result:%v; expected %v", result, tt.expected)
			}

		})
	}
}

func TestIsChipsLenghtReasonably(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct chips lenght (3-15 chips)[1]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("="),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Correct chips lenght (3-15 chips)[2]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("*"),
				components.NewChip("4"),
			},
			expected: true,
		},
		{
			name: "Correct chips lenght (3-15 chips)[3]",
			chipSet: []components.Chip{
				components.NewChip("1"),
				components.NewChip("5"),
				components.NewChip("="),
				components.NewChip("15"),
				components.NewChip("+"),
				components.NewChip("8"),
				components.NewChip("9"),
				components.NewChip("7"),
				components.NewChip("*"),
				components.NewChip("0"),
				components.NewChip("/"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("1"),
				components.NewChip("5"),
			},
			expected: true,
		},
		{
			name: "Incorrect chips lenght (not between 3-15 chips)[1]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("="),
			},
			expected: false,
		},
		{
			name: "Incorrect chips lenght (not between 3-15 chips)[2]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("*"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("7"),
				components.NewChip("2"),
				components.NewChip("/"),
				components.NewChip("3"),
				components.NewChip("="),
				components.NewChip("1"),
				components.NewChip("6"),
				components.NewChip("+"),
				components.NewChip("8"),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.IsChipsLenghtReasonably(tt.chipSet)

			if result != tt.expected {
				t.Errorf("IsChipsLenghtReasonably() length: %d = result:%v; expected %v", len(tt.chipSet), result, tt.expected)
			}

		})
	}
}

func TestIsOperationPlaceOnUnusualLocation(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct Operation Location[1]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("*"),
				components.NewChip("4"),
			},
			expected: false,
		},
		{
			name: "Correct Operation Location[2]",
			chipSet: []components.Chip{
				components.NewChip("-"),
				components.NewChip("2"),
				components.NewChip("+"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("2"),
			},
			expected: false,
		},
		{
			name: "Incorrect Operation Location[1]",
			chipSet: []components.Chip{
				components.NewChip("+"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("-"),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Incorrect Operation Location[2]",
			chipSet: []components.Chip{
				components.NewChip("*"),
				components.NewChip("3"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("/"),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Incorrect Operation Location[3]",
			chipSet: []components.Chip{
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("-"),
				components.NewChip("2"),
				components.NewChip("-"),
			},
			expected: true,
		},
		{
			name: "Incorrect Operation Location[4]",
			chipSet: []components.Chip{
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("-"),
				components.NewChip("2"),
				components.NewChip("/"),
			},
			expected: true,
		},
		{
			name: "Incorrect Operation Location[5]",
			chipSet: []components.Chip{
				components.NewChip("5"),
				components.NewChip("/"),
				components.NewChip("7"),
				components.NewChip("1"),
				components.NewChip("="),
				components.NewChip("-"),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.IsOperationPlaceOnUnusualLocation(tt.chipSet)

			if result != tt.expected {
				t.Errorf("IsOperationPlaceOnUnusualLocation() = result:%v; expected %v", result, tt.expected)
			}

		})
	}
}

func TestAreOperationNextToEachOther(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct equation (operation aren't next to each other)[1]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("*"),
				components.NewChip("4"),
			},
			expected: false,
		},
		{
			name: "Correct equation (operation aren't next to each other)[2]",
			chipSet: []components.Chip{
				components.NewChip("8"),
				components.NewChip("="),
				components.NewChip("-"),
				components.NewChip("6"),
				components.NewChip("+"),
				components.NewChip("14"),
			},
			expected: false,
		},
		{
			name: "Correct equation (operation aren't next to each other)[3]",
			chipSet: []components.Chip{
				components.NewChip("3"),
				components.NewChip("-"),
				components.NewChip("9"),
				components.NewChip("="),
				components.NewChip("-"),
				components.NewChip("6"),
			},
			expected: false,
		},
		{
			name: "Incorrect equation (operation aren't next to each other)[1]",
			chipSet: []components.Chip{
				components.NewChip("-"),
				components.NewChip("2"),
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("6"),
				components.NewChip("+"),
				components.NewChip("-"),
				components.NewChip("4"),
			},
			expected: true,
		},
		{
			name: "Incorrect equation (operation aren't next to each other)[2]",
			chipSet: []components.Chip{
				components.NewChip("15"),
				components.NewChip("="),
				components.NewChip("*"),
				components.NewChip("6"),
				components.NewChip("+"),
				components.NewChip("9"),
			},
			expected: true,
		},
		{
			name: "Incorrect equation (operation aren't next to each other)[3]",
			chipSet: []components.Chip{
				components.NewChip("3"),
				components.NewChip("+"),
				components.NewChip("="),
				components.NewChip("1"),
				components.NewChip("+"),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Incorrect equation (operation aren't next to each other)[4]",
			chipSet: []components.Chip{
				components.NewChip("4"),
				components.NewChip("-"),
				components.NewChip("-"),
				components.NewChip("3"),
				components.NewChip("="),
				components.NewChip("7"),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.AreOperationNextToEachOther(tt.chipSet)

			if result != tt.expected {
				t.Errorf("AreOperationNextToEachOther() = result:%v; expected %v", result, tt.expected)
			}

		})
	}
}

func TestAreTwoDigitNumbersStackedWithTheOtherNumber(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct 2-digit numbers stacking equation [1]",
			chipSet: []components.Chip{
				components.NewChip("12"),
				components.NewChip("="),
				components.NewChip("1"),
				components.NewChip("9"),
				components.NewChip("-"),
				components.NewChip("7"),
			},
			expected: false,
		},
		{
			name: "Incorrect 2-digit numbers stacking equation [1]",
			chipSet: []components.Chip{
				components.NewChip("2"),
				components.NewChip("12"),
				components.NewChip("="),
				components.NewChip("2"),
				components.NewChip("1"),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Incorrect 2-digit numbers stacking equation [2]",
			chipSet: []components.Chip{
				components.NewChip("4"),
				components.NewChip("5"),
				components.NewChip("="),
				components.NewChip("12"),
				components.NewChip("1"),
				components.NewChip("4"),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.AreTwoDigitNumbersStackedWithTheOtherNumber(tt.chipSet)

			if result != tt.expected {
				t.Errorf("AreTwoDigitNumbersStackedWithTheOtherNumber() = result:%v; expected %v", result, tt.expected)
			}

		})
	}
}

func TestAreOneDigitNumbersReasonablyStacked(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct 1-digit numbers stacking reasonably [1]",
			chipSet: []components.Chip{
				components.NewChip("12"),
				components.NewChip("="),
				components.NewChip("1"),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Correct 1-digit numbers stacking reasonably [2]",
			chipSet: []components.Chip{
				components.NewChip("12"),
				components.NewChip("*"),
				components.NewChip("11"),
				components.NewChip("="),
				components.NewChip("1"),
				components.NewChip("3"),
				components.NewChip("2"),
			},
			expected: true,
		},
		{
			name: "Incorrect 1-digit numbers stacking reasonably [1]",
			chipSet: []components.Chip{
				components.NewChip("5"),
				components.NewChip("3"),
				components.NewChip("2"),
				components.NewChip("*"),
				components.NewChip("2"),
				components.NewChip("="),
				components.NewChip("1"),
				components.NewChip("0"),
				components.NewChip("6"),
				components.NewChip("4"),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.AreOneDigitNumbersReasonablyStacked(tt.chipSet)

			if result != tt.expected {
				t.Errorf("AreOneDigitNumbersReasonablyStacked() = result:%v; expected %v", result, tt.expected)
			}

		})
	}
}

func TestIsZeroAtTheBeginWithNumberFormed(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct no 0 at the beginning of formed number [1]",
			chipSet: []components.Chip{
				components.NewChip("20"),
				components.NewChip("="),
				components.NewChip("2"),
				components.NewChip("0"),
			},
			expected: false,
		},
		// {
		// 	name: "Incorrect 0 at the beginning of formed number [1]",
		// 	chipSet: []components.Chip{
		// 		components.NewChip("0"),
		// 		components.NewChip("5"),
		// 		components.NewChip("+"),
		// 		components.NewChip("8"),
		// 		components.NewChip("="),
		// 		components.NewChip("1"),
		// 		components.NewChip("3"),
		// 	},
		// 	expected: true,
		// },
		// {
		// 	name: "Incorrect 0 at the beginning of formed number [2]",
		// 	chipSet: []components.Chip{
		// 		components.NewChip("9"),
		// 		components.NewChip("*"),
		// 		components.NewChip("4"),
		// 		components.NewChip("="),
		// 		components.NewChip("0"),
		// 		components.NewChip("3"),
		// 		components.NewChip("6"),
		// 	},
		// 	expected: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.IsZeroAtTheBeginWithNumberFormed(tt.chipSet)

			if result != tt.expected {
				t.Errorf("IsZeroAtTheBeginWithNumberFormed() = result:%v; expected %v", result, tt.expected)
			}

		})
	}
}

func TestHasDivideByZero(t *testing.T) {
	tests := []struct {
		name     string
		chipSet  []components.Chip
		expected bool
	}{
		{
			name: "Correct equation (no divide by 0) [1]",
			chipSet: []components.Chip{
				components.NewChip("4"),
				components.NewChip("="),
				components.NewChip("4"),
				components.NewChip("+"),
				components.NewChip("0"),
				components.NewChip("/"),
				components.NewChip("9"),
				components.NewChip("2"),
			},
			expected: false,
		},
		{
			name: "Correct equation (no divide by 0) [1]",
			chipSet: []components.Chip{
				components.NewChip("0"),
				components.NewChip("*"),
				components.NewChip("4"),
				components.NewChip("/"),
				components.NewChip("4"),
				components.NewChip("7"),
				components.NewChip("="),
				components.NewChip("0"),
			},
			expected: false,
		},
		{
			name: "Incorrect divide by 0 [1]",
			chipSet: []components.Chip{
				components.NewChip("12"),
				components.NewChip("/"),
				components.NewChip("0"),
				components.NewChip("+"),
				components.NewChip("1"),
				components.NewChip("="),
				components.NewChip("1"),
			},
			expected: true,
		},
		{
			name: "Incorrect divide by 0 [2]",
			chipSet: []components.Chip{
				components.NewChip("8"),
				components.NewChip("="),
				components.NewChip("14"),
				components.NewChip("/"),
				components.NewChip("0"),
				components.NewChip("+"),
				components.NewChip("8"),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.HasDivideByZero(tt.chipSet)

			if result != tt.expected {
				t.Errorf("HasDivideByZero() = result:%v; expected %v", result, tt.expected)
			}

		})
	}
}
