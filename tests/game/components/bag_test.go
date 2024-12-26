package components_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"testing"
)

func TestNewBag(t *testing.T) {
	tests := []struct {
		name            string
		expectedChips   components.Chip
		expectedTotal   int
		expectedMaxChip int
	}{
		{"Chip in bag: Zero", components.NewChip(string(constants.Zero)), 5, 5},
		{"Chip in bag: One", components.NewChip(string(constants.One)), 6, 6},
		{"Chip in bag: Two", components.NewChip(string(constants.Two)), 6, 6},
		{"Chip in bag: Three", components.NewChip(string(constants.Three)), 5, 5},
		{"Chip in bag: Four", components.NewChip(string(constants.Four)), 5, 5},
		{"Chip in bag: Five", components.NewChip(string(constants.Five)), 4, 4},
		{"Chip in bag: Six", components.NewChip(string(constants.Six)), 4, 4},
		{"Chip in bag: Seven", components.NewChip(string(constants.Seven)), 4, 4},
		{"Chip in bag: Eight", components.NewChip(string(constants.Eight)), 4, 4},
		{"Chip in bag: Nine", components.NewChip(string(constants.Nine)), 4, 4},
		{"Chip in bag: Ten", components.NewChip(string(constants.Ten)), 2, 2},
		{"Chip in bag: Eleven", components.NewChip(string(constants.Eleven)), 1, 1},
		{"Chip in bag: Twelve", components.NewChip(string(constants.Twelve)), 2, 2},
		{"Chip in bag: Thirteen", components.NewChip(string(constants.Thirteen)), 1, 1},
		{"Chip in bag: Fourteen", components.NewChip(string(constants.Fourteen)), 1, 1},
		{"Chip in bag: Fifteen", components.NewChip(string(constants.Fifteen)), 1, 1},
		{"Chip in bag: Sixteen", components.NewChip(string(constants.Sixteen)), 1, 1},
		{"Chip in bag: Seventeen", components.NewChip(string(constants.Seventeen)), 1, 1},
		{"Chip in bag: Eighteen", components.NewChip(string(constants.Eighteen)), 1, 1},
		{"Chip in bag: Nineteen", components.NewChip(string(constants.Nineteen)), 1, 1},
		{"Chip in bag: Twenty", components.NewChip(string(constants.Twenty)), 1, 1},
		{"Chip in bag: Addition", components.NewChip(string(constants.Addition)), 4, 4},
		{"Chip in bag: Subtraction", components.NewChip(string(constants.Subtraction)), 4, 4},
		{"Chip in bag: Multiply", components.NewChip(string(constants.Multiply)), 4, 4},
		{"Chip in bag: Division", components.NewChip(string(constants.Division)), 4, 4},
		{"Chip in bag: Add_sub", components.NewChip(string(constants.Add_sub)), 5, 5},
		{"Chip in bag: Multi_divide", components.NewChip(string(constants.Multi_divide)), 4, 4},
		{"Chip in bag: Equal", components.NewChip(string(constants.Equal)), 11, 11},
		{"Chip in bag: Blank", components.NewChip(string(constants.Blank)), 4, 4},
	}

	result := components.NewBag()

	// Unit Test
	if result.TotalChipLeft != constants.MaxChipInBag {
		t.Errorf("NewBag() = TotalChipLeft:%d; expected: %d", result.TotalChipLeft, constants.MaxChipInBag)
	}

	if result.MaxChip != constants.MaxChipInBag {
		t.Errorf("NewBag() = MaxChip:%d; expected: %d", result.MaxChip, constants.MaxChipInBag)
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if result.ChipCollectors[i].Chips != tt.expectedChips {
				t.Errorf("NewBag() = Chip:%s; expected: %s", result.ChipCollectors[i].Chips, tt.expectedChips)
			}

			if result.ChipCollectors[i].Total != tt.expectedTotal {
				t.Errorf("NewBag() = Total:%d; expected: %d", result.ChipCollectors[i].Total, tt.expectedTotal)
			}

			if result.ChipCollectors[i].MaxChip != tt.expectedMaxChip {
				t.Errorf("NewBag() = MaxChip:%d; expected: %d", result.ChipCollectors[i].MaxChip, tt.expectedMaxChip)
			}
		})
	}
}

