package player

import "poker_game/card"

type Player struct {
	Name   string
	Hand   []card.Card
	Result HandResult
}

func NewPlayer(name string) *Player {
	return &Player{Name: name}
}
