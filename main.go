package main

import (
	"A-MATH/models/game"
	"fmt"
)

func main() {
	// c, _ := game.NewChip("x")
	// d, _ := game.NewChip("/")
	// fmt.Println(c == d)

	// a, _ := game.NewChip("1")
	// fmt.Println(a)

	// b, e := game.NewSquare("H8")
	// if e != nil {
	// 	fmt.Println(e.Error())
	// }
	// fmt.Println(b)

	f := game.NewBoard()
	fmt.Println(f)

	// g := game.NewBag()
	// fmt.Println(g)

	// r := game.NewRack(1)
	// fmt.Println(r)

}
