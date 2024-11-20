package components

import (
	"A-MATH/err"
	"A-MATH/game/constants"
	"A-MATH/game/utils"

	"strconv"
	"strings"
)

var totalSet = [29]int{
	5, 6, 6, 5, 5,
	4, 4, 4, 4, 4,
	2, 1, 2, 1, 1,
	1, 1, 1, 1, 1,
	1, 4, 4, 4, 4,
	5, 4, 11, 4}

var chipSet = utils.ChipSet

// ChipCollector("0", 5),  ChipCollector("1", 6),  ChipCollector("2", 6),  ChipCollector("3", 5),  ChipCollector("4", 5),
// ChipCollector("5", 4),  ChipCollector("6", 4),  ChipCollector("7", 4),  ChipCollector("8", 4),  ChipCollector("9", 4),
// ChipCollector("10", 2), ChipCollector("11", 1), ChipCollector("12", 2), ChipCollector("13", 1), ChipCollector("14", 1),
// ChipCollector("15", 1), ChipCollector("16", 1), ChipCollector("17", 1), ChipCollector("18", 1), ChipCollector("19", 1),
// ChipCollector("20", 1), ChipCollector("+", 4),  ChipCollector("-", 4),   ChipCollector("x", 4), ChipCollector("%", 4),
// ChipCollector("+/-", 5), ChipCollector("x/%", 4), ChipCollector("=", 11), ChipCollector("Blank", 4)

type Bag struct {
	ChipCollectors []chipCollectors
	TotalChipLeft  int
	MaxChip        int
}

func NewBag() Bag {
	var chipCollectors []chipCollectors
	var totalChip int

	for i := range len(chipSet) {
		c, _ := NewChip(chipSet[i])
		totalChip += totalSet[i]
		n := NewChipCollector(c, totalSet[i])
		chipCollectors = append(chipCollectors, n)
	}

	return Bag{chipCollectors, totalChip, totalChip}
}

func (b *Bag) getIndex(c chip) (int, error) {
	for i, cs := range chipSet {
		if cs == c.Value {
			return i, nil
		}
	}
	return -1, err.New(int(constants.BadRequest), string(constants.InvalidInputForChip))
}

func (b *Bag) DecreaseChip(c chip) error {
	index, e := b.getIndex(c)
	if e != nil {
		return e
	}

	totalChipInBag := b.ChipCollectors[index].Total
	if 0 < totalChipInBag {
		b.ChipCollectors[index].DecreaseChip()
		b.TotalChipLeft--
	}
	return nil
}

func (b *Bag) IncreaseChip(c chip) error {
	index, e := b.getIndex(c)
	if e != nil {
		return e
	}

	totalChipInBag := b.ChipCollectors[index].Total
	maximumChipInBag := b.ChipCollectors[index].MaxChip
	if 0 < totalChipInBag && totalChipInBag < maximumChipInBag {
		b.ChipCollectors[index].IncreaseChip()
		b.TotalChipLeft++
	}
	return nil
}

func (b Bag) String() string {
	var sb strings.Builder

	for i := range len(totalSet) {
		sb.WriteString("[")
		sb.WriteString(b.ChipCollectors[i].Chips.Value)
		sb.WriteString("]")
		sb.WriteString(strconv.Itoa(b.ChipCollectors[i].Total))
		sb.WriteString("/")
		sb.WriteString(strconv.Itoa(b.ChipCollectors[i].MaxChip))
		sb.WriteString("\t")

		if i%5 == 4 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
