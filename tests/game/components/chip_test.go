package components_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"testing"
)

func TestNewChip(t *testing.T) {
	tests := []struct {
		name             string
		value            string
		expectedScore    int
		expectedChipType string
	}{
		{"Chip: Zero", string(constants.Zero), 1, string(constants.OneDigitNumberType)},
		{"Chip: One", string(constants.One), 1, string(constants.OneDigitNumberType)},
		{"Chip: Two", string(constants.Two), 1, string(constants.OneDigitNumberType)},
		{"Chip: Three", string(constants.Three), 1, string(constants.OneDigitNumberType)},
		{"Chip: Four", string(constants.Four), 2, string(constants.OneDigitNumberType)},
		{"Chip: Five", string(constants.Five), 2, string(constants.OneDigitNumberType)},
		{"Chip: Six", string(constants.Six), 2, string(constants.OneDigitNumberType)},
		{"Chip: Seven", string(constants.Seven), 2, string(constants.OneDigitNumberType)},
		{"Chip: Eight", string(constants.Eight), 2, string(constants.OneDigitNumberType)},
		{"Chip: Nine", string(constants.Nine), 2, string(constants.OneDigitNumberType)},
		{"Chip: Ten", string(constants.Ten), 3, string(constants.TwoDigitNumberType)},
		{"Chip: Eleven", string(constants.Eleven), 4, string(constants.TwoDigitNumberType)},
		{"Chip: Twelve", string(constants.Twelve), 3, string(constants.TwoDigitNumberType)},
		{"Chip: Thirteen", string(constants.Thirteen), 6, string(constants.TwoDigitNumberType)},
		{"Chip: Fourteen", string(constants.Fourteen), 4, string(constants.TwoDigitNumberType)},
		{"Chip: Fifteen", string(constants.Fifteen), 4, string(constants.TwoDigitNumberType)},
		{"Chip: Sixteen", string(constants.Sixteen), 4, string(constants.TwoDigitNumberType)},
		{"Chip: Seventeen", string(constants.Seventeen), 6, string(constants.TwoDigitNumberType)},
		{"Chip: Eighteen", string(constants.Eighteen), 4, string(constants.TwoDigitNumberType)},
		{"Chip: Nineteen", string(constants.Nineteen), 7, string(constants.TwoDigitNumberType)},
		{"Chip: Twenty", string(constants.Twenty), 5, string(constants.TwoDigitNumberType)},
		{"Chip: Addition", string(constants.Addition), 2, string(constants.OperatorType)},
		{"Chip: Subtraction", string(constants.Subtraction), 2, string(constants.OperatorType)},
		{"Chip: Multiply", string(constants.Multiply), 2, string(constants.OperatorType)},
		{"Chip: Division", string(constants.Division), 2, string(constants.OperatorType)},
		{"Chip: Add_sub", string(constants.Add_sub), 1, string(constants.AlterOperatorType)},
		{"Chip: Multi_divide", string(constants.Multi_divide), 1, string(constants.AlterOperatorType)},
		{"Chip: Equal", string(constants.Equal), 1, string(constants.EqualType)},
		{"Chip: Blank", string(constants.Blank), 0, string(constants.BlankType)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := components.NewChip(tt.value)

			if result.Score != tt.expectedScore {
				t.Errorf("NewChip(%s) = Score:%d; expected %d", tt.value, result.Score, tt.expectedScore)
			}

			if result.ChipType != tt.expectedChipType {
				t.Errorf("NewChip(%s) = Chip Type:%s; expected %s", tt.value, result.ChipType, tt.expectedChipType)
			}
		})
	}
}
