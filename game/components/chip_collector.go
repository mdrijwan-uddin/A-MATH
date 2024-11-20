package components

import (
	"strconv"
	"strings"
)

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
	var sb strings.Builder
	sb.WriteString("{")
	sb.WriteString(c.Chips.Value)
	sb.WriteString("|")
	sb.WriteString(strconv.Itoa(c.Total))
	sb.WriteString("/")
	sb.WriteString(strconv.Itoa(c.MaxChip))
	sb.WriteString("}")
	return sb.String()
}
