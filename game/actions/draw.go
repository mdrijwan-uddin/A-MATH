package actions

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"

	"math/rand"
)

// Cheating Method
func (g *Game) DrawSpecificChip(player *InGamePlayers, chip components.Chip) {
	g.Bag.DecreaseChip(chip)
	player.Racks.Add(chip)
}

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

	g.DrawSpecificChip(player, chip)
}

func (g *Game) FullyDraw(player *InGamePlayers) {
	rack := player.Racks
	if rack.IsFull() {
		return
	}

	chipsNeedtoDraw := constants.MaxChipInRack - rack.GetTotalChip()
	for i := 0; i < chipsNeedtoDraw; i++ {
		g.Draw(player)
	}
}

func (g *Game) Exchange(player *InGamePlayers, chips []components.Chip) {
	tempChips := []components.Chip{}
	for _, ch := range chips {
		if player.Racks.IsChipFound(ch) {
			tempChips = append(tempChips, ch)
			player.Racks.Remove(ch)
		}
	}
	g.FullyDraw(player)
	for _, tc := range tempChips {
		g.Bag.IncreaseChip(tc)
	}
}
