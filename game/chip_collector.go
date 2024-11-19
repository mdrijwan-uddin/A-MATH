package game

import "strconv"

type chipCollectors struct {
	Chips   chip
	Total   int
	MaxChip int
}

func NewChipCollector(c chip, total int) chipCollectors {
	return chipCollectors{c, total, total}
}

func (c *chipCollectors) IncreaseChip() {
	if c.Total != c.MaxChip {
		c.Total++
	}
}

func (c *chipCollectors) DecreaseChip() {
	if c.Total > 0 {
		c.Total--
	}
}

func (c chipCollectors) String() string {
	return "{" + c.Chips.Value + "|" +
		strconv.Itoa(c.Total) +
		"/" + strconv.Itoa(c.MaxChip) + "}"
}
