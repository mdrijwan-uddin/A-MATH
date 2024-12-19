package calculations

// Fraction represents a Fraction with numerator and denominator
type Fraction struct {
	Numerator   int
	Denominator int
}

func NewFraction(Numerator, Denominator int) Fraction {
	f := Fraction{Numerator, Denominator}
	f.simplify()
	return f
}

// Adds a fraction and an integer
func (f *Fraction) AddBy(i int) {
	f.Numerator = f.Numerator + i*f.Denominator
	f.simplify()
}

// Adds a fraction by a fraction
func (f *Fraction) AddFractionBy(s Fraction) {
	if f.Denominator == s.Denominator {
		f.Numerator = f.Numerator + s.Numerator
	} else {
		f.Numerator = (f.Numerator * s.Denominator) + (s.Numerator * f.Denominator)
		f.Denominator = f.Denominator * s.Denominator
	}
	f.simplify()
}

// Subtracts an integer from a fraction
func (f *Fraction) SubtractBy(i int) {
	f.Numerator = f.Numerator - i*f.Denominator
	f.simplify()
}

// An integer Subtracts fraction
func (f *Fraction) IntegerSubtractFraction(i int) {
	f.Numerator = i*f.Denominator - f.Numerator
	f.simplify()
}

// Subtracts a fraction by a fraction
func (f *Fraction) SubtractFractionBy(s Fraction) {
	if f.Denominator == s.Denominator {
		f.Numerator = f.Numerator - s.Numerator
	} else {
		f.Numerator = (f.Numerator * s.Denominator) - (s.Numerator * f.Denominator)
		f.Denominator = f.Denominator * s.Denominator
	}
	f.simplify()
}

// Multiplies a fraction by an integer
func (f *Fraction) MultiplyBy(i int) {
	f.Numerator = f.Numerator * i
	f.simplify()
}

// Multiplies a fraction by a fraction
func (f *Fraction) MultiplyFractionBy(s Fraction) {
	f.Numerator = f.Numerator * s.Numerator
	f.Denominator = f.Denominator * s.Denominator
	f.simplify()
}

// Divides a fraction by an integer
func (f *Fraction) DivideBy(i int) {
	f.Denominator = f.Denominator * i
	f.simplify()
}

// Divides a fraction by a fraction
func (f *Fraction) DivideFractionBy(s Fraction) {
	f.Numerator = f.Numerator * s.Denominator
	f.Denominator = f.Denominator * s.Numerator
	f.simplify()
}

// An integer Divides fraction
func (f *Fraction) IntegerDivideFraction(i int) {
	temp := f.Numerator
	f.Numerator = i * f.Denominator
	f.Denominator = temp
	f.simplify()
}

// Check is Empty or not
func (f *Fraction) IsEmpty() bool {
	return f.Numerator == 0 && f.Denominator == 0
}

// gcd finds the greatest common divisor of two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// simplify reduces a fraction to its simplest form
func (f *Fraction) simplify() {
	divisor := gcd(abs(f.Numerator), abs(f.Denominator))
	f.Numerator /= divisor
	f.Denominator /= divisor

	// Ensure the denominator is positive
	if f.Denominator < 0 {
		f.Numerator = -f.Numerator
		f.Denominator = -f.Denominator
	}
}
