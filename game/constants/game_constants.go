package constants

type values string
type chipTypes string
type squareTypes string
type directions string

const (
	A values = "A"
	B values = "B"
	C values = "C"
	D values = "D"
	E values = "E"
	F values = "F"
	G values = "G"
	H values = "H"
	I values = "I"
	J values = "J"
	K values = "K"
	L values = "L"
	M values = "M"
	N values = "N"
	O values = "O"

	Zero         values = "0"
	One          values = "1"
	Two          values = "2"
	Three        values = "3"
	Four         values = "4"
	Five         values = "5"
	Six          values = "6"
	Seven        values = "7"
	Eight        values = "8"
	Nine         values = "9"
	Ten          values = "10"
	Eleven       values = "11"
	Twelve       values = "12"
	Thirteen     values = "13"
	Fourteen     values = "14"
	Fifteen      values = "15"
	Sixteen      values = "16"
	Seventeen    values = "17"
	Eighteen     values = "18"
	Nineteen     values = "19"
	Twenty       values = "20"
	Addition     values = "+"
	Subtraction  values = "-"
	Multiply     values = "*"
	Division     values = "/"
	Add_sub      values = "+-"
	Multi_divide values = "*/"
	Equal        values = "="
	Blank        values = "~"

	OneDigitNumberType chipTypes = "1 Digit Number"
	TwoDigitNumberType chipTypes = "2 Digit Number"
	OperatorType       chipTypes = "Normal Operator"
	AlterOperatorType  chipTypes = "Alternative Operator"
	EqualType          chipTypes = "Equal Operator"
	BlankType          chipTypes = "Blank"

	NormalSquare squareTypes = "Normal Square"
	OrangeSquare squareTypes = "Orange Square"
	BlueSquare   squareTypes = "Blue Square"
	YellowSquare squareTypes = "Yellow Square"
	RedSquare    squareTypes = "Red Square"
	CenterSquare squareTypes = "Center Square"

	MaxChipInRack = 8
	BoardRange    = 15

	Left  directions = "left"
	Right directions = "right"
	Up    directions = "up"
	Down  directions = "down"

	RandomLetter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)
