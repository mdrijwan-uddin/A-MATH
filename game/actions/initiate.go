package actions

import (
	"A-MATH/game/components"
	"A-MATH/game/players"
	"A-MATH/game/utils"
	"strconv"
	"strings"
)

type Game struct {
	ID          string
	GamePlayers []InGamePlayers
	Board       *components.Board
	Bag         *components.Bag
	Turn        int
	Winner      string
}

type InGamePlayers struct {
	Players players.Player
	Score   int
	Racks   *components.Rack
}

func NewInGamePlayers(player players.Player) InGamePlayers {
	rack := components.NewRack()
	return InGamePlayers{player, 0, &rack}
}

func NewGame(players [2]players.Player) Game {
	gameID := utils.RandomString(8)
	gamePlayers := []InGamePlayers{NewInGamePlayers(players[0]), NewInGamePlayers(players[1])}
	board := components.NewBoard()
	bag := components.NewBag()
	turn := 1
	winner := ""
	return Game{gameID, gamePlayers, &board, &bag, turn, winner}
}

func (p InGamePlayers) String() string {
	var sb strings.Builder
	sb.WriteString(p.Players.String())
	sb.WriteString(" |Score: ")
	sb.WriteString(strconv.Itoa(p.Score))
	sb.WriteString(" |")
	sb.WriteString(p.Racks.String())
	sb.WriteString(" |")
	return sb.String()
}

func (g Game) String() string {
	var sb strings.Builder
	sb.WriteString("{Game ID: ")
	sb.WriteString(g.ID)
	sb.WriteString("}\n")
	sb.WriteString(g.GamePlayers[0].String())
	sb.WriteString("\n")
	sb.WriteString(g.GamePlayers[1].String())
	sb.WriteString("\n")
	sb.WriteString("----------------------------------------------------------------------------------\n")
	sb.WriteString(g.Board.String())
	sb.WriteString("----------------------------------------------------------------------------------\n")
	sb.WriteString(g.Bag.String())
	return sb.String()
}
