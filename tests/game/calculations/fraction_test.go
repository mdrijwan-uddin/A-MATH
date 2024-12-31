package calculations_test

import (
	"A-MATH/game/calculations"
	"testing"
)

func TestNewFraction(t *testing.T) {
	tests := []struct {
		name             string
		baseFraction     calculations.Fraction
		comparedFraction calculations.Fraction
	}{
		{
			name:             "Correct same fraction (1÷5)",
			baseFraction:     calculations.NewFraction(1, 5),
			comparedFraction: calculations.NewFraction(1, 5),
		},
		{
			name:             "Correct same fraction (2÷5)",
			baseFraction:     calculations.NewFraction(2, 5),
			comparedFraction: calculations.NewFraction(4, 10),
		},
		{
			name:             "Correct same fraction (4÷7)",
			baseFraction:     calculations.NewFraction(8, 14),
			comparedFraction: calculations.NewFraction(12, 21),
		},
		{
			name:             "Correct same fraction (4÷3)",
			baseFraction:     calculations.NewFraction(4, 3),
			comparedFraction: calculations.NewFraction(8, 6),
		},
		{
			name:             "Correct same fraction (1÷6)",
			baseFraction:     calculations.NewFraction(2, 12),
			comparedFraction: calculations.NewFraction(1, 6),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v should equal to %v",
					tt.baseFraction, tt.comparedFraction,
				)
			}
		})
	}
}

func TestAddFraction(t *testing.T) {
	tests := []struct {
		name             string
		baseFraction     calculations.Fraction
		integerToAdd     int
		comparedFraction calculations.Fraction
	}{
		{
			name:             "Correct same fraction (6÷7)+4",
			baseFraction:     calculations.NewFraction(6, 7),
			integerToAdd:     4,
			comparedFraction: calculations.NewFraction(34, 7),
		},
		{
			name:             "Correct same fraction (7÷8)+5",
			baseFraction:     calculations.NewFraction(7, 8),
			integerToAdd:     5,
			comparedFraction: calculations.NewFraction(47, 8),
		},
		{
			name:             "Correct same fraction (9÷5)+4",
			baseFraction:     calculations.NewFraction(9, 5),
			integerToAdd:     4,
			comparedFraction: calculations.NewFraction(29, 5),
		},
		{
			name:             "Correct same fraction (6÷7)+4",
			baseFraction:     calculations.NewFraction(11, 6),
			integerToAdd:     7,
			comparedFraction: calculations.NewFraction(53, 6),
		},
		{
			name:             "Correct same fraction (3÷4)+2",
			baseFraction:     calculations.NewFraction(3, 4),
			integerToAdd:     2,
			comparedFraction: calculations.NewFraction(11, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.AddBy(tt.integerToAdd)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v add %d should equal to with %v",
					tt.baseFraction, tt.integerToAdd, tt.comparedFraction,
				)
			}
		})
	}
}

func TestFractionAddFraction(t *testing.T) {
	tests := []struct {
		name              string
		baseFraction      calculations.Fraction
		fractionToBeAdded calculations.Fraction
		comparedFraction  calculations.Fraction
	}{
		{
			name:              "Correct same fraction (4÷7)+(2÷7)",
			baseFraction:      calculations.NewFraction(4, 7),
			fractionToBeAdded: calculations.NewFraction(2, 7),
			comparedFraction:  calculations.NewFraction(6, 7),
		},
		{
			name:              "Correct same fraction (3÷10)+(5÷4)",
			baseFraction:      calculations.NewFraction(3, 10),
			fractionToBeAdded: calculations.NewFraction(5, 4),
			comparedFraction:  calculations.NewFraction(31, 20),
		},
		{
			name:              "Correct same fraction (7÷8)+(5÷12)",
			baseFraction:      calculations.NewFraction(7, 8),
			fractionToBeAdded: calculations.NewFraction(5, 12),
			comparedFraction:  calculations.NewFraction(31, 24),
		},
		{
			name:              "Correct same fraction (9÷10)+(4÷7)",
			baseFraction:      calculations.NewFraction(9, 10),
			fractionToBeAdded: calculations.NewFraction(4, 7),
			comparedFraction:  calculations.NewFraction(103, 70),
		},
		{
			name:              "Correct same fraction (7÷2)+(3÷8)",
			baseFraction:      calculations.NewFraction(7, 2),
			fractionToBeAdded: calculations.NewFraction(3, 8),
			comparedFraction:  calculations.NewFraction(31, 8),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.AddFractionBy(tt.fractionToBeAdded)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v add %v should equal to with %v",
					tt.baseFraction, tt.fractionToBeAdded, tt.comparedFraction,
				)
			}
		})
	}
}

