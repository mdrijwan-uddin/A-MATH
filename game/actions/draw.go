package actions

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"

	"math/rand"
)

// Cheating Method
func (g *Game) DrawSpecificChip(player *inGamePlayers, chip components.Chip) {
	g.Bag.DecreaseChip(chip)
	player.Racks.Add(chip)
}

// Normal Method
func (g *Game) SingleDraw(player *inGamePlayers) {
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

	g.DrawSpecificChip(player, chip)
}

// Draw until full
func (g *Game) RefillRack(player *inGamePlayers) {
	rack := player.Racks
	if rack.IsFull() {
		return
	}

	chipsNeedtoDraw := constants.MaxChipInRack - rack.GetTotalChip()
	for i := 0; i < chipsNeedtoDraw; i++ {
		g.SingleDraw(player)
	}
}
