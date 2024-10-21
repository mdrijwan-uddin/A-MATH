package model

type values string
type score int
type chipType string

type squareType string

const (
	ZeroPoint score = iota
	OnePoint
	TwoPoint
	ThreePoint
	FourPoint
	FivePoint
	SixPoint
	SevenPoint

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
	Multiply     values = "x"
	Division     values = "/"
	Add_sub      values = "+-"
	Multi_divide values = "x/"
	Equal        values = "="
	Blank        values = ""

	OneDigitNumberType chipType = "1 Digit Number"
	TwoDigitNumberType chipType = "2 Digit Number"
	OperatorType       chipType = "Operator"
	EqualType          chipType = "Equal"
	BlankType          chipType = "Blank"

	NormalSquare squareType = "Normal Square"
	CenterSquare squareType = "Center Square"
	OrangeSquare squareType = "Orange Square"
	BlueSquare   squareType = "Blue Square"
	YellowSquare squareType = "Yellow Square"
	RedSquare    squareType = "Red Square"
)