func TestSubtractFraction(t *testing.T) {
	tests := []struct {
		name              string
		baseFraction      calculations.Fraction
		integerToSubtract int
		comparedFraction  calculations.Fraction
	}{
		{
			name:              "Correct same fraction (17÷4)-3",
			baseFraction:      calculations.NewFraction(17, 4),
			integerToSubtract: 3,
			comparedFraction:  calculations.NewFraction(5, 4),
		},
		{
			name:              "Correct same fraction (7÷5)-2",
			baseFraction:      calculations.NewFraction(7, 5),
			integerToSubtract: 2,
			comparedFraction:  calculations.NewFraction(-3, 5),
		},
		{
			name:              "Correct same fraction (9÷4)-3",
			baseFraction:      calculations.NewFraction(9, 4),
			integerToSubtract: 3,
			comparedFraction:  calculations.NewFraction(-3, 4),
		},
		{
			name:              "Correct same fraction (11÷6)-1",
			baseFraction:      calculations.NewFraction(11, 6),
			integerToSubtract: 1,
			comparedFraction:  calculations.NewFraction(5, 6),
		},
		{
			name:              "Correct same fraction (23÷12)-4",
			baseFraction:      calculations.NewFraction(23, 12),
			integerToSubtract: 4,
			comparedFraction:  calculations.NewFraction(-25, 12),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.SubtractBy(tt.integerToSubtract)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v subtract %d should equal to with %v",
					tt.baseFraction, tt.integerToSubtract, tt.comparedFraction,
				)
			}
		})
	}
}

func TestIntergerSubtractFraction(t *testing.T) {
	tests := []struct {
		name               string
		baseInteger        int
		fractionToSubtract calculations.Fraction
		comparedFraction   calculations.Fraction
	}{
		{
			name:               "Correct same fraction 4-(5÷12)",
			baseInteger:        4,
			fractionToSubtract: calculations.NewFraction(5, 12),
			comparedFraction:   calculations.NewFraction(43, 12),
		},
		{
			name:               "Correct same fraction 3-(2÷5)",
			baseInteger:        3,
			fractionToSubtract: calculations.NewFraction(2, 5),
			comparedFraction:   calculations.NewFraction(13, 5),
		},
		{
			name:               "Correct same fraction 5-(7÷8)",
			baseInteger:        5,
			fractionToSubtract: calculations.NewFraction(7, 8),
			comparedFraction:   calculations.NewFraction(33, 8),
		},
		{
			name:               "Correct same fraction 2-(3÷4)",
			baseInteger:        2,
			fractionToSubtract: calculations.NewFraction(3, 4),
			comparedFraction:   calculations.NewFraction(5, 4),
		},
		{
			name:               "Correct same fraction 9-(3÷11)",
			baseInteger:        9,
			fractionToSubtract: calculations.NewFraction(3, 11),
			comparedFraction:   calculations.NewFraction(96, 11),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fractionToSubtract.IntegerSubtractFraction(tt.baseInteger)

			if tt.fractionToSubtract != tt.comparedFraction {
				t.Errorf("%d subtract %v should equal to with %v",
					tt.baseInteger, tt.fractionToSubtract, tt.comparedFraction,
				)
			}
		})
	}
}

func TestFractionSubtractFraction(t *testing.T) {
	tests := []struct {
		name                   string
		baseFraction           calculations.Fraction
		fractionToBeSubtracted calculations.Fraction
		comparedFraction       calculations.Fraction
	}{
		{
			name:                   "Correct same fraction (4÷7)-(3÷7)",
			baseFraction:           calculations.NewFraction(4, 7),
			fractionToBeSubtracted: calculations.NewFraction(3, 7),
			comparedFraction:       calculations.NewFraction(1, 7),
		},
		{
			name:                   "Correct same fraction (7÷4)-(5÷9)",
			baseFraction:           calculations.NewFraction(7, 4),
			fractionToBeSubtracted: calculations.NewFraction(5, 9),
			comparedFraction:       calculations.NewFraction(43, 36),
		},
		{
			name:                   "Correct same fraction (5÷6)-(1÷4)",
			baseFraction:           calculations.NewFraction(5, 6),
			fractionToBeSubtracted: calculations.NewFraction(1, 4),
			comparedFraction:       calculations.NewFraction(7, 12),
		},
		{
			name:                   "Correct same fraction (9÷10)-(2÷5)",
			baseFraction:           calculations.NewFraction(9, 10),
			fractionToBeSubtracted: calculations.NewFraction(2, 5),
			comparedFraction:       calculations.NewFraction(1, 2),
		},
		{
			name:                   "Correct same fraction (7÷8)-(3÷5)",
			baseFraction:           calculations.NewFraction(7, 8),
			fractionToBeSubtracted: calculations.NewFraction(3, 5),
			comparedFraction:       calculations.NewFraction(11, 40),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.SubtractFractionBy(tt.fractionToBeSubtracted)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v subtract %v should equal to with %v",
					tt.baseFraction, tt.fractionToBeSubtracted, tt.comparedFraction,
				)
			}
		})
	}
}

