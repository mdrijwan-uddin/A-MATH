package calculations_test

import (
	"A-MATH/game/calculations"
	"testing"
)

func TestNewFraction(t *testing.T) {
	numerator := 1
	denominator := 5

	expectedNumerator := 1
	expectedDenominator := 5

	fraction := calculations.NewFraction(numerator, denominator)

	if fraction.Numerator != expectedNumerator || fraction.Denominator != expectedDenominator {
		t.Errorf("NewFraction(%d/%d) will be %d/%d But result is %d/%d",
			numerator, denominator, expectedNumerator, expectedDenominator, fraction.Numerator, fraction.Denominator)
	}
}

func TestAddFraction(t *testing.T) {
	numerator := 6
	denominator := 7
	integer := 4

	expectedNumerator := 34
	expectedDenominator := 7

	fraction := calculations.NewFraction(numerator, denominator)
	fraction.AddBy(integer)

	if fraction.Numerator != expectedNumerator || fraction.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).AddBy(%d) will be %d/%d But result is %d/%d",
			numerator, denominator, integer, expectedNumerator, expectedDenominator, fraction.Numerator, fraction.Denominator)
	}
}

func TestFractionAddFraction1(t *testing.T) {
	fNumerator := 4
	fDenominator := 7
	sNumerator := 2
	sDenominator := 7

	expectedNumerator := 6
	expectedDenominator := 7

	f := calculations.NewFraction(fNumerator, fDenominator)
	s := calculations.NewFraction(sNumerator, sDenominator)
	f.AddFractionBy(s)

	if f.Numerator != expectedNumerator || f.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).AddFractionBy(%d/%d) will be %d/%d But result is %d/%d",
			fNumerator, fDenominator, sNumerator, sDenominator, expectedNumerator, expectedDenominator, f.Numerator, f.Denominator)
	}
}

func TestFractionAddFraction2(t *testing.T) {
	fNumerator := 3
	fDenominator := 10
	sNumerator := 5
	sDenominator := 4

	expectedNumerator := 31
	expectedDenominator := 20

	f := calculations.NewFraction(fNumerator, fDenominator)
	s := calculations.NewFraction(sNumerator, sDenominator)
	f.AddFractionBy(s)

	if f.Numerator != expectedNumerator || f.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).AddFractionBy(%d/%d) will be %d/%d But result is %d/%d",
			fNumerator, fDenominator, sNumerator, sDenominator, expectedNumerator, expectedDenominator, f.Numerator, f.Denominator)
	}
}

func TestSubtractFraction(t *testing.T) {
	numerator := 17
	denominator := 4
	integer := 3

	expectedNumerator := 5
	expectedDenominator := 4

	fraction := calculations.NewFraction(numerator, denominator)
	fraction.SubtractBy(integer)

	if fraction.Numerator != expectedNumerator || fraction.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).SubtractBy(%d) will be %d/%d But result is %d/%d",
			numerator, denominator, integer, expectedNumerator, expectedDenominator, fraction.Numerator, fraction.Denominator)
	}
}

func TestFractionSubtractFraction1(t *testing.T) {
	fNumerator := 9
	fDenominator := 10
	sNumerator := 4
	sDenominator := 10

	expectedNumerator := 1
	expectedDenominator := 2

	f := calculations.NewFraction(fNumerator, fDenominator)
	s := calculations.NewFraction(sNumerator, sDenominator)
	f.SubtractFractionBy(s)

	if f.Numerator != expectedNumerator || f.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).SubtractFractionBy(%d/%d) will be %d/%d But result is %d/%d",
			fNumerator, fDenominator, sNumerator, sDenominator, expectedNumerator, expectedDenominator, f.Numerator, f.Denominator)
	}
}

func TestFractionSubtractFraction2(t *testing.T) {
	fNumerator := 7
	fDenominator := 4
	sNumerator := 5
	sDenominator := 9

	expectedNumerator := 43
	expectedDenominator := 36

	f := calculations.NewFraction(fNumerator, fDenominator)
	s := calculations.NewFraction(sNumerator, sDenominator)
	f.SubtractFractionBy(s)

	if f.Numerator != expectedNumerator || f.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).SubtractFractionBy(%d/%d) will be %d/%d But result is %d/%d",
			fNumerator, fDenominator, sNumerator, sDenominator, expectedNumerator, expectedDenominator, f.Numerator, f.Denominator)
	}
}

