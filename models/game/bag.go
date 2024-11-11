package game

import (
	"A-MATH/utils"
	"fmt"
)

type ChipCollector struct {
	chip     Chip
	chipLeft int
}

type Bag struct {
	ChipCollectors []ChipCollector
	totalChipLeft  int
}

func NewChipCollector(chip Chip, chipLeft int) ChipCollector {
	// fmt.Println(chip)
	return ChipCollector{chip, chipLeft}
}

func NewBag() Bag {
	var ChipCollectors []ChipCollector
	var totalChipLeft int

	Chips := utils.ChipSet
	totalChipSet := [29]int{
		5, 6, 6, 5, 5,
		4, 4, 4, 4, 4,
		2, 1, 2, 1, 1,
		1, 1, 1, 1, 1,
		1, 4, 4, 5, 4,
		4, 4, 11, 4}

	for i := range len(Chips) {
		fmt.Println(Chips)
		c, e := NewChip(Chips[i])
		if e != nil {
			fmt.Println(e)
		}
		totalChipLeft += totalChipSet[i]
		n := NewChipCollector(c, totalChipSet[i])
		ChipCollectors = append(ChipCollectors, n)

	}

	return Bag{ChipCollectors, totalChipLeft}
}

// ChipCollector("0", 5), ChipCollector("1", 6), ChipCollector("2", 6),
// ChipCollector("3", 5), ChipCollector("4", 5), ChipCollector("5", 4),
// ChipCollector("6", 4), ChipCollector("7", 4), ChipCollector("8", 4),
// ChipCollector("9", 4), ChipCollector("10", 2), ChipCollector("11", 1),
// ChipCollector("12", 2), ChipCollector("13", 1), ChipCollector("14", 1),
// ChipCollector("15", 1), ChipCollector("16", 1), ChipCollector("17", 1),
// ChipCollector("18", 1), ChipCollector("19", 1), ChipCollector("20", 1),
// ChipCollector("+", 4), ChipCollector("-", 4), ChipCollector("+/-", 5),
// ChipCollector("x", 4), ChipCollector("%", 4), ChipCollector("x/%", 4),
// ChipCollector("=", 11), ChipCollector("Blank", 4),
