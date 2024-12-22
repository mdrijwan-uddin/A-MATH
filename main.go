package main

import (
	"A-MATH/game/actions"
	"A-MATH/game/calculations"
	"A-MATH/game/components"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"A-MATH/game/players"
	"A-MATH/game/rules"
	"fmt"
)

func main() {

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

	// p1 := players.NewPlayer(1, "p1")
	// p2 := players.NewPlayer(2, "p2")
	// ng := actions.NewGame([2]players.Player{p1, p2})
	// a, _ := components.NewChip("1")
	// b, _ := components.NewChip("18")
	// c, _ := components.NewChip("=")
	// d, _ := components.NewChip("=")
	// e, _ := components.NewChip("7")
	// f, _ := components.NewChip("+-")
	// g, _ := components.NewChip("+")
	// h, _ := components.NewChip("0")
	// ng.DrawSpecificChip(&ng.GamePlayers[0], a)
	// ng.DrawSpecificChip(&ng.GamePlayers[0], b)
	// ng.DrawSpecificChip(&ng.GamePlayers[0], c)
	// ng.DrawSpecificChip(&ng.GamePlayers[0], d)
	// ng.DrawSpecificChip(&ng.GamePlayers[0], e)
	// ng.DrawSpecificChip(&ng.GamePlayers[0], f)
	// ng.DrawSpecificChip(&ng.GamePlayers[0], g)
	// ng.DrawSpecificChip(&ng.GamePlayers[0], h)
	// fmt.Println(ng)

	// ng.Exchange(&ng.GamePlayers[0], []components.Chip{b, c, g})
	// fmt.Println(ng)

	//--------------------------------------------------------------------------

	p1 := players.NewPlayer(1, "p1")
	p2 := players.NewPlayer(2, "p2")
	ng := actions.NewGame([2]players.Player{p1, p2})

	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{1, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 8}, components.NewChip("6")))
	chips = append(chips, models.NewChipForPlacing([2]int{3, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 8}, components.NewChip("0")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 8}, components.NewChip("/")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("13")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("16")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 8}, components.NewChip("/")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 8}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{13, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{14, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{15, 8}, components.NewChip("5")))

	chips = append(chips, models.NewChipForPlacing([2]int{2, 3}, components.NewChip("3")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 4}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 5}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 6}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 7}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 9}, components.NewChip("2")))

	chips = append(chips, models.NewChipForPlacing([2]int{4, 4}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 5}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 6}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 7}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 9}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 10}, components.NewChip("6")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 11}, components.NewChip("5")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 12}, components.NewChip("9")))

	chips = append(chips, models.NewChipForPlacing([2]int{6, 4}, components.NewChip("3")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 5}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 6}, components.NewChip("10")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 7}, components.NewChip("=")))

	chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("20")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 11}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 12}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 13}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 14}, components.NewChip("/")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 15}, components.NewChip("6")))

	chips = append(chips, models.NewChipForPlacing([2]int{9, 12}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 12}, components.NewChip("11")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 12}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 12}, components.NewChip("9")))

	chips = append(chips, models.NewChipForPlacing([2]int{7, 14}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{9, 14}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{10, 14}, components.NewChip("*")))
	chips = append(chips, models.NewChipForPlacing([2]int{11, 14}, components.NewChip("4")))
	chips = append(chips, models.NewChipForPlacing([2]int{12, 14}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{13, 14}, components.NewChip("2")))
	chips = append(chips, models.NewChipForPlacing([2]int{14, 14}, components.NewChip("8")))

	for _, ch := range chips {
		ng.Board.Add(ch.Position, ch.Chip)
		ng.Bag.DecreaseChip(ch.Chip)
	}
	fmt.Println(ng)

	// // for _, ch := range chipForPlacing {
	// // 	fmt.Println(ng.Board.GetSquare(ch.Position))
	// // }

	// // var position [][2]int
	// // for _, chip := range chipForPlacing {
	// // 	position = append(position, chip.Position)
	// // }
	// // fmt.Println(rules.IsChipPlaceOnCenterSquare(position))
	// // fmt.Println(rules.IsChipPlaceOnVerticalOrHorizontal(position))

	var chipForPlacing []models.ChipForPlacing

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{1, 5}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{3, 5}, components.NewChip("3")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{5, 5}, components.NewChip("12")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{7, 5}, components.NewChip("3")))

	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 7}, components.NewChip("9")))
	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 9}, components.NewChip("3")))
	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 10}, components.NewChip("=")))
	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 11}, components.NewChip("17")))
	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{11, 13}, components.NewChip("1")))

	var coordinates [][2]int
	for _, chips := range chipForPlacing {
		coordinates = append(coordinates, chips.Position)
	}

	isVertical, _ := rules.IsChipPlaceOnVerticalOrHorizontal(coordinates)
	isStraightLine, _ := rules.IsChipPlacingOnStraightLineOrSeparated(*ng.Board, coordinates, isVertical)

	connector := mappings.CrossConnector(*ng.Board, coordinates, isVertical)

	mapping := mappings.StraightMapping(*ng.Board, chipForPlacing, connector, isVertical, isStraightLine)

	fmt.Println("---------------------------------------------------------")
	fmt.Println(mapping)
	fmt.Println("---------------------------------------------------------")

	result, err := calculations.Management(mapping)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	// calculations.ChipCalculationSeperation(mapping[0])

	// fmt.Println("Score:", calculations.ScoreCalculation(mapping[0]))

	// fmt.Println("---------------------------------------------------------")
	// fmt.Println(isVertical)
	// fmt.Println(isHorizontal)
	// fmt.Println(isStraightLine)
	// fmt.Println(isSeperated)

	//--------------------------------------------------------------------------
	// p1 := players.NewPlayer(1, "p1")
	// p2 := players.NewPlayer(2, "p2")
	// ng := actions.NewGame([2]players.Player{p1, p2})

	// var chips []models.ChipForPlacing

	// chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("3")))
	// chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("*")))
	// chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("1")))
	// chips = append(chips, models.NewChipForPlacing([2]int{9, 8}, components.NewChip("=")))
	// chips = append(chips, models.NewChipForPlacing([2]int{10, 8}, components.NewChip("3")))

	// chips = append(chips, models.NewChipForPlacing([2]int{8, 6}, components.NewChip("6")))
	// chips = append(chips, models.NewChipForPlacing([2]int{8, 7}, components.NewChip("=")))
	// chips = append(chips, models.NewChipForPlacing([2]int{8, 9}, components.NewChip("+")))
	// chips = append(chips, models.NewChipForPlacing([2]int{8, 10}, components.NewChip("5")))

	// for _, ch := range chips {
	// 	ng.Board.Add(ch.Position, ch.Chip)
	// 	ng.Bag.DecreaseChip(ch.Chip)
	// }
	// fmt.Println(ng)

	// var chipForPlacing []models.ChipForPlacing

	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 2}, components.NewChip("4")))
	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 3}, components.NewChip("+")))
	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 4}, components.NewChip("2")))
	// chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{8, 5}, components.NewChip("=")))

	// var coordinates [][2]int
	// for _, chips := range chipForPlacing {
	// 	coordinates = append(coordinates, chips.Position)
	// }

	// isVertical, _ := rules.IsChipPlaceOnVerticalOrHorizontal(coordinates)

	// connector := mappings.StraightConnector(*ng.Board, coordinates, isVertical)

	// mapping := mappings.StraightMapping(*ng.Board, chipForPlacing, connector, isVertical)

	// fmt.Println(mapping)

	//--------------------------------------------------------------------------

}
