package components

import (
	"A-MATH/err"
	"A-MATH/game/constants"
	"A-MATH/game/utils"

	"strconv"
	"strings"
)

var chipSet = utils.ChipSet
var totalSet = utils.TotalSet

type Bag struct {
	ChipCollectors []chipCollectors
	TotalChipLeft  int
	MaxChip        int
}

func NewBag() Bag {
	var chipCollectors []chipCollectors
	var totalChip int

	for i := range len(chipSet) {
		totalChip += totalSet[i]
		n := NewChipCollector(NewChip(chipSet[i]), totalSet[i])
		chipCollectors = append(chipCollectors, n)
	}

	// In-bag tiles needed
	// ("0", 5),  ("1", 6),  ("2", 6),  ("3", 5),  ("4", 5),
	// ("5", 4),  ("6", 4),  ("7", 4),  ("8", 4),  ("9", 4),
	// ("10", 2), ("11", 1), ("12", 2), ("13", 1), ("14", 1),
	// ("15", 1), ("16", 1), ("17", 1), ("18", 1), ("19", 1),
	// ("20", 1), ("+", 4),  ("-", 4),  ("x", 4),  ("%", 4),
	// ("+-", 5), (*/", 4),  ("=", 11), ("~", 4)

	return Bag{chipCollectors, totalChip, totalChip}
}

func (b *Bag) getIndex(c Chip) (int, error) {
	for i, cs := range chipSet {
		if cs == c.Value {
			return i, nil
		}
	}
	return -1, err.New(int(constants.BadRequest), string(constants.InvalidInputForChip))
}

func (b *Bag) DecreaseChip(c Chip) error {
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

func (b *Bag) IncreaseChip(c Chip) error {
	index, e := b.getIndex(c)
	if e != nil {
		return e
	}

	totalChipInBag := b.ChipCollectors[index].Total
	maximumChipInBag := b.ChipCollectors[index].MaxChip
	if 0 <= totalChipInBag && totalChipInBag < maximumChipInBag {
		b.ChipCollectors[index].IncreaseChip()
		b.TotalChipLeft++
	}
	return nil
}

func (b Bag) String() string {
	var sb strings.Builder

	for i, ch := range b.ChipCollectors {
		sb.WriteString("[")
		sb.WriteString(ch.Chips.Value)
		sb.WriteString("]")
		sb.WriteString(strconv.Itoa(ch.Total))
		sb.WriteString("/")
		sb.WriteString(strconv.Itoa(ch.MaxChip))
		sb.WriteString("\t")

		if i%5 == 4 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
