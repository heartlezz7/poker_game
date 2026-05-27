package player

import (
	"poker_game/card"
	"testing"
)

// h is a short builder for building hands in test tables.
func h(cards ...card.Card) []card.Card { return cards }

func c(s card.Suit, r card.Rank) card.Card { return card.Card{Suit: s, Rank: r} }

func TestEvaluate(t *testing.T) {
	cases := []struct {
		name     string
		hand     []card.Card
		wantRank HandRank
		// wantTiebreak is optional; nil skips the check.
		wantTiebreak []card.Rank
	}{
		// --- Royal Flush ---
		{
			name:         "royal flush in spades",
			hand:         h(c(card.Spades, card.Ten), c(card.Spades, card.Jack), c(card.Spades, card.Queen), c(card.Spades, card.King), c(card.Spades, card.Ace)),
			wantRank:     RoyalFlush,
			wantTiebreak: []card.Rank{card.Ace},
		},
		{
			name:         "royal flush in hearts",
			hand:         h(c(card.Hearts, card.Ten), c(card.Hearts, card.Jack), c(card.Hearts, card.Queen), c(card.Hearts, card.King), c(card.Hearts, card.Ace)),
			wantRank:     RoyalFlush,
			wantTiebreak: []card.Rank{card.Ace},
		},

		// --- Straight Flush ---
		{
			name:         "straight flush 9-K",
			hand:         h(c(card.Clubs, card.Nine), c(card.Clubs, card.Ten), c(card.Clubs, card.Jack), c(card.Clubs, card.Queen), c(card.Clubs, card.King)),
			wantRank:     StraightFlush,
			wantTiebreak: []card.Rank{card.King},
		},
		{
			name:         "steel wheel (A-2-3-4-5 same suit)",
			hand:         h(c(card.Diamonds, card.Ace), c(card.Diamonds, card.Two), c(card.Diamonds, card.Three), c(card.Diamonds, card.Four), c(card.Diamonds, card.Five)),
			wantRank:     StraightFlush,
			wantTiebreak: []card.Rank{card.Five},
		},
		{
			name:         "straight flush 5-9",
			hand:         h(c(card.Hearts, card.Five), c(card.Hearts, card.Six), c(card.Hearts, card.Seven), c(card.Hearts, card.Eight), c(card.Hearts, card.Nine)),
			wantRank:     StraightFlush,
			wantTiebreak: []card.Rank{card.Nine},
		},

		// --- Four of a Kind ---
		{
			name:         "four aces with king kicker",
			hand:         h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.Ace), c(card.Spades, card.Ace), c(card.Clubs, card.King)),
			wantRank:     FourOfAKind,
			wantTiebreak: []card.Rank{card.Ace, card.King},
		},
		{
			name:         "four twos with three kicker",
			hand:         h(c(card.Clubs, card.Two), c(card.Diamonds, card.Two), c(card.Hearts, card.Two), c(card.Spades, card.Two), c(card.Clubs, card.Three)),
			wantRank:     FourOfAKind,
			wantTiebreak: []card.Rank{card.Two, card.Three},
		},

		// --- Full House ---
		{
			name:         "full house: kings full of twos",
			hand:         h(c(card.Clubs, card.King), c(card.Diamonds, card.King), c(card.Hearts, card.King), c(card.Spades, card.Two), c(card.Clubs, card.Two)),
			wantRank:     FullHouse,
			wantTiebreak: []card.Rank{card.King, card.Two},
		},
		{
			name:         "full house: threes full of aces",
			hand:         h(c(card.Clubs, card.Three), c(card.Diamonds, card.Three), c(card.Hearts, card.Three), c(card.Spades, card.Ace), c(card.Clubs, card.Ace)),
			wantRank:     FullHouse,
			wantTiebreak: []card.Rank{card.Three, card.Ace},
		},

		// --- Flush ---
		{
			name:         "ace-high flush",
			hand:         h(c(card.Spades, card.Two), c(card.Spades, card.Five), c(card.Spades, card.Nine), c(card.Spades, card.Jack), c(card.Spades, card.Ace)),
			wantRank:     Flush,
			wantTiebreak: []card.Rank{card.Ace, card.Jack, card.Nine, card.Five, card.Two},
		},
		{
			name:         "king-high flush",
			hand:         h(c(card.Hearts, card.Three), c(card.Hearts, card.Six), c(card.Hearts, card.Eight), c(card.Hearts, card.Ten), c(card.Hearts, card.King)),
			wantRank:     Flush,
			wantTiebreak: []card.Rank{card.King, card.Ten, card.Eight, card.Six, card.Three},
		},

		// --- Straight ---
		{
			name:         "broadway straight (10-A)",
			hand:         h(c(card.Spades, card.Ten), c(card.Hearts, card.Jack), c(card.Diamonds, card.Queen), c(card.Clubs, card.King), c(card.Spades, card.Ace)),
			wantRank:     Straight,
			wantTiebreak: []card.Rank{card.Ace},
		},
		{
			name:         "mid straight (5-9)",
			hand:         h(c(card.Spades, card.Five), c(card.Hearts, card.Six), c(card.Diamonds, card.Seven), c(card.Clubs, card.Eight), c(card.Hearts, card.Nine)),
			wantRank:     Straight,
			wantTiebreak: []card.Rank{card.Nine},
		},
		{
			name:         "wheel straight (A-5)",
			hand:         h(c(card.Spades, card.Ace), c(card.Hearts, card.Two), c(card.Diamonds, card.Three), c(card.Clubs, card.Four), c(card.Hearts, card.Five)),
			wantRank:     Straight,
			wantTiebreak: []card.Rank{card.Five},
		},
		{
			name:         "9-K straight mixed suits",
			hand:         h(c(card.Spades, card.Nine), c(card.Hearts, card.Ten), c(card.Diamonds, card.Jack), c(card.Clubs, card.Queen), c(card.Spades, card.King)),
			wantRank:     Straight,
			wantTiebreak: []card.Rank{card.King},
		},

		// --- Three of a Kind ---
		{
			name:         "three queens with A,2 kickers",
			hand:         h(c(card.Clubs, card.Queen), c(card.Diamonds, card.Queen), c(card.Hearts, card.Queen), c(card.Spades, card.Ace), c(card.Clubs, card.Two)),
			wantRank:     ThreeOfAKind,
			wantTiebreak: []card.Rank{card.Queen, card.Ace, card.Two},
		},
		{
			name:         "three sevens with J,4 kickers",
			hand:         h(c(card.Clubs, card.Seven), c(card.Diamonds, card.Seven), c(card.Hearts, card.Seven), c(card.Spades, card.Jack), c(card.Clubs, card.Four)),
			wantRank:     ThreeOfAKind,
			wantTiebreak: []card.Rank{card.Seven, card.Jack, card.Four},
		},

		// --- Two Pair ---
		{
			name:         "aces and kings with queen kicker",
			hand:         h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.King), c(card.Spades, card.King), c(card.Clubs, card.Queen)),
			wantRank:     TwoPair,
			wantTiebreak: []card.Rank{card.Ace, card.King, card.Queen},
		},
		{
			name:         "tens and threes with five kicker",
			hand:         h(c(card.Clubs, card.Ten), c(card.Diamonds, card.Ten), c(card.Hearts, card.Three), c(card.Spades, card.Three), c(card.Clubs, card.Five)),
			wantRank:     TwoPair,
			wantTiebreak: []card.Rank{card.Ten, card.Three, card.Five}, // top pair, bottom pair, kicker
		},

		// --- One Pair ---
		{
			name:         "pair of aces with K,Q,J kickers",
			hand:         h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.King), c(card.Spades, card.Queen), c(card.Clubs, card.Jack)),
			wantRank:     OnePair,
			wantTiebreak: []card.Rank{card.Ace, card.King, card.Queen, card.Jack},
		},
		{
			name:         "pair of twos with low kickers",
			hand:         h(c(card.Clubs, card.Two), c(card.Diamonds, card.Two), c(card.Hearts, card.Six), c(card.Spades, card.Four), c(card.Clubs, card.Three)),
			wantRank:     OnePair,
			wantTiebreak: []card.Rank{card.Two, card.Six, card.Four, card.Three},
		},

		// --- High Card ---
		{
			name:         "ace high",
			hand:         h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Jack), c(card.Hearts, card.Eight), c(card.Spades, card.Four), c(card.Hearts, card.Two)),
			wantRank:     HighCard,
			wantTiebreak: []card.Rank{card.Ace, card.Jack, card.Eight, card.Four, card.Two},
		},
		{
			name:         "seven high",
			hand:         h(c(card.Clubs, card.Seven), c(card.Diamonds, card.Five), c(card.Hearts, card.Four), c(card.Spades, card.Three), c(card.Hearts, card.Two)),
			wantRank:     HighCard,
			wantTiebreak: []card.Rank{card.Seven, card.Five, card.Four, card.Three, card.Two},
		},
		{
			name:     "almost-straight (gap) is high card",
			hand:     h(c(card.Clubs, card.Two), c(card.Diamonds, card.Three), c(card.Hearts, card.Four), c(card.Spades, card.Five), c(card.Hearts, card.Seven)),
			wantRank: HighCard,
		},
		{
			name:     "four-suit near-flush is high card",
			hand:     h(c(card.Spades, card.Two), c(card.Spades, card.Five), c(card.Spades, card.Nine), c(card.Spades, card.Jack), c(card.Hearts, card.Ace)),
			wantRank: HighCard,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Evaluate(tc.hand)
			if err != nil {
				t.Fatalf("Evaluate returned error: %v", err)
			}
			if res.Rank != tc.wantRank {
				t.Fatalf("Rank = %s, want %s", res.Rank, tc.wantRank)
			}
			if tc.wantTiebreak != nil {
				if !rankSliceEqual(res.Tiebreak, tc.wantTiebreak) {
					t.Fatalf("Tiebreak = %v, want %v", res.Tiebreak, tc.wantTiebreak)
				}
			}
		})
	}
}

