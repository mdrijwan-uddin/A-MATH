package components_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"testing"
)

func TestNewBoard(t *testing.T) {
	tests := [][]struct {
		name                string
		Coordinate          [2]int
		expectedSquareType  string
		expectedChipPlaceOn components.Chip
	}{
		{
			{"Board Square: A1", [2]int{1, 1}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: B1", [2]int{2, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C1", [2]int{3, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D1", [2]int{4, 1}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: E1", [2]int{5, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F1", [2]int{6, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G1", [2]int{7, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H1", [2]int{8, 1}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: I1", [2]int{9, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J1", [2]int{10, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K1", [2]int{11, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L1", [2]int{12, 1}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: M1", [2]int{13, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N1", [2]int{14, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O1", [2]int{15, 1}, string(constants.RedSquare), components.Chip{}},
		},
		{
			{"Board Square: A2", [2]int{1, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B2", [2]int{2, 2}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: C2", [2]int{3, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D2", [2]int{4, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E2", [2]int{5, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F2", [2]int{6, 2}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G2", [2]int{7, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H2", [2]int{8, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I2", [2]int{9, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J2", [2]int{10, 2}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K2", [2]int{11, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L2", [2]int{12, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M2", [2]int{13, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N2", [2]int{14, 2}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: O2", [2]int{15, 2}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A3", [2]int{1, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B3", [2]int{2, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C3", [2]int{3, 3}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: D3", [2]int{4, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E3", [2]int{5, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F3", [2]int{6, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G3", [2]int{7, 3}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H3", [2]int{8, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I3", [2]int{9, 3}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J3", [2]int{10, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K3", [2]int{11, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L3", [2]int{12, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M3", [2]int{13, 3}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: N3", [2]int{14, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O3", [2]int{15, 3}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A4", [2]int{1, 4}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: B4", [2]int{2, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C4", [2]int{3, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D4", [2]int{4, 4}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: E4", [2]int{5, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F4", [2]int{6, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G4", [2]int{7, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H4", [2]int{8, 4}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: I4", [2]int{9, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J4", [2]int{10, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K4", [2]int{11, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L4", [2]int{12, 4}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: M4", [2]int{13, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N4", [2]int{14, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O4", [2]int{15, 4}, string(constants.OrangeSquare), components.Chip{}},
		},
		{
			{"Board Square: A5", [2]int{1, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B5", [2]int{2, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C5", [2]int{3, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D5", [2]int{4, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E5", [2]int{5, 5}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: F5", [2]int{6, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G5", [2]int{7, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H5", [2]int{8, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I5", [2]int{9, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J5", [2]int{10, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K5", [2]int{11, 5}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: L5", [2]int{12, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M5", [2]int{13, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N5", [2]int{14, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O5", [2]int{15, 5}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A6", [2]int{1, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B6", [2]int{2, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: C6", [2]int{3, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D6", [2]int{4, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E6", [2]int{5, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F6", [2]int{6, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G6", [2]int{7, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H6", [2]int{8, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I6", [2]int{9, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J6", [2]int{10, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K6", [2]int{11, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L6", [2]int{12, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M6", [2]int{13, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N6", [2]int{14, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: O6", [2]int{15, 6}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A7", [2]int{1, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B7", [2]int{2, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C7", [2]int{3, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: D7", [2]int{4, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E7", [2]int{5, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F7", [2]int{6, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G7", [2]int{7, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H7", [2]int{8, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I7", [2]int{9, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J7", [2]int{10, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K7", [2]int{11, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L7", [2]int{12, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M7", [2]int{13, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: N7", [2]int{14, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O7", [2]int{15, 7}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A8", [2]int{1, 8}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: B8", [2]int{2, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C8", [2]int{3, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D8", [2]int{4, 8}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: E8", [2]int{5, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F8", [2]int{6, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G8", [2]int{7, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H8", [2]int{8, 8}, string(constants.CenterSquare), components.Chip{}},
			{"Board Square: I8", [2]int{9, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J8", [2]int{10, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K8", [2]int{11, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L8", [2]int{12, 8}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: M8", [2]int{13, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N8", [2]int{14, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O8", [2]int{15, 8}, string(constants.RedSquare), components.Chip{}},
		},
		{
			{"Board Square: A9", [2]int{1, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B9", [2]int{2, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C9", [2]int{3, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: D9", [2]int{4, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E9", [2]int{5, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F9", [2]int{6, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G9", [2]int{7, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H9", [2]int{8, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I9", [2]int{9, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J9", [2]int{10, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K9", [2]int{11, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L9", [2]int{12, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M9", [2]int{13, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: N9", [2]int{14, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O9", [2]int{15, 9}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A10", [2]int{1, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B10", [2]int{2, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: C10", [2]int{3, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D10", [2]int{4, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E10", [2]int{5, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F10", [2]int{6, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G10", [2]int{7, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H10", [2]int{8, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I10", [2]int{9, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J10", [2]int{10, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K10", [2]int{11, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L10", [2]int{12, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M10", [2]int{13, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N10", [2]int{14, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: O10", [2]int{15, 10}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A11", [2]int{1, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B11", [2]int{2, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C11", [2]int{3, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D11", [2]int{4, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E11", [2]int{5, 11}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: F11", [2]int{6, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G11", [2]int{7, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H11", [2]int{8, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I11", [2]int{9, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J11", [2]int{10, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K11", [2]int{11, 11}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: L11", [2]int{12, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M11", [2]int{13, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N11", [2]int{14, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O11", [2]int{15, 11}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A12", [2]int{1, 12}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: B12", [2]int{2, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C12", [2]int{3, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D12", [2]int{4, 12}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: E12", [2]int{5, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F12", [2]int{6, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G12", [2]int{7, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H12", [2]int{8, 12}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: I12", [2]int{9, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J12", [2]int{10, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K12", [2]int{11, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L12", [2]int{12, 12}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: M12", [2]int{13, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N12", [2]int{14, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O12", [2]int{15, 12}, string(constants.OrangeSquare), components.Chip{}},
		},
		{
			{"Board Square: A13", [2]int{1, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B13", [2]int{2, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C13", [2]int{3, 13}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: D13", [2]int{4, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E13", [2]int{5, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F13", [2]int{6, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G13", [2]int{7, 13}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H13", [2]int{8, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I13", [2]int{9, 13}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J13", [2]int{10, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K13", [2]int{11, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L13", [2]int{12, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M13", [2]int{13, 13}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: N13", [2]int{14, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O13", [2]int{15, 13}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A14", [2]int{1, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B14", [2]int{2, 14}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: C14", [2]int{3, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D14", [2]int{4, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E14", [2]int{5, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F14", [2]int{6, 14}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G14", [2]int{7, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H14", [2]int{8, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I14", [2]int{9, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J14", [2]int{10, 14}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K14", [2]int{11, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L14", [2]int{12, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M14", [2]int{13, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N14", [2]int{14, 14}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: O14", [2]int{15, 14}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A15", [2]int{1, 15}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: B15", [2]int{2, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C15", [2]int{3, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D15", [2]int{4, 15}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: E15", [2]int{5, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F15", [2]int{6, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G15", [2]int{7, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H15", [2]int{8, 15}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: I15", [2]int{9, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J15", [2]int{10, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K15", [2]int{11, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L15", [2]int{12, 15}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: M15", [2]int{13, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N15", [2]int{14, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O15", [2]int{15, 15}, string(constants.RedSquare), components.Chip{}},
		},
	}

	board := components.NewBoard()
	testEmptyBoard(t, board)

	for j, test := range tests {
		for i, tt := range test {
			t.Run(tt.name, func(t *testing.T) {
				testSquareInBoard(t, board, i, j, tt.Coordinate, tt.expectedSquareType, tt.expectedChipPlaceOn)
			})
		}
	}
}

func TestNewBoardWithChipAdded(t *testing.T) {
	tests := [][]struct {
		name                string
		Coordinate          [2]int
		expectedSquareType  string
		expectedChipPlaceOn components.Chip
	}{
		{
			{"Board Square: A1", [2]int{1, 1}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: B1", [2]int{2, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C1", [2]int{3, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D1", [2]int{4, 1}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: E1", [2]int{5, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F1", [2]int{6, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G1", [2]int{7, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H1", [2]int{8, 1}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: I1", [2]int{9, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J1", [2]int{10, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K1", [2]int{11, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L1", [2]int{12, 1}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: M1", [2]int{13, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N1", [2]int{14, 1}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O1", [2]int{15, 1}, string(constants.RedSquare), components.Chip{}},
		},
		{
			{"Board Square: A2", [2]int{1, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B2", [2]int{2, 2}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: C2", [2]int{3, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D2", [2]int{4, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E2", [2]int{5, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F2", [2]int{6, 2}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G2", [2]int{7, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H2", [2]int{8, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I2", [2]int{9, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J2", [2]int{10, 2}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K2", [2]int{11, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L2", [2]int{12, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M2", [2]int{13, 2}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N2", [2]int{14, 2}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: O2", [2]int{15, 2}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A3", [2]int{1, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B3", [2]int{2, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C3", [2]int{3, 3}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: D3", [2]int{4, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E3", [2]int{5, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F3", [2]int{6, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G3", [2]int{7, 3}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H3", [2]int{8, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I3", [2]int{9, 3}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J3", [2]int{10, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K3", [2]int{11, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L3", [2]int{12, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M3", [2]int{13, 3}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: N3", [2]int{14, 3}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O3", [2]int{15, 3}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A4", [2]int{1, 4}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: B4", [2]int{2, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C4", [2]int{3, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D4", [2]int{4, 4}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: E4", [2]int{5, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F4", [2]int{6, 4}, string(constants.NormalSquare), components.NewChip("5")},
			{"Board Square: G4", [2]int{7, 4}, string(constants.NormalSquare), components.NewChip("3")},
			{"Board Square: H4", [2]int{8, 4}, string(constants.OrangeSquare), components.NewChip("=")},
			{"Board Square: I4", [2]int{9, 4}, string(constants.NormalSquare), components.NewChip("5")},
			{"Board Square: J4", [2]int{10, 4}, string(constants.NormalSquare), components.NewChip("3")},
			{"Board Square: K4", [2]int{11, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L4", [2]int{12, 4}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: M4", [2]int{13, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N4", [2]int{14, 4}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O4", [2]int{15, 4}, string(constants.OrangeSquare), components.Chip{}},
		},
		{
			{"Board Square: A5", [2]int{1, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B5", [2]int{2, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C5", [2]int{3, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D5", [2]int{4, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E5", [2]int{5, 5}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: F5", [2]int{6, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G5", [2]int{7, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H5", [2]int{8, 5}, string(constants.NormalSquare), components.NewChip("1")},
			{"Board Square: I5", [2]int{9, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J5", [2]int{10, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K5", [2]int{11, 5}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: L5", [2]int{12, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M5", [2]int{13, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N5", [2]int{14, 5}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O5", [2]int{15, 5}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A6", [2]int{1, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B6", [2]int{2, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: C6", [2]int{3, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D6", [2]int{4, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E6", [2]int{5, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F6", [2]int{6, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G6", [2]int{7, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H6", [2]int{8, 6}, string(constants.NormalSquare), components.NewChip("4")},
			{"Board Square: I6", [2]int{9, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J6", [2]int{10, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K6", [2]int{11, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L6", [2]int{12, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M6", [2]int{13, 6}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N6", [2]int{14, 6}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: O6", [2]int{15, 6}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A7", [2]int{1, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B7", [2]int{2, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C7", [2]int{3, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: D7", [2]int{4, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E7", [2]int{5, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F7", [2]int{6, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G7", [2]int{7, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H7", [2]int{8, 7}, string(constants.NormalSquare), components.NewChip("=")},
			{"Board Square: I7", [2]int{9, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J7", [2]int{10, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K7", [2]int{11, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L7", [2]int{12, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M7", [2]int{13, 7}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: N7", [2]int{14, 7}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O7", [2]int{15, 7}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A8", [2]int{1, 8}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: B8", [2]int{2, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C8", [2]int{3, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D8", [2]int{4, 8}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: E8", [2]int{5, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F8", [2]int{6, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G8", [2]int{7, 8}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H8", [2]int{8, 8}, string(constants.CenterSquare), components.NewChip("1")},
			{"Board Square: I8", [2]int{9, 8}, string(constants.NormalSquare), components.NewChip("+")},
			{"Board Square: J8", [2]int{10, 8}, string(constants.NormalSquare), components.NewChip("0")},
			{"Board Square: K8", [2]int{11, 8}, string(constants.NormalSquare), components.NewChip("*")},
			{"Board Square: L8", [2]int{12, 8}, string(constants.OrangeSquare), components.NewChip("5")},
			{"Board Square: M8", [2]int{13, 8}, string(constants.NormalSquare), components.NewChip("3")},
			{"Board Square: N8", [2]int{14, 8}, string(constants.NormalSquare), components.NewChip("=")},
			{"Board Square: O8", [2]int{15, 8}, string(constants.RedSquare), components.NewChip("1")},
		},
		{
			{"Board Square: A9", [2]int{1, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B9", [2]int{2, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C9", [2]int{3, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: D9", [2]int{4, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E9", [2]int{5, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F9", [2]int{6, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G9", [2]int{7, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H9", [2]int{8, 9}, string(constants.NormalSquare), components.NewChip("4")},
			{"Board Square: I9", [2]int{9, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J9", [2]int{10, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K9", [2]int{11, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L9", [2]int{12, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M9", [2]int{13, 9}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: N9", [2]int{14, 9}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O9", [2]int{15, 9}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A10", [2]int{1, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B10", [2]int{2, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: C10", [2]int{3, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D10", [2]int{4, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E10", [2]int{5, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F10", [2]int{6, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G10", [2]int{7, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H10", [2]int{8, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I10", [2]int{9, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J10", [2]int{10, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K10", [2]int{11, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L10", [2]int{12, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M10", [2]int{13, 10}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N10", [2]int{14, 10}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: O10", [2]int{15, 10}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A11", [2]int{1, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B11", [2]int{2, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C11", [2]int{3, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D11", [2]int{4, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E11", [2]int{5, 11}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: F11", [2]int{6, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G11", [2]int{7, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H11", [2]int{8, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I11", [2]int{9, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J11", [2]int{10, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K11", [2]int{11, 11}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: L11", [2]int{12, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M11", [2]int{13, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N11", [2]int{14, 11}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O11", [2]int{15, 11}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A12", [2]int{1, 12}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: B12", [2]int{2, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C12", [2]int{3, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D12", [2]int{4, 12}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: E12", [2]int{5, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F12", [2]int{6, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G12", [2]int{7, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H12", [2]int{8, 12}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: I12", [2]int{9, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J12", [2]int{10, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K12", [2]int{11, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L12", [2]int{12, 12}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: M12", [2]int{13, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N12", [2]int{14, 12}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O12", [2]int{15, 12}, string(constants.OrangeSquare), components.Chip{}},
		},
		{
			{"Board Square: A13", [2]int{1, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B13", [2]int{2, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C13", [2]int{3, 13}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: D13", [2]int{4, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E13", [2]int{5, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F13", [2]int{6, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G13", [2]int{7, 13}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: H13", [2]int{8, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I13", [2]int{9, 13}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: J13", [2]int{10, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K13", [2]int{11, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L13", [2]int{12, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M13", [2]int{13, 13}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: N13", [2]int{14, 13}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O13", [2]int{15, 13}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A14", [2]int{1, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: B14", [2]int{2, 14}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: C14", [2]int{3, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D14", [2]int{4, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: E14", [2]int{5, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F14", [2]int{6, 14}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: G14", [2]int{7, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H14", [2]int{8, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: I14", [2]int{9, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J14", [2]int{10, 14}, string(constants.BlueSquare), components.Chip{}},
			{"Board Square: K14", [2]int{11, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L14", [2]int{12, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: M14", [2]int{13, 14}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N14", [2]int{14, 14}, string(constants.YellowSquare), components.Chip{}},
			{"Board Square: O14", [2]int{15, 14}, string(constants.NormalSquare), components.Chip{}},
		},
		{
			{"Board Square: A15", [2]int{1, 15}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: B15", [2]int{2, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: C15", [2]int{3, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: D15", [2]int{4, 15}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: E15", [2]int{5, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: F15", [2]int{6, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: G15", [2]int{7, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: H15", [2]int{8, 15}, string(constants.RedSquare), components.Chip{}},
			{"Board Square: I15", [2]int{9, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: J15", [2]int{10, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: K15", [2]int{11, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: L15", [2]int{12, 15}, string(constants.OrangeSquare), components.Chip{}},
			{"Board Square: M15", [2]int{13, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: N15", [2]int{14, 15}, string(constants.NormalSquare), components.Chip{}},
			{"Board Square: O15", [2]int{15, 15}, string(constants.RedSquare), components.Chip{}},
		},
	}

	board := components.NewBoard()
	AddChipToBoard(&board)

	var chipForPlacing []models.ChipForPlacing
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 1}, components.NewChip("7")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 2}, components.NewChip("*")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 3}, components.NewChip("2")))
	removedChips := RemoveChipFromBoard(&board, chipForPlacing)

	testChipRemovedFromBoard(t, chipForPlacing, removedChips)
	testFilledBoard(t, board)

	for j, test := range tests {
		for i, tt := range test {
			t.Run(tt.name, func(t *testing.T) {
				testSquareInBoard(t, board, i, j, tt.Coordinate, tt.expectedSquareType, tt.expectedChipPlaceOn)
			})
		}
	}
}

func testSquareInBoard(
	t *testing.T,
	board components.Board,
	i, j int,
	expectedCoord [2]int,
	expectedSquareType string,
	expectedChipPlaceOn components.Chip,
) {
	if !board.IsValidSquare(expectedCoord) {
		t.Errorf("Coordinate(%d, %d) should be valid", i+1, j+1)
	}

	square := board.GetSquare([2]int{i + 1, j + 1})
	if square.Coordinate != expectedCoord {
		t.Errorf("NewBoard() = Coordinate:(%d, %d) ; expected (%d, %d)", i+1, j+1, expectedCoord[0], expectedCoord[1])
	}

	if square.SquareType != expectedSquareType {
		t.Errorf("NewBoard(%d, %d) = SquareType:() ; expected (%v)", i+1, j+1, expectedSquareType)
	}

	if square.ChipPlaceOn != expectedChipPlaceOn {
		t.Errorf("NewBoard(%d, %d) = ChipPlaceOn:() ; expected (%v)", i+1, j+1, expectedChipPlaceOn)
	}
}

func testEmptyBoard(t *testing.T, board components.Board) {
	if !board.IsEmpty() {
		t.Errorf("NewBoard() should be empty")
	}
}

func testFilledBoard(t *testing.T, board components.Board) {
	if board.IsEmpty() {
		t.Errorf("Filed Board() should not be empty")
	}
}

func AddChipToBoard(board *components.Board) {
	var chipForPlacing []models.ChipForPlacing
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("+")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("0")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 8}, components.NewChip("*")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{12, 8}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{13, 8}, components.NewChip("3")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{14, 8}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{15, 8}, components.NewChip("1")))

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 1}, components.NewChip("7")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 2}, components.NewChip("*")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 3}, components.NewChip("2")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 4}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 5}, components.NewChip("1")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("4")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("=")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("4")))

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{6, 4}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{7, 4}, components.NewChip("3")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{9, 4}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{10, 4}, components.NewChip("3")))

	for _, chips := range chipForPlacing {
		board.Add(chips.Coordinate, chips.Chip)
	}
}

func RemoveChipFromBoard(board *components.Board, chipForPlacing []models.ChipForPlacing) []components.Chip {
	var removedChip []components.Chip
	for _, chip := range chipForPlacing {
		removedChip = append(removedChip, board.Remove(chip.Coordinate))
	}
	return removedChip
}

func testChipRemovedFromBoard(t *testing.T, chipForPlacing []models.ChipForPlacing, removedChips []components.Chip) {
	if len(chipForPlacing) != len(removedChips) {
		t.Errorf("Removed chips Board() got an Error")
	}

	for i, chip := range removedChips {
		if chipForPlacing[i].Chip != chip {
			t.Errorf("Removed chips Board() got an Error")
		}
	}
}
