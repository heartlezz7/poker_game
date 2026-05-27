package main

import (
	"fmt"
	"poker_game/card"
	"poker_game/player"
	"strings"
)

const (
	cardsPerHand = 5
	numPlayers   = 4
)

func main() {
	deck := card.NewDeck()
	deck.Shuffle()

	players := make([]*player.Player, numPlayers)
	for i := range players {
		players[i] = player.NewPlayer(fmt.Sprintf("Player %d", i+1))
	}

	if err := deal(deck, players); err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, p := range players {
		res, err := player.Evaluate(p.Hand)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		p.Result = res
		fmt.Printf("%s: %s -> %s\n", p.Name, formatHand(p.Hand), res.Rank)
	}

	fmt.Printf("\nCards left in deck: %d\n", deck.Remaining())

	winners := findWinners(players)
	if len(winners) == 1 {
		fmt.Printf("\n*** Winner is %s with %s! ***\n", winners[0].Name, winners[0].Result.Rank)
	} else {
		names := make([]string, len(winners))
		for i, w := range winners {
			names[i] = w.Name
		}
		fmt.Printf("*** Tie between %s with %s! ***\n", strings.Join(names, ", "), winners[0].Result.Rank)
	}
}

func deal(deck *card.Deck, players []*player.Player) error {
	for _, p := range players {
		cards, err := deck.DrawN(cardsPerHand)
		if err != nil {
			return err
		}
		p.Hand = cards
	}
	return nil
}

func formatHand(hand []card.Card) string {
	parts := make([]string, len(hand))
	for i, c := range hand {
		parts[i] = c.String()
	}
	return strings.Join(parts, " ")
}

func findWinners(players []*player.Player) []*player.Player {
	winners := []*player.Player{players[0]}
	for _, p := range players[1:] {
		cmp := player.Compare(p.Result, winners[0].Result)
		switch {
		case cmp > 0:
			winners = []*player.Player{p}
		case cmp == 0:
			winners = append(winners, p)
		}
	}
	return winners
}
