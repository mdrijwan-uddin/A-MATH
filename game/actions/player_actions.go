package actions

import (
	"A-MATH/game/components"
	"A-MATH/game/models"
)

func (g *Game) Submit(player *inGamePlayers, c []models.ChipForPlacing) {
	if !g.isPlayersTurn(player) {
		return //add error
	}

	for _, chip := range c {
		g.Board.Add(chip.Position, chip.Chip)
		player.Racks.Remove(chip.Chip)
	}

	g.RefillRack(player)
	g.EndTurn(player)
}

func (g *Game) Exchange(player *inGamePlayers, chips []components.Chip) {
	if !g.isPlayersTurn(player) {
		return //add error
	}

	tempChips := []components.Chip{}
	for _, ch := range chips {
		if player.Racks.IsChipFound(ch) {
			tempChips = append(tempChips, ch)
			player.Racks.Remove(ch)
		}
	}

	g.RefillRack(player)
	for _, tc := range tempChips {
		g.Bag.IncreaseChip(tc)
	}

	g.EndTurn(player)
}

func (g *Game) EndTurn(player *inGamePlayers) {
	if !g.isPlayersTurn(player) {
		return //add error
	}
	g.Turn++
}

func (g *Game) isPlayersTurn(player *inGamePlayers) bool {
	playersTurn := (g.Turn + 1) % 2
	return g.GamePlayers[playersTurn].Players.ID == player.Players.ID
}