func TestBagDecrease(t *testing.T) {
	tests := []struct {
		name            string
		expectedChips   components.Chip
		expectedTotal   int
		expectedMaxChip int
	}{
		{"Chip in bag: Zero (after decreasing)", components.NewChip(string(constants.Zero)), 4, 5},
		{"Chip in bag: One (after decreasing)", components.NewChip(string(constants.One)), 5, 6},
		{"Chip in bag: Two (after decreasing)", components.NewChip(string(constants.Two)), 6, 6},
		{"Chip in bag: Three (after decreasing)", components.NewChip(string(constants.Three)), 0, 5},
		{"Chip in bag: Four (after decreasing)", components.NewChip(string(constants.Four)), 5, 5},
		{"Chip in bag: Five (after decreasing)", components.NewChip(string(constants.Five)), 4, 4},
		{"Chip in bag: Six (after decreasing)", components.NewChip(string(constants.Six)), 3, 4},
		{"Chip in bag: Seven (after decreasing)", components.NewChip(string(constants.Seven)), 2, 4},
		{"Chip in bag: Eight (after decreasing)", components.NewChip(string(constants.Eight)), 3, 4},
		{"Chip in bag: Nine (after decreasing)", components.NewChip(string(constants.Nine)), 4, 4},
		{"Chip in bag: Ten (after decreasing)", components.NewChip(string(constants.Ten)), 2, 2},
		{"Chip in bag: Eleven (after decreasing)", components.NewChip(string(constants.Eleven)), 1, 1},
		{"Chip in bag: Twelve (after decreasing)", components.NewChip(string(constants.Twelve)), 2, 2},
		{"Chip in bag: Thirteen (after decreasing)", components.NewChip(string(constants.Thirteen)), 1, 1},
		{"Chip in bag: Fourteen (after decreasing)", components.NewChip(string(constants.Fourteen)), 0, 1},
		{"Chip in bag: Fifteen (after decreasing)", components.NewChip(string(constants.Fifteen)), 1, 1},
		{"Chip in bag: Sixteen (after decreasing)", components.NewChip(string(constants.Sixteen)), 1, 1},
		{"Chip in bag: Seventeen (after decreasing)", components.NewChip(string(constants.Seventeen)), 1, 1},
		{"Chip in bag: Eighteen (after decreasing)", components.NewChip(string(constants.Eighteen)), 1, 1},
		{"Chip in bag: Nineteen (after decreasing)", components.NewChip(string(constants.Nineteen)), 1, 1},
		{"Chip in bag: Twenty (after decreasing)", components.NewChip(string(constants.Twenty)), 1, 1},
		{"Chip in bag: Addition (after decreasing)", components.NewChip(string(constants.Addition)), 2, 4},
		{"Chip in bag: Subtraction (after decreasing)", components.NewChip(string(constants.Subtraction)), 4, 4},
		{"Chip in bag: Multiply (after decreasing)", components.NewChip(string(constants.Multiply)), 4, 4},
		{"Chip in bag: Division (after decreasing)", components.NewChip(string(constants.Division)), 3, 4},
		{"Chip in bag: Add_sub (after decreasing)", components.NewChip(string(constants.Add_sub)), 5, 5},
		{"Chip in bag: Multi_divide (after decreasing)", components.NewChip(string(constants.Multi_divide)), 3, 4},
		{"Chip in bag: Equal (after decreasing)", components.NewChip(string(constants.Equal)), 8, 11},
		{"Chip in bag: Blank (after decreasing)", components.NewChip(string(constants.Blank)), 3, 4},
	}

	result := components.NewBag()

	chipsForPick := mockChipsForPick()
	for _, chip := range chipsForPick {
		result.DecreaseChip(chip)
	}

	// Unit Test
	if result.TotalChipLeft != 80 {
		t.Errorf("Bag(after decreasing) = TotalChipLeft:%d; expected: %d", result.TotalChipLeft, 80)
	}

	if result.MaxChip != constants.MaxChipInBag {
		t.Errorf("Bag(after decreasing) = MaxChip:%d; expected: %d", result.MaxChip, constants.MaxChipInBag)
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if result.ChipCollectors[i].Chips != tt.expectedChips {
				t.Errorf("Bag(after decreasing) = Chip:%s; expected: %s", result.ChipCollectors[i].Chips, tt.expectedChips)
			}

			if result.ChipCollectors[i].Total != tt.expectedTotal {
				t.Errorf("Bag(after decreasing) = Total:%d; expected: %d", result.ChipCollectors[i].Total, tt.expectedTotal)
			}

			if result.ChipCollectors[i].MaxChip != tt.expectedMaxChip {
				t.Errorf("Bag(after decreasing) = MaxChip:%d; expected: %d", result.ChipCollectors[i].MaxChip, tt.expectedMaxChip)
			}
		})
	}
}

