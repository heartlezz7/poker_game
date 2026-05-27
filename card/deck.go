package card

import (
	"errors"
	"math/rand"
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	d := &Deck{cards: make([]Card, 0, 52)}
	for s := Clubs; s <= Spades; s++ {
		for r := Two; r <= Ace; r++ {
			d.cards = append(d.cards, Card{Suit: s, Rank: r})
		}
	}
	return d
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Draw() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c, nil
}

func (d *Deck) DrawN(n int) ([]Card, error) {
	if n > len(d.cards) {
		return nil, errors.New("not enough cards in deck")
	}
	out := make([]Card, n)
	for i := range n {
		c, _ := d.Draw()
		out[i] = c
	}
	return out, nil
}

func (d *Deck) Remaining() int {
	return len(d.cards)
}
