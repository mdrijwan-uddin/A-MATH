package game

type Square struct {
	Position   [2]int
	SquareType string
	IsPlaced   bool
}

func NewSquare(pos string) Square {
	// position := setSquarePosition(pos)
	return Square{[2]int{0, 0}, pos, false}
}

func SetSquarePosition(pos string) {
	//
}
