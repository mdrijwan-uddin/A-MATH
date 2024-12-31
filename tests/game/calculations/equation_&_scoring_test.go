package calculations_test

import (
	"A-MATH/game/calculations"
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/models"
	"testing"
)

func TestProcessCalculating(t *testing.T) {
	tests := []struct {
		name            string
		squareType      []string
		chip            []string
		isPlacedOnBoard []bool
		expectedResult  bool
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
			expectedResult:  true,
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
			expectedResult:  true,
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
			expectedResult:  true,
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
			expectedResult:  true,
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
			expectedResult: true,
			expectedScore:  87,
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
			expectedResult:  true,
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
			expectedResult: true,
			expectedScore:  96,
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
			expectedResult:  true,
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
			expectedResult:  true,
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
			expectedResult: true,
			expectedScore:  20,
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
			expectedResult:  true,
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
			expectedResult:  true,
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
			expectedResult:  true,
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
			expectedResult:  true,
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
			expectedResult:  true,
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

			if result != tt.expectedResult {
				t.Errorf("%v is return %v but expect result: %v", tt.name, result, tt.expectedResult)
			}

			if score != tt.expectedScore {
				t.Errorf("%v is return %d but expect score: %d", tt.name, score, tt.expectedScore)
			}
		})
	}

}
