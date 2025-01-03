package components_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"testing"
)

func TestNewSquare(t *testing.T) {
	tests := []struct {
		name                string
		Coordinate          [2]int
		expectedSquareType  string
		expectedChipPlaceOn components.Chip
	}{
		{"Square: A1", [2]int{1, 1}, string(constants.RedSquare), components.Chip{}},
		{"Square: B1", [2]int{2, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C1", [2]int{3, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D1", [2]int{4, 1}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: E1", [2]int{5, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F1", [2]int{6, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G1", [2]int{7, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H1", [2]int{8, 1}, string(constants.RedSquare), components.Chip{}},
		{"Square: I1", [2]int{9, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J1", [2]int{10, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K1", [2]int{11, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L1", [2]int{12, 1}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: M1", [2]int{13, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N1", [2]int{14, 1}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O1", [2]int{15, 1}, string(constants.RedSquare), components.Chip{}},

		{"Square: A2", [2]int{1, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B2", [2]int{2, 2}, string(constants.YellowSquare), components.Chip{}},
		{"Square: C2", [2]int{3, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D2", [2]int{4, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E2", [2]int{5, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F2", [2]int{6, 2}, string(constants.BlueSquare), components.Chip{}},
		{"Square: G2", [2]int{7, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H2", [2]int{8, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I2", [2]int{9, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J2", [2]int{10, 2}, string(constants.BlueSquare), components.Chip{}},
		{"Square: K2", [2]int{11, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L2", [2]int{12, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M2", [2]int{13, 2}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N2", [2]int{14, 2}, string(constants.YellowSquare), components.Chip{}},
		{"Square: O2", [2]int{15, 2}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A3", [2]int{1, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B3", [2]int{2, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C3", [2]int{3, 3}, string(constants.YellowSquare), components.Chip{}},
		{"Square: D3", [2]int{4, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E3", [2]int{5, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F3", [2]int{6, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G3", [2]int{7, 3}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: H3", [2]int{8, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I3", [2]int{9, 3}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: J3", [2]int{10, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K3", [2]int{11, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L3", [2]int{12, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M3", [2]int{13, 3}, string(constants.YellowSquare), components.Chip{}},
		{"Square: N3", [2]int{14, 3}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O3", [2]int{15, 3}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A4", [2]int{1, 4}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: B4", [2]int{2, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C4", [2]int{3, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D4", [2]int{4, 4}, string(constants.YellowSquare), components.Chip{}},
		{"Square: E4", [2]int{5, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F4", [2]int{6, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G4", [2]int{7, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H4", [2]int{8, 4}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: I4", [2]int{9, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J4", [2]int{10, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K4", [2]int{11, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L4", [2]int{12, 4}, string(constants.YellowSquare), components.Chip{}},
		{"Square: M4", [2]int{13, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N4", [2]int{14, 4}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O4", [2]int{15, 4}, string(constants.OrangeSquare), components.Chip{}},

		{"Square: A5", [2]int{1, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B5", [2]int{2, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C5", [2]int{3, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D5", [2]int{4, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E5", [2]int{5, 5}, string(constants.BlueSquare), components.Chip{}},
		{"Square: F5", [2]int{6, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G5", [2]int{7, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H5", [2]int{8, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I5", [2]int{9, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J5", [2]int{10, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K5", [2]int{11, 5}, string(constants.BlueSquare), components.Chip{}},
		{"Square: L5", [2]int{12, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M5", [2]int{13, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N5", [2]int{14, 5}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O5", [2]int{15, 5}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A6", [2]int{1, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B6", [2]int{2, 6}, string(constants.BlueSquare), components.Chip{}},
		{"Square: C6", [2]int{3, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D6", [2]int{4, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E6", [2]int{5, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F6", [2]int{6, 6}, string(constants.BlueSquare), components.Chip{}},
		{"Square: G6", [2]int{7, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H6", [2]int{8, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I6", [2]int{9, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J6", [2]int{10, 6}, string(constants.BlueSquare), components.Chip{}},
		{"Square: K6", [2]int{11, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L6", [2]int{12, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M6", [2]int{13, 6}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N6", [2]int{14, 6}, string(constants.BlueSquare), components.Chip{}},
		{"Square: O6", [2]int{15, 6}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A7", [2]int{1, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B7", [2]int{2, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C7", [2]int{3, 7}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: D7", [2]int{4, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E7", [2]int{5, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F7", [2]int{6, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G7", [2]int{7, 7}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: H7", [2]int{8, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I7", [2]int{9, 7}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: J7", [2]int{10, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K7", [2]int{11, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L7", [2]int{12, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M7", [2]int{13, 7}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: N7", [2]int{14, 7}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O7", [2]int{15, 7}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A8", [2]int{1, 8}, string(constants.RedSquare), components.Chip{}},
		{"Square: B8", [2]int{2, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C8", [2]int{3, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D8", [2]int{4, 8}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: E8", [2]int{5, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F8", [2]int{6, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G8", [2]int{7, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H8", [2]int{8, 8}, string(constants.CenterSquare), components.Chip{}},
		{"Square: I8", [2]int{9, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J8", [2]int{10, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K8", [2]int{11, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L8", [2]int{12, 8}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: M8", [2]int{13, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N8", [2]int{14, 8}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O8", [2]int{15, 8}, string(constants.RedSquare), components.Chip{}},

		{"Square: A9", [2]int{1, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B9", [2]int{2, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C9", [2]int{3, 9}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: D9", [2]int{4, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E9", [2]int{5, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F9", [2]int{6, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G9", [2]int{7, 9}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: H9", [2]int{8, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I9", [2]int{9, 9}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: J9", [2]int{10, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K9", [2]int{11, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L9", [2]int{12, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M9", [2]int{13, 9}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: N9", [2]int{14, 9}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O9", [2]int{15, 9}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A10", [2]int{1, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B10", [2]int{2, 10}, string(constants.BlueSquare), components.Chip{}},
		{"Square: C10", [2]int{3, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D10", [2]int{4, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E10", [2]int{5, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F10", [2]int{6, 10}, string(constants.BlueSquare), components.Chip{}},
		{"Square: G10", [2]int{7, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H10", [2]int{8, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I10", [2]int{9, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J10", [2]int{10, 10}, string(constants.BlueSquare), components.Chip{}},
		{"Square: K10", [2]int{11, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L10", [2]int{12, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M10", [2]int{13, 10}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N10", [2]int{14, 10}, string(constants.BlueSquare), components.Chip{}},
		{"Square: O10", [2]int{15, 10}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A11", [2]int{1, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B11", [2]int{2, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C11", [2]int{3, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D11", [2]int{4, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E11", [2]int{5, 11}, string(constants.BlueSquare), components.Chip{}},
		{"Square: F11", [2]int{6, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G11", [2]int{7, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H11", [2]int{8, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I11", [2]int{9, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J11", [2]int{10, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K11", [2]int{11, 11}, string(constants.BlueSquare), components.Chip{}},
		{"Square: L11", [2]int{12, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M11", [2]int{13, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N11", [2]int{14, 11}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O11", [2]int{15, 11}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A12", [2]int{1, 12}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: B12", [2]int{2, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C12", [2]int{3, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D12", [2]int{4, 12}, string(constants.YellowSquare), components.Chip{}},
		{"Square: E12", [2]int{5, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F12", [2]int{6, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G12", [2]int{7, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H12", [2]int{8, 12}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: I12", [2]int{9, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J12", [2]int{10, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K12", [2]int{11, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L12", [2]int{12, 12}, string(constants.YellowSquare), components.Chip{}},
		{"Square: M12", [2]int{13, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N12", [2]int{14, 12}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O12", [2]int{15, 12}, string(constants.OrangeSquare), components.Chip{}},

		{"Square: A13", [2]int{1, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B13", [2]int{2, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C13", [2]int{3, 13}, string(constants.YellowSquare), components.Chip{}},
		{"Square: D13", [2]int{4, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E13", [2]int{5, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F13", [2]int{6, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G13", [2]int{7, 13}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: H13", [2]int{8, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I13", [2]int{9, 13}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: J13", [2]int{10, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K13", [2]int{11, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L13", [2]int{12, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M13", [2]int{13, 13}, string(constants.YellowSquare), components.Chip{}},
		{"Square: N13", [2]int{14, 13}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O13", [2]int{15, 13}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A14", [2]int{1, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: B14", [2]int{2, 14}, string(constants.YellowSquare), components.Chip{}},
		{"Square: C14", [2]int{3, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D14", [2]int{4, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: E14", [2]int{5, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F14", [2]int{6, 14}, string(constants.BlueSquare), components.Chip{}},
		{"Square: G14", [2]int{7, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H14", [2]int{8, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: I14", [2]int{9, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J14", [2]int{10, 14}, string(constants.BlueSquare), components.Chip{}},
		{"Square: K14", [2]int{11, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L14", [2]int{12, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: M14", [2]int{13, 14}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N14", [2]int{14, 14}, string(constants.YellowSquare), components.Chip{}},
		{"Square: O14", [2]int{15, 14}, string(constants.NormalSquare), components.Chip{}},

		{"Square: A15", [2]int{1, 15}, string(constants.RedSquare), components.Chip{}},
		{"Square: B15", [2]int{2, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: C15", [2]int{3, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: D15", [2]int{4, 15}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: E15", [2]int{5, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: F15", [2]int{6, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: G15", [2]int{7, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: H15", [2]int{8, 15}, string(constants.RedSquare), components.Chip{}},
		{"Square: I15", [2]int{9, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: J15", [2]int{10, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: K15", [2]int{11, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: L15", [2]int{12, 15}, string(constants.OrangeSquare), components.Chip{}},
		{"Square: M15", [2]int{13, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: N15", [2]int{14, 15}, string(constants.NormalSquare), components.Chip{}},
		{"Square: O15", [2]int{15, 15}, string(constants.RedSquare), components.Chip{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := components.NewSquare(tt.Coordinate)

			if result.SquareType != tt.expectedSquareType {
				t.Errorf("NewSquare(%d, %d) = SquareType:%s; expected %s", tt.Coordinate[0], tt.Coordinate[1], result.SquareType, tt.expectedSquareType)
			}

			if result.ChipPlaceOn != tt.expectedChipPlaceOn {
				t.Errorf("NewSquare(%d, %d) = ChipPlaceOn:%s; expected %s", tt.Coordinate[0], tt.Coordinate[1], result.ChipPlaceOn, tt.expectedChipPlaceOn)
			}
		})
	}
}