func TestBagIncreaseChip(t *testing.T) {
	tests := []struct {
		name            string
		expectedChips   components.Chip
		expectedTotal   int
		expectedMaxChip int
	}{
		{"Chip in bag: Zero (after increasing)", components.NewChip(string(constants.Zero)), 4, 5},
		{"Chip in bag: One (after increasing)", components.NewChip(string(constants.One)), 5, 6},
		{"Chip in bag: Two (after increasing)", components.NewChip(string(constants.Two)), 6, 6},
		{"Chip in bag: Three (after increasing)", components.NewChip(string(constants.Three)), 0, 5},
		{"Chip in bag: Four (after increasing)", components.NewChip(string(constants.Four)), 5, 5},
		{"Chip in bag: Five (after increasing)", components.NewChip(string(constants.Five)), 4, 4},
		{"Chip in bag: Six (after increasing)", components.NewChip(string(constants.Six)), 4, 4},
		{"Chip in bag: Seven (after increasing)", components.NewChip(string(constants.Seven)), 3, 4},
		{"Chip in bag: Eight (after increasing)", components.NewChip(string(constants.Eight)), 3, 4},
		{"Chip in bag: Nine (after increasing)", components.NewChip(string(constants.Nine)), 4, 4},
		{"Chip in bag: Ten (after increasing)", components.NewChip(string(constants.Ten)), 2, 2},
		{"Chip in bag: Eleven (after increasing)", components.NewChip(string(constants.Eleven)), 1, 1},
		{"Chip in bag: Twelve (after increasing)", components.NewChip(string(constants.Twelve)), 2, 2},
		{"Chip in bag: Thirteen (after increasing)", components.NewChip(string(constants.Thirteen)), 1, 1},
		{"Chip in bag: Fourteen (after increasing)", components.NewChip(string(constants.Fourteen)), 0, 1},
		{"Chip in bag: Fifteen (after increasing)", components.NewChip(string(constants.Fifteen)), 1, 1},
		{"Chip in bag: Sixteen (after increasing)", components.NewChip(string(constants.Sixteen)), 1, 1},
		{"Chip in bag: Seventeen (after increasing)", components.NewChip(string(constants.Seventeen)), 1, 1},
		{"Chip in bag: Eighteen (after increasing)", components.NewChip(string(constants.Eighteen)), 1, 1},
		{"Chip in bag: Nineteen (after increasing)", components.NewChip(string(constants.Nineteen)), 1, 1},
		{"Chip in bag: Twenty (after increasing)", components.NewChip(string(constants.Twenty)), 1, 1},
		{"Chip in bag: Addition (after increasing)", components.NewChip(string(constants.Addition)), 3, 4},
		{"Chip in bag: Subtraction (after increasing)", components.NewChip(string(constants.Subtraction)), 4, 4},
		{"Chip in bag: Multiply (after increasing)", components.NewChip(string(constants.Multiply)), 4, 4},
		{"Chip in bag: Division (after increasing)", components.NewChip(string(constants.Division)), 3, 4},
		{"Chip in bag: Add_sub (after increasing)", components.NewChip(string(constants.Add_sub)), 5, 5},
		{"Chip in bag: Multi_divide (after increasing)", components.NewChip(string(constants.Multi_divide)), 3, 4},
		{"Chip in bag: Equal (after increasing)", components.NewChip(string(constants.Equal)), 10, 11},
		{"Chip in bag: Blank (after increasing)", components.NewChip(string(constants.Blank)), 3, 4},
	}

	result := components.NewBag()

	chipsForPick := mockChipsForPick()

	for _, chip := range chipsForPick {
		result.DecreaseChip(chip)
	}

	chipsForPutBack := mockchipsForPutBack()
	for _, chip := range chipsForPutBack {
		result.IncreaseChip(chip)
	}

	// Unit Test
	if result.TotalChipLeft != 85 {
		t.Errorf("Bag(after increasing) = TotalChipLeft:%d; expected: %d", result.TotalChipLeft, 85)
	}

	if result.MaxChip != constants.MaxChipInBag {
		t.Errorf("Bag(after increasing) = MaxChip:%d; expected: %d", result.MaxChip, constants.MaxChipInBag)
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if result.ChipCollectors[i].Chips != tt.expectedChips {
				t.Errorf("Bag(after increasing) = Chip:%s; expected: %s", result.ChipCollectors[i].Chips, tt.expectedChips)
			}

			if result.ChipCollectors[i].Total != tt.expectedTotal {
				t.Errorf("Bag(after increasing) = Total:%d; expected: %d", result.ChipCollectors[i].Total, tt.expectedTotal)
			}

			if result.ChipCollectors[i].MaxChip != tt.expectedMaxChip {
				t.Errorf("Bag(after increasing) = MaxChip:%d; expected: %d", result.ChipCollectors[i].MaxChip, tt.expectedMaxChip)
			}
		})
	}
}

func mockChipsForPick() []components.Chip {
	var chipsForPick []components.Chip
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Addition)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Addition)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Blank)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Division)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Eight)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Equal)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Equal)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Equal)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Fourteen)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Multi_divide)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.One)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Six)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Seven)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Seven)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Three)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Three)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Three)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Three)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Three)))
	chipsForPick = append(chipsForPick, components.NewChip(string(constants.Zero)))
	return chipsForPick
}

func mockchipsForPutBack() []components.Chip {
	var chipsForPutBack []components.Chip
	chipsForPutBack = append(chipsForPutBack, components.NewChip(string(constants.Seven)))
	chipsForPutBack = append(chipsForPutBack, components.NewChip(string(constants.Addition)))
	chipsForPutBack = append(chipsForPutBack, components.NewChip(string(constants.Equal)))
	chipsForPutBack = append(chipsForPutBack, components.NewChip(string(constants.Six)))
	chipsForPutBack = append(chipsForPutBack, components.NewChip(string(constants.Equal)))
	return chipsForPutBack
}
