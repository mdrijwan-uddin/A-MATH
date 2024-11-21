package actions

import (
	"A-MATH/game/components"
	"math/rand"
)

func (g *Game) Draw(player *InGamePlayers) {
	totalChipLeft := g.Bag.TotalChipLeft
	if totalChipLeft == 0 {
		return
	}

	index := rand.Intn(totalChipLeft)

	var chip components.Chip
	for _, ch := range g.Bag.ChipCollectors {
		if index > ch.Total {
			index -= ch.Total
		} else {
			chip = ch.Chips
			break
		}
	}

	g.Bag.DecreaseChip(chip)
	player.Racks.Add(chip)
}