func TestMultiplyFraction(t *testing.T) {
	tests := []struct {
		name              string
		baseFraction      calculations.Fraction
		integerToMultiply int
		comparedFraction  calculations.Fraction
	}{
		{
			name:              "Correct same fraction ((-9)÷(-4))×3",
			baseFraction:      calculations.NewFraction(-9, -4),
			integerToMultiply: 3,
			comparedFraction:  calculations.NewFraction(27, 4),
		},
		{
			name:              "Correct same fraction (8÷17)×2",
			baseFraction:      calculations.NewFraction(8, 17),
			integerToMultiply: 2,
			comparedFraction:  calculations.NewFraction(16, 17),
		},
		{
			name:              "Correct same fraction (3÷4)×5",
			baseFraction:      calculations.NewFraction(3, 4),
			integerToMultiply: 5,
			comparedFraction:  calculations.NewFraction(15, 4),
		},
		{
			name:              "Correct same fraction (7÷10)×3",
			baseFraction:      calculations.NewFraction(7, 10),
			integerToMultiply: 3,
			comparedFraction:  calculations.NewFraction(21, 10),
		},
		{
			name:              "Correct same fraction (9÷8)×4",
			baseFraction:      calculations.NewFraction(9, 8),
			integerToMultiply: 4,
			comparedFraction:  calculations.NewFraction(9, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.MultiplyBy(tt.integerToMultiply)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v multiply %d should equal to with %v",
					tt.baseFraction, tt.integerToMultiply, tt.comparedFraction,
				)
			}
		})
	}
}

func TestFractionMultiplyFraction(t *testing.T) {
	tests := []struct {
		name                   string
		baseFraction           calculations.Fraction
		fractionToBeMultiplied calculations.Fraction
		comparedFraction       calculations.Fraction
	}{
		{
			name:                   "Correct same fraction (5÷9)×(3÷10)",
			baseFraction:           calculations.NewFraction(5, 9),
			fractionToBeMultiplied: calculations.NewFraction(3, 10),
			comparedFraction:       calculations.NewFraction(1, 6),
		},
		{
			name:                   "Correct same fraction (4÷3)×(7÷5)",
			baseFraction:           calculations.NewFraction(4, 3),
			fractionToBeMultiplied: calculations.NewFraction(7, 5),
			comparedFraction:       calculations.NewFraction(28, 15),
		},
		{
			name:                   "Correct same fraction (2÷3)×(4÷5)",
			baseFraction:           calculations.NewFraction(2, 3),
			fractionToBeMultiplied: calculations.NewFraction(4, 5),
			comparedFraction:       calculations.NewFraction(8, 15),
		},
		{
			name:                   "Correct same fraction (7÷8)×(3÷4)",
			baseFraction:           calculations.NewFraction(7, 8),
			fractionToBeMultiplied: calculations.NewFraction(3, 4),
			comparedFraction:       calculations.NewFraction(21, 32),
		},
		{
			name:                   "Correct same fraction (5÷6)×(9÷11)",
			baseFraction:           calculations.NewFraction(5, 6),
			fractionToBeMultiplied: calculations.NewFraction(9, 11),
			comparedFraction:       calculations.NewFraction(15, 22),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.MultiplyFractionBy(tt.fractionToBeMultiplied)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v multiply %v should equal to with %v",
					tt.baseFraction, tt.fractionToBeMultiplied, tt.comparedFraction,
				)
			}
		})
	}
}