func TestEvaluate_Errors(t *testing.T) {
	cases := []struct {
		name string
		hand []card.Card
	}{
		{"empty hand", h()},
		{"four cards", h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.Ace), c(card.Spades, card.Ace))},
		{"six cards", h(c(card.Clubs, card.Two), c(card.Diamonds, card.Three), c(card.Hearts, card.Four), c(card.Spades, card.Five), c(card.Hearts, card.Six), c(card.Clubs, card.Seven))},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := Evaluate(tc.hand); err == nil {
				t.Fatalf("expected error for %s, got nil", tc.name)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	// Reusable evaluated hands.
	royalFlush := mustEval(t, h(c(card.Spades, card.Ten), c(card.Spades, card.Jack), c(card.Spades, card.Queen), c(card.Spades, card.King), c(card.Spades, card.Ace)))
	straightFlushKing := mustEval(t, h(c(card.Clubs, card.Nine), c(card.Clubs, card.Ten), c(card.Clubs, card.Jack), c(card.Clubs, card.Queen), c(card.Clubs, card.King)))
	straightFlushNine := mustEval(t, h(c(card.Hearts, card.Five), c(card.Hearts, card.Six), c(card.Hearts, card.Seven), c(card.Hearts, card.Eight), c(card.Hearts, card.Nine)))
	quadAces := mustEval(t, h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.Ace), c(card.Spades, card.Ace), c(card.Clubs, card.Two)))
	quadKings := mustEval(t, h(c(card.Clubs, card.King), c(card.Diamonds, card.King), c(card.Hearts, card.King), c(card.Spades, card.King), c(card.Clubs, card.Two)))
	fullKingsOverTwos := mustEval(t, h(c(card.Clubs, card.King), c(card.Diamonds, card.King), c(card.Hearts, card.King), c(card.Spades, card.Two), c(card.Clubs, card.Two)))
	fullThreesOverAces := mustEval(t, h(c(card.Clubs, card.Three), c(card.Diamonds, card.Three), c(card.Hearts, card.Three), c(card.Spades, card.Ace), c(card.Clubs, card.Ace)))
	aceFlush := mustEval(t, h(c(card.Spades, card.Two), c(card.Spades, card.Five), c(card.Spades, card.Nine), c(card.Spades, card.Jack), c(card.Spades, card.Ace)))
	kingFlush := mustEval(t, h(c(card.Hearts, card.Three), c(card.Hearts, card.Six), c(card.Hearts, card.Eight), c(card.Hearts, card.Ten), c(card.Hearts, card.King)))
	broadway := mustEval(t, h(c(card.Spades, card.Ten), c(card.Hearts, card.Jack), c(card.Diamonds, card.Queen), c(card.Clubs, card.King), c(card.Spades, card.Ace)))
	wheel := mustEval(t, h(c(card.Spades, card.Ace), c(card.Hearts, card.Two), c(card.Diamonds, card.Three), c(card.Clubs, card.Four), c(card.Hearts, card.Five)))
	tripsQueens := mustEval(t, h(c(card.Clubs, card.Queen), c(card.Diamonds, card.Queen), c(card.Hearts, card.Queen), c(card.Spades, card.Seven), c(card.Clubs, card.Two)))
	tripsSevens := mustEval(t, h(c(card.Clubs, card.Seven), c(card.Diamonds, card.Seven), c(card.Hearts, card.Seven), c(card.Spades, card.Ace), c(card.Clubs, card.King)))
	twoPairAcesKings := mustEval(t, h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.King), c(card.Spades, card.King), c(card.Clubs, card.Two)))
	twoPairAcesKingsBigKicker := mustEval(t, h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.King), c(card.Spades, card.King), c(card.Clubs, card.Queen)))
	twoPairAcesQueens := mustEval(t, h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.Queen), c(card.Spades, card.Queen), c(card.Clubs, card.Two)))
	pairAcesKKicker := mustEval(t, h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.King), c(card.Spades, card.Seven), c(card.Clubs, card.Two)))
	pairAcesQKicker := mustEval(t, h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Ace), c(card.Hearts, card.Queen), c(card.Spades, card.Seven), c(card.Clubs, card.Two)))
	pairKings := mustEval(t, h(c(card.Clubs, card.King), c(card.Diamonds, card.King), c(card.Hearts, card.Ace), c(card.Spades, card.Seven), c(card.Clubs, card.Two)))
	highAce := mustEval(t, h(c(card.Clubs, card.Ace), c(card.Diamonds, card.Jack), c(card.Hearts, card.Eight), c(card.Spades, card.Four), c(card.Hearts, card.Two)))
	highKing := mustEval(t, h(c(card.Clubs, card.King), c(card.Diamonds, card.Jack), c(card.Hearts, card.Eight), c(card.Spades, card.Four), c(card.Hearts, card.Two)))

	cases := []struct {
		name string
		a, b HandResult
		want int
	}{
		// rank-vs-rank
		{"royal beats straight flush", royalFlush, straightFlushKing, 1},
		{"straight flush beats quads", straightFlushKing, quadAces, 1},
		{"quads beat full house", quadAces, fullKingsOverTwos, 1},
		{"full house beats flush", fullKingsOverTwos, aceFlush, 1},
		{"flush beats straight", aceFlush, broadway, 1},
		{"straight beats trips", broadway, tripsQueens, 1},
		{"trips beat two pair", tripsQueens, twoPairAcesKings, 1},
		{"two pair beat one pair", twoPairAcesKings, pairAcesKKicker, 1},
		{"one pair beats high card", pairAcesKKicker, highAce, 1},

		// same-rank tiebreakers
		{"higher straight flush wins", straightFlushKing, straightFlushNine, 1},
		{"quads compare by quad rank", quadAces, quadKings, 1},
		{"full house compares by trips first", fullKingsOverTwos, fullThreesOverAces, 1},
		{"higher flush wins", aceFlush, kingFlush, 1},
		{"broadway beats wheel", broadway, wheel, 1},
		{"trips compare by trip rank", tripsQueens, tripsSevens, 1},
		{"two pair compare by top pair", twoPairAcesKings, twoPairAcesQueens, 1},
		{"two pair compare by kicker when pairs match", twoPairAcesKingsBigKicker, twoPairAcesKings, 1},
		{"pair compares by kicker when pair matches", pairAcesKKicker, pairAcesQKicker, 1},
		{"higher pair wins", pairAcesKKicker, pairKings, 1},
		{"higher high card wins", highAce, highKing, 1},

		// ties
		{"identical hands tie", broadway, broadway, 0},
		{"identical flushes tie", aceFlush, aceFlush, 0},

		// reverse direction
		{"reverse: lower hand loses", highKing, highAce, -1},
		{"reverse: pair loses to two pair", pairAcesKKicker, twoPairAcesKings, -1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Compare(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Compare = %d, want %d (a=%s %v, b=%s %v)",
					got, tc.want, tc.a.Rank, tc.a.Tiebreak, tc.b.Rank, tc.b.Tiebreak)
			}
		})
	}
}

func mustEval(t *testing.T, hand []card.Card) HandResult {
	t.Helper()
	res, err := Evaluate(hand)
	if err != nil {
		t.Fatalf("Evaluate: %v", err)
	}
	return res
}

func rankSliceEqual(a, b []card.Rank) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