func TestIntergerSubtractFraction(t *testing.T) {
	numerator := 5
	denominator := 12
	integer := 4

	expectedNumerator := 43
	expectedDenominator := 12

	fraction := calculations.NewFraction(numerator, denominator)
	fraction.IntegerSubtractFraction(integer)

	if fraction.Numerator != expectedNumerator || fraction.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).IntegerSubtractFraction(%d) will be %d/%d But result is %d/%d",
			numerator, denominator, integer, expectedNumerator, expectedDenominator, fraction.Numerator, fraction.Denominator)
	}
}

func TestMultiplyFraction(t *testing.T) {
	numerator := -9
	denominator := -14
	integer := 7

	expectedNumerator := 9
	expectedDenominator := 2

	fraction := calculations.NewFraction(numerator, denominator)
	fraction.MultiplyBy(integer)

	if fraction.Numerator != expectedNumerator || fraction.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).MultiplyBy(%d) will be %d/%d But result is %d/%d",
			numerator, denominator, integer, expectedNumerator, expectedDenominator, fraction.Numerator, fraction.Denominator)
	}
}

func TestFractionMultiplyFraction(t *testing.T) {
	fNumerator := 5
	fDenominator := 9
	sNumerator := 3
	sDenominator := 10

	expectedNumerator := 1
	expectedDenominator := 6

	f := calculations.NewFraction(fNumerator, fDenominator)
	s := calculations.NewFraction(sNumerator, sDenominator)
	f.MultiplyFractionBy(s)

	if f.Numerator != expectedNumerator || f.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).MultiplyFractionBy(%d/%d) will be %d/%d But result is %d/%d",
			fNumerator, fDenominator, sNumerator, sDenominator, expectedNumerator, expectedDenominator, f.Numerator, f.Denominator)
	}
}

func TestDivideFraction(t *testing.T) {
	numerator := 4
	denominator := 9
	integer := 2

	expectedNumerator := 2
	expectedDenominator := 9

	fraction := calculations.NewFraction(numerator, denominator)
	fraction.DivideBy(integer)

	if fraction.Numerator != expectedNumerator || fraction.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).DivideBy(%d) will be %d/%d But result is %d/%d",
			numerator, denominator, integer, expectedNumerator, expectedDenominator, fraction.Numerator, fraction.Denominator)
	}
}

func TestIntergerDivideFraction(t *testing.T) {
	numerator := 4
	denominator := 7
	integer := 8

	expectedNumerator := 14
	expectedDenominator := 1

	fraction := calculations.NewFraction(numerator, denominator)
	fraction.IntegerDivideFraction(integer)

	if fraction.Numerator != expectedNumerator || fraction.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).IntegerDivideFraction(%d) will be %d/%d But result is %d/%d",
			numerator, denominator, integer, expectedNumerator, expectedDenominator, fraction.Numerator, fraction.Denominator)
	}
}

func TestFractionDivideFraction(t *testing.T) {
	fNumerator := -14
	fDenominator := 19
	sNumerator := 7
	sDenominator := 4

	expectedNumerator := -8
	expectedDenominator := 19

	f := calculations.NewFraction(fNumerator, fDenominator)
	s := calculations.NewFraction(sNumerator, sDenominator)
	f.DivideFractionBy(s)

	if f.Numerator != expectedNumerator || f.Denominator != expectedDenominator {
		t.Errorf("Fraction(%d/%d).MultiplyFractionBy(%d/%d) will be %d/%d But result is %d/%d",
			fNumerator, fDenominator, sNumerator, sDenominator, expectedNumerator, expectedDenominator, f.Numerator, f.Denominator)
	}
}

func TestIsFractionEmpty(t *testing.T) {
	fraction := calculations.Fraction{}

	if !fraction.IsEmpty() {
		t.Errorf("IsEmpty() func is incorrect")
	}
}
