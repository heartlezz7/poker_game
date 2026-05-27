package main

import "testing"

func TestStraights(t *testing.T) {
	cases := []struct {
		name string
		hand []Card
		want HandRank
		high Rank
	}{
		{
			name: "9-10-J-Q-K",
			hand: []Card{
				{Spades, Nine}, {Hearts, Ten}, {Diamonds, Jack}, {Clubs, Queen}, {Spades, King},
			},
			want: Straight,
			high: King,
		},
		{
			name: "10-J-Q-K-A (broadway)",
			hand: []Card{
				{Spades, Ten}, {Hearts, Jack}, {Diamonds, Queen}, {Clubs, King}, {Spades, Ace},
			},
			want: Straight,
			high: Ace,
		},
		{
			name: "A-2-3-4-5 (wheel)",
			hand: []Card{
				{Spades, Ace}, {Hearts, Two}, {Diamonds, Three}, {Clubs, Four}, {Spades, Five},
			},
			want: Straight,
			high: Five,
		},
		{
			name: "9-10-J-Q-K all spades (straight flush)",
			hand: []Card{
				{Spades, Nine}, {Spades, Ten}, {Spades, Jack}, {Spades, Queen}, {Spades, King},
			},
			want: StraightFlush,
			high: King,
		},
		{
			name: "not a straight",
			hand: []Card{
				{Spades, Two}, {Hearts, Four}, {Diamonds, Six}, {Clubs, Eight}, {Spades, Ten},
			},
			want: HighCard,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Evaluate(tc.hand)
			if err != nil {
				t.Fatalf("Evaluate: %v", err)
			}
			if res.Rank != tc.want {
				t.Fatalf("rank = %s, want %s", res.Rank, tc.want)
			}
			if tc.high != 0 && res.Tiebreak[0] != tc.high {
				t.Fatalf("high = %s, want %s", res.Tiebreak[0], tc.high)
			}
		})
	}
}
