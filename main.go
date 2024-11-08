package main

import (
	"A-MATH/models/game"
	"fmt"
)

func main() {
	// c := game.NewChip("x")
	// d := game.NewChip("/")
	// fmt.Println(c == d)

	// a := game.NewChip("1")
	// fmt.Println(a)

	b, e := game.NewSquare("A2")
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(b)
}
