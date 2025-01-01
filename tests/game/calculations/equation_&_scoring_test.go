package calculations_test

import (
	"A-MATH/game/calculations"
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"testing"
)

func TestCorrectProcessCalculatingAndScore(t *testing.T) {
	tests := []struct {
		name            string
		squareType      []string
		chip            []string
		isPlacedOnBoard []bool
		expectedScore   int
	}{
		{
			name: "Correct Equation(3=3)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"3", "=", "3"},
			isPlacedOnBoard: []bool{false, true, false},
			expectedScore:   3,
		},
		{
			name: "Correct Equation(18=18)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
			},
			chip:            []string{"1", "8", "=", "18"},
			isPlacedOnBoard: []bool{false, false, true, false},
			expectedScore:   12,
		},
		{
			name: "Correct Equation(144÷12=12)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"1", "4", "4", "/", "12", "=", "1", "2"},
			isPlacedOnBoard: []bool{false, false, false, true, false, false, false, false},
			expectedScore:   18,
		},
		{
			name: "Correct Equation(14-7+11=18)",
			squareType: []string{
				string(constants.CenterSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.RedSquare),
			},
			chip:            []string{"14", "-", "7", "+", "11", "=", "1", "8"},
			isPlacedOnBoard: []bool{false, false, false, false, false, false, false, false},
			expectedScore:   90,
		},
		{
			name: "Correct Equation(9×2=18=14-7+11=18)",
			squareType: []string{
				string(constants.RedSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.CenterSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.RedSquare),
			},
			chip: []string{"9", "*", "2", "=", "1", "8", "=",
				"14", "-", "7", "+", "11", "=", "1", "8"},
			isPlacedOnBoard: []bool{false, false, false, false, false, false,
				false, true, true, true, true, true, true, true, true},
			expectedScore: 87,
		},
		{
			name: "Correct Equation(15+7-12=10)",
			squareType: []string{
				string(constants.BlueSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.BlueSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"15", "+", "7", "-", "1", "2", "=", "10"},
			isPlacedOnBoard: []bool{false, false, false, false, false, false, true, false},
			expectedScore:   26,
		},
		{
			name: "Correct Equation(96÷18×0+9-0÷43=9)",
			squareType: []string{
				string(constants.RedSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.CenterSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.RedSquare),
			},
			chip: []string{"9", "6", "/", "18", "*", "0", "+",
				"9", "-", "0", "/", "4", "3", "=", "9"},
			isPlacedOnBoard: []bool{false, false, false, false, false, false,
				false, true, true, true, true, true, true, true, true},
			expectedScore: 96,
		},
		{
			name: "Correct Equation(-4+2=-6+4)",
			squareType: []string{
				string(constants.YellowSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.YellowSquare),
			},
			chip:            []string{"-", "4", "+", "2", "=", "-", "6", "+", "4"},
			isPlacedOnBoard: []bool{false, false, false, false, true, false, false, false, false},
			expectedScore:   64,
		},
		{
			name: "Correct Equation(-3×3=9-18)",
			squareType: []string{
				string(constants.YellowSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.YellowSquare),
			},
			chip:            []string{"-", "3", "*", "3", "=", "9", "-", "1", "8"},
			isPlacedOnBoard: []bool{false, false, false, false, true, false, false, false, false},
			expectedScore:   56,
		},
		{
			name: "Correct Equation(-1÷3+4=11÷3)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.BlueSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.BlueSquare),
			},
			chip: []string{"-", "1", "/", "3", "+", "4", "=", "11", "/", "3"},
			isPlacedOnBoard: []bool{true, false, false, false, false,
				false, false, false, false, true},
			expectedScore: 20,
		},
		{
			name: "Correct Equation(81÷9+6=15)",
			squareType: []string{
				string(constants.YellowSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.BlueSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"8", "1", "/", "9", "+", "6", "=", "15"},
			isPlacedOnBoard: []bool{false, false, false, false, false, false, true, false},
			expectedScore:   40,
		},
		{
			name: "Correct Equation(19+17-12=24)",
			squareType: []string{
				string(constants.BlueSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.BlueSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"19", "-", "12", "+", "17", "=", "2", "4"},
			isPlacedOnBoard: []bool{false, false, false, true, false, false, false, true},
			expectedScore:   50,
		},
		{
			name: "Correct Equation(27×3÷1=81)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.RedSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.RedSquare),
			},
			chip:            []string{"2", "7", "*", "3", "/", "1", "=", "8", "1"},
			isPlacedOnBoard: []bool{false, false, false, true, false, false, false, false, false},
			expectedScore:   126,
		},
		{
			name: "Correct Equation(2÷6=1÷3)",
			squareType: []string{
				string(constants.BlueSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.BlueSquare),
			},
			chip:            []string{"2", "/", "6", "=", "1", "/", "3"},
			isPlacedOnBoard: []bool{false, false, false, true, false, false, false},
			expectedScore:   14,
		},
		{
			name: "Correct Equation(36×2=80-8)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.YellowSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"3", "6", "*", "2", "=", "8", "0", "-", "8"},
			isPlacedOnBoard: []bool{false, true, false, false, false, false, false, false, false},
			expectedScore:   30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chipForCalculating []models.ChipForCalculating
			for i := range len(tt.squareType) {
				chipForCalculating = append(chipForCalculating, models.NewChipForCalculating(
					tt.squareType[i],
					components.NewChip(tt.chip[i]),
					tt.isPlacedOnBoard[i],
				))
			}

			result := calculations.ProcessCalculating(chipForCalculating)
			score := calculations.Score(chipForCalculating)

			if result != true {
				t.Errorf("%v is return %v but expect result: %v", tt.name, result, true)
			}

			if score != tt.expectedScore {
				t.Errorf("%v is return %d but expect score: %d", tt.name, score, tt.expectedScore)
			}
		})
	}

}

