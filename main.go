package main

import (
	"A-MATH/game/actions"
	"A-MATH/game/components"
	"A-MATH/game/players"
	"fmt"
)

func main() {
	// c, _ := game.NewChip("x")
	// d, _ := game.NewChip("/")
	// fmt.Println(c == d)

	//--------------------------------------------------------------------------

	// a, _ := game.NewChip("1")
	// fmt.Println(a)

	//--------------------------------------------------------------------------

	// b, e := game.NewSquare("H8")
	// if e != nil {
	// 	fmt.Println(e.Error())
	// }
	// fmt.Println(b)

	//--------------------------------------------------------------------------

	// a, _ := game.NewChip("1")
	// f := game.NewBoard()

	// pos := [2]int{4, 1}

	// f.Add(pos, a)
	// fmt.Println(f)

	// f.Remove(pos)
	// fmt.Println(f)

	//--------------------------------------------------------------------------

	// c, _ := components.NewChip("10")

	// g := components.NewBag()

	// g.DecreaseChip(c)
	// g.DecreaseChip(c)
	// fmt.Println(g)
	// fmt.Println("--------------------------------------------------------------")
	// g.IncreaseChip(c)
	// fmt.Println(g)

	//--------------------------------------------------------------------------

	// a, _ := game.NewChip("1")
	// b, _ := game.NewChip("19")
	// c, _ := game.NewChip("=")
	// d, _ := game.NewChip("=")
	// e, _ := game.NewChip("7")
	// f, _ := game.NewChip("+-")
	// g, _ := game.NewChip("+")
	// h, _ := game.NewChip("0")

	// r := game.NewRack(1)
	// fmt.Println(r)

	// r.Add(a)
	// r.Add(b)
	// r.Add(c)
	// r.Add(d)
	// r.Add(e)
	// r.Add(f)
	// r.Add(g)
	// r.Add(h)
	// fmt.Println(r)
	// fmt.Println(r.IsFull())

	// r.Remove(a)
	// fmt.Println(r)

	// r.Remove(e)
	// fmt.Println(r)

	//--------------------------------------------------------------------------

	// c, _ := components.NewChip("*")
	// e := components.NewChipCollector(c, 4)
	// fmt.Println(e)

	// e.DecreaseChip()
	// fmt.Println(e)

	// e.IncreaseChip()
	// fmt.Println(e)

	//--------------------------------------------------------------------------

	p1 := players.NewPlayer(1, "p1")
	p2 := players.NewPlayer(2, "p2")
	ng := actions.NewGame([2]players.Player{p1, p2})
	// ng.FullyDraw(&ng.GamePlayers[0])
	// ng.FullyDraw(&ng.GamePlayers[1])
	a, _ := components.NewChip("1")
	b, _ := components.NewChip("18")
	c, _ := components.NewChip("=")
	d, _ := components.NewChip("=")
	e, _ := components.NewChip("7")
	f, _ := components.NewChip("+-")
	g, _ := components.NewChip("+")
	h, _ := components.NewChip("0")
	ng.DrawSpecificChip(&ng.GamePlayers[0], a)
	ng.DrawSpecificChip(&ng.GamePlayers[0], b)
	ng.DrawSpecificChip(&ng.GamePlayers[0], c)
	ng.DrawSpecificChip(&ng.GamePlayers[0], d)
	ng.DrawSpecificChip(&ng.GamePlayers[0], e)
	ng.DrawSpecificChip(&ng.GamePlayers[0], f)
	ng.DrawSpecificChip(&ng.GamePlayers[0], g)
	ng.DrawSpecificChip(&ng.GamePlayers[0], h)
	fmt.Println(ng)

	ng.Exchange(&ng.GamePlayers[0], []components.Chip{b, c, g})
	fmt.Println(ng)
}