func TestDivideFraction(t *testing.T) {
	tests := []struct {
		name             string
		baseFraction     calculations.Fraction
		integerToDivide  int
		comparedFraction calculations.Fraction
	}{
		{
			name:             "Correct same fraction (4÷9)÷2",
			baseFraction:     calculations.NewFraction(4, 9),
			integerToDivide:  2,
			comparedFraction: calculations.NewFraction(4, 18),
		},
		{
			name:             "Correct same fraction (5÷6)÷3",
			baseFraction:     calculations.NewFraction(5, 6),
			integerToDivide:  3,
			comparedFraction: calculations.NewFraction(5, 18),
		},
		{
			name:             "Correct same fraction (7÷10)÷3",
			baseFraction:     calculations.NewFraction(7, 10),
			integerToDivide:  3,
			comparedFraction: calculations.NewFraction(7, 30),
		},
		{
			name:             "Correct same fraction ((-9)÷8)÷4",
			baseFraction:     calculations.NewFraction(-9, 8),
			integerToDivide:  4,
			comparedFraction: calculations.NewFraction(-9, 32),
		},
		{
			name:             "Correct same fraction (1÷2)÷5",
			baseFraction:     calculations.NewFraction(1, 2),
			integerToDivide:  5,
			comparedFraction: calculations.NewFraction(1, 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.DivideBy(tt.integerToDivide)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v divide %d should equal to with %v",
					tt.baseFraction, tt.integerToDivide, tt.comparedFraction,
				)
			}
		})
	}
}

func TestIntergerDivideFraction(t *testing.T) {
	tests := []struct {
		name             string
		baseInteger      int
		fractionToDivide calculations.Fraction
		comparedFraction calculations.Fraction
	}{
		{
			name:             "Correct same fraction 8÷(4÷7)",
			baseInteger:      8,
			fractionToDivide: calculations.NewFraction(4, 7),
			comparedFraction: calculations.NewFraction(14, 1),
		},
		{
			name:             "Correct same fraction 9÷(4÷7)",
			baseInteger:      9,
			fractionToDivide: calculations.NewFraction(4, 7),
			comparedFraction: calculations.NewFraction(63, 4),
		},
		{
			name:             "Correct same fraction 2÷(3÷4)",
			baseInteger:      2,
			fractionToDivide: calculations.NewFraction(3, 4),
			comparedFraction: calculations.NewFraction(8, 3),
		},
		{
			name:             "Correct same fraction 5÷(7÷10)",
			baseInteger:      5,
			fractionToDivide: calculations.NewFraction(7, 10),
			comparedFraction: calculations.NewFraction(50, 7),
		},
		{
			name:             "Correct same fraction 4÷(9÷8)",
			baseInteger:      4,
			fractionToDivide: calculations.NewFraction(9, 8),
			comparedFraction: calculations.NewFraction(32, 9),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fractionToDivide.IntegerDivideFraction(tt.baseInteger)

			if tt.fractionToDivide != tt.comparedFraction {
				t.Errorf("%d divide %v should equal to with %v",
					tt.baseInteger, tt.fractionToDivide, tt.comparedFraction,
				)
			}
		})
	}
}

func TestFractionDivideFraction(t *testing.T) {
	tests := []struct {
		name                string
		baseFraction        calculations.Fraction
		fractionToBeDivided calculations.Fraction
		comparedFraction    calculations.Fraction
	}{
		{
			name:                "Correct same fraction (14÷19)÷(7÷4)",
			baseFraction:        calculations.NewFraction(14, 19),
			fractionToBeDivided: calculations.NewFraction(7, 4),
			comparedFraction:    calculations.NewFraction(8, 19),
		},
		{
			name:                "Correct same fraction (5÷3)÷(2÷7)",
			baseFraction:        calculations.NewFraction(5, 3),
			fractionToBeDivided: calculations.NewFraction(2, 7),
			comparedFraction:    calculations.NewFraction(35, 6),
		},
		{
			name:                "Correct same fraction (3÷4)÷(2÷5)",
			baseFraction:        calculations.NewFraction(3, 4),
			fractionToBeDivided: calculations.NewFraction(2, 5),
			comparedFraction:    calculations.NewFraction(15, 8),
		},
		{
			name:                "Correct same fraction (7÷10)÷(3÷8)",
			baseFraction:        calculations.NewFraction(7, 10),
			fractionToBeDivided: calculations.NewFraction(3, 8),
			comparedFraction:    calculations.NewFraction(28, 15),
		},
		{
			name:                "Correct same fraction ((-5)÷6)÷(4÷9)",
			baseFraction:        calculations.NewFraction(-5, 6),
			fractionToBeDivided: calculations.NewFraction(4, 9),
			comparedFraction:    calculations.NewFraction(-15, 8),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.baseFraction.DivideFractionBy(tt.fractionToBeDivided)

			if tt.baseFraction != tt.comparedFraction {
				t.Errorf("%v divide %v should equal to with %v",
					tt.baseFraction, tt.fractionToBeDivided, tt.comparedFraction,
				)
			}
		})
	}
}

func TestIsFractionEmpty(t *testing.T) {
	fraction := calculations.Fraction{}

	if !fraction.IsEmpty() {
		t.Errorf("IsEmpty() func is incorrect")
	}
}
