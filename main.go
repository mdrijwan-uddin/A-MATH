package main

import (
	"A-MATH/game/actions"
	"A-MATH/game/calculations"
	"A-MATH/game/components"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"A-MATH/game/players"
	"fmt"
)

func main() {

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
		ng.Board.Add(ch.Coordinate, ch.Chip)
		ng.Bag.DecreaseChip(ch.Chip)
	}
	fmt.Println(ng)

	var chipForPlacing []models.ChipForPlacing

	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{1, 5}, components.NewChip("5")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{3, 5}, components.NewChip("3")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{5, 5}, components.NewChip("12")))
	chipForPlacing = append(chipForPlacing, models.NewChipForPlacing([2]int{7, 5}, components.NewChip("3")))

	mapping, e := mappings.Management(*ng.Board, chipForPlacing)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("---------------------------------------------------------")
	fmt.Println(mapping)
	fmt.Println("---------------------------------------------------------")

	result, err := calculations.Management(mapping)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
