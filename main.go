package main

import (
	"fmt"
	"strings"
)

const (
	cardsPerHand = 5
	numPlayers   = 4
)

type Player struct {
	Name   string
	Hand   []Card
	Result HandResult
}

func main() {
	deck := NewDeck()
	deck.Shuffle()

	players := make([]*Player, numPlayers)
	for i := range players {
		players[i] = &Player{Name: fmt.Sprintf("Player %d", i+1)}
	}

	if err := deal(deck, players); err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, p := range players {
		res, err := Evaluate(p.Hand)
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

func deal(deck *Deck, players []*Player) error {
	for _, p := range players {
		cards, err := deck.DrawN(cardsPerHand)
		if err != nil {
			return err
		}
		p.Hand = cards
	}
	return nil
}

func formatHand(hand []Card) string {
	parts := make([]string, len(hand))
	for i, c := range hand {
		parts[i] = c.String()
	}
	return strings.Join(parts, " ")
}

func findWinners(players []*Player) []*Player {
	winners := []*Player{players[0]}
	for _, p := range players[1:] {
		cmp := Compare(p.Result, winners[0].Result)
		switch {
		case cmp > 0:
			winners = []*Player{p}
		case cmp == 0:
			winners = append(winners, p)
		}
	}
	return winners
}
