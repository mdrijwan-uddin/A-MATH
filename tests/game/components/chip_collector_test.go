package components_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"testing"
)

func TestNewChipCollector(t *testing.T) {
	tests := []struct {
		name            string
		chips           components.Chip
		total           int
		expectedChips   components.Chip
		expectedTotal   int
		expectedMaxChip int
	}{
		{
			name:            "Chips: One (Collector)",
			chips:           components.NewChip(string(constants.One)),
			total:           6,
			expectedChips:   components.NewChip(string(constants.One)),
			expectedTotal:   6,
			expectedMaxChip: 6,
		},
		{
			name:            "Chips: Fifteen (Collector)",
			chips:           components.NewChip(string(constants.Fifteen)),
			total:           1,
			expectedChips:   components.NewChip(string(constants.Fifteen)),
			expectedTotal:   1,
			expectedMaxChip: 1,
		},
		{
			name:            "Chip: Equal (Collector)",
			chips:           components.NewChip(string(constants.Equal)),
			total:           11,
			expectedChips:   components.NewChip(string(constants.Equal)),
			expectedTotal:   11,
			expectedMaxChip: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := components.NewChipCollector(tt.chips, tt.total)

			if result.Chips != tt.expectedChips {
				t.Errorf("NewChipCollector(%s, %d) = Chip:%s; expected %s", tt.chips, tt.total, result.Chips, tt.expectedChips)
			}

			if result.Total != tt.expectedTotal {
				t.Errorf("NewChipCollector(%s, %d) = Total:%d; expected %d", tt.chips, tt.total, result.Total, tt.expectedTotal)
			}

			if result.MaxChip != tt.expectedMaxChip {
				t.Errorf("NewChipCollector(%s, %d) = Max Chip:%d; expected %d", tt.chips, tt.total, result.MaxChip, tt.expectedMaxChip)
			}
		})
	}
}

func TestDecreaseChip(t *testing.T) {
	tests := []struct {
		name            string
		chips           components.Chip
		total           int
		expectedChips   components.Chip
		expectedTotal   int
		expectedMaxChip int
	}{
		{
			name:            "Chips: One (Decreasing)",
			chips:           components.NewChip(string(constants.One)),
			total:           6,
			expectedChips:   components.NewChip(string(constants.One)),
			expectedTotal:   4,
			expectedMaxChip: 6,
		},
		{
			name:            "Chips: Fifteen (Decreasing)",
			chips:           components.NewChip(string(constants.Fifteen)),
			total:           1,
			expectedChips:   components.NewChip(string(constants.Fifteen)),
			expectedTotal:   0,
			expectedMaxChip: 1,
		},
		{
			name:            "Chip: Equal (Decreasing)",
			chips:           components.NewChip(string(constants.Equal)),
			total:           11,
			expectedChips:   components.NewChip(string(constants.Equal)),
			expectedTotal:   9,
			expectedMaxChip: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := components.NewChipCollector(tt.chips, tt.total)
			result.DecreaseChip()
			result.DecreaseChip()

			if result.Chips != tt.expectedChips {
				t.Errorf("NewChipCollector(%s, %d).DecreaseChip().DecreaseChip() = Chip:%s; expected %s", tt.chips, tt.total, result.Chips, tt.expectedChips)
			}

			if result.Total != tt.expectedTotal {
				t.Errorf("NewChipCollector(%s, %d).DecreaseChip().DecreaseChip() = Total:%d; expected %d", tt.chips, tt.total, result.Total, tt.expectedTotal)
			}

			if result.MaxChip != tt.expectedMaxChip {
				t.Errorf("NewChipCollector(%s, %d).DecreaseChip().DecreaseChip() = Max Chip:%d; expected %d", tt.chips, tt.total, result.MaxChip, tt.expectedMaxChip)
			}
		})
	}
}

func TestIncreaseChip(t *testing.T) {
	tests := []struct {
		name            string
		chips           components.Chip
		total           int
		expectedChips   components.Chip
		expectedTotal   int
		expectedMaxChip int
	}{
		{
			name:            "Chips: One (Increasing)",
			chips:           components.NewChip(string(constants.One)),
			total:           6,
			expectedChips:   components.NewChip(string(constants.One)),
			expectedTotal:   5,
			expectedMaxChip: 6,
		},
		{
			name:            "Chips: Fifteen (Increasing)",
			chips:           components.NewChip(string(constants.Fifteen)),
			total:           1,
			expectedChips:   components.NewChip(string(constants.Fifteen)),
			expectedTotal:   1,
			expectedMaxChip: 1,
		},
		{
			name:            "Chip: Equal (Increasing)",
			chips:           components.NewChip(string(constants.Equal)),
			total:           11,
			expectedChips:   components.NewChip(string(constants.Equal)),
			expectedTotal:   10,
			expectedMaxChip: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := components.NewChipCollector(tt.chips, tt.total)
			result.DecreaseChip()
			result.DecreaseChip()
			result.IncreaseChip()

			if result.Chips != tt.expectedChips {
				t.Errorf("NewChipCollector(%s, %d).DecreaseChip().DecreaseChip().IncreaseChip() = Chip:%s; expected %s", tt.chips, tt.total, result.Chips, tt.expectedChips)
			}

			if result.Total != tt.expectedTotal {
				t.Errorf("NewChipCollector(%s, %d).DecreaseChip().DecreaseChip().IncreaseChip() = Total:%d; expected %d", tt.chips, tt.total, result.Total, tt.expectedTotal)
			}

			if result.MaxChip != tt.expectedMaxChip {
				t.Errorf("NewChipCollector(%s, %d).DecreaseChip().DecreaseChip().IncreaseChip() = Max Chip:%d; expected %d", tt.chips, tt.total, result.MaxChip, tt.expectedMaxChip)
			}
		})
	}
}