func TestIncorrectProcessCalculating(t *testing.T) {
	tests := []struct {
		name            string
		squareType      []string
		chip            []string
		isPlacedOnBoard []bool
	}{
		{
			name: "Incorrect Equation(4=3)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"4", "=", "3"},
			isPlacedOnBoard: []bool{false, true, false},
		},
		{
			name: "Incorrect Equation(152÷12=12)",
			squareType: []string{
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
			},
			chip:            []string{"1", "5", "2", "/", "12", "=", "1", "2"},
			isPlacedOnBoard: []bool{false, false, false, true, false, false, false, false},
		},
		{
			name: "Incorrect Equation(18-7+17=27)",
			squareType: []string{
				string(constants.CenterSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.RedSquare),
			},
			chip:            []string{"18", "-", "7", "+", "17", "=", "2", "7"},
			isPlacedOnBoard: []bool{false, false, false, false, false, false, false, false},
		},
		{
			name: "Incorrect Equation(7×2=18=14-7+11=18)",
			squareType: []string{
				string(constants.RedSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.CenterSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.RedSquare),
			},
			chip: []string{"7", "*", "2", "=", "1", "8", "=",
				"14", "-", "7", "+", "11", "=", "1", "8"},
			isPlacedOnBoard: []bool{false, false, false, false, false, false,
				false, true, true, true, true, true, true, true, true},
		},
		{
			name: "Incorrect Equation(86×0÷9=47×0+3=3)",
			squareType: []string{
				string(constants.RedSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.CenterSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.RedSquare),
			},
			chip: []string{"8", "6", "*", "0", "/", "9", "=",
				"4", "7", "*", "0", "+", "3", "=", "3"},
			isPlacedOnBoard: []bool{true, true, true, true, true, true,
				true, true, false, false, false, false, false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chipForCalculating []models.ChipForCalculating
			for i := range len(tt.squareType) {
				chipForCalculating = append(chipForCalculating, models.NewChipForCalculating(
					tt.squareType[i],
					components.NewChip(tt.chip[i]),
					tt.isPlacedOnBoard[i],
				))
			}

			result := calculations.ProcessCalculating(chipForCalculating)

			if result != false {
				t.Errorf("%v is return %v but expect result: %v", tt.name, result, false)
			}

		})
	}

}
