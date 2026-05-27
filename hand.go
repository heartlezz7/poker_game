package main

import (
	"errors"
	"sort"
)

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func (h HandRank) String() string {
	return [...]string{
		"High Card", "One Pair", "Two Pair", "Three of a Kind",
		"Straight", "Flush", "Full House", "Four of a Kind",
		"Straight Flush", "Royal Flush",
	}[h]
}

type HandResult struct {
	Rank     HandRank
	Tiebreak []Rank
}

// Evaluate ranks a 5-card poker hand. Tiebreak ranks are returned in
// descending priority order so two HandResults of the same Rank can be
// compared by lexicographic order of Tiebreak.
func Evaluate(hand []Card) (HandResult, error) {
	if len(hand) != 5 {
		return HandResult{}, errors.New("hand must contain exactly 5 cards")
	}

	ranks := make([]Rank, 5)
	for i, c := range hand {
		ranks[i] = c.Rank
	}
	sort.Slice(ranks, func(i, j int) bool { return ranks[i] > ranks[j] })

	flush := isFlush(hand)
	straight, straightHigh := isStraight(ranks)

	counts := rankCounts(ranks)
	groups := groupByCount(counts)

	switch {
	case flush && straight && straightHigh == Ace:
		return HandResult{Rank: RoyalFlush, Tiebreak: []Rank{straightHigh}}, nil
	case flush && straight:
		return HandResult{Rank: StraightFlush, Tiebreak: []Rank{straightHigh}}, nil
	case groups[4] != nil:
		return HandResult{Rank: FourOfAKind, Tiebreak: append(groups[4], groups[1]...)}, nil
	case groups[3] != nil && groups[2] != nil:
		return HandResult{Rank: FullHouse, Tiebreak: append(groups[3], groups[2]...)}, nil
	case flush:
		return HandResult{Rank: Flush, Tiebreak: ranks}, nil
	case straight:
		return HandResult{Rank: Straight, Tiebreak: []Rank{straightHigh}}, nil
	case groups[3] != nil:
		return HandResult{Rank: ThreeOfAKind, Tiebreak: append(groups[3], groups[1]...)}, nil
	case len(groups[2]) == 2:
		return HandResult{Rank: TwoPair, Tiebreak: append(groups[2], groups[1]...)}, nil
	case len(groups[2]) == 1:
		return HandResult{Rank: OnePair, Tiebreak: append(groups[2], groups[1]...)}, nil
	default:
		return HandResult{Rank: HighCard, Tiebreak: ranks}, nil
	}
}

// Compare returns 1 if a beats b, -1 if b beats a, 0 if tied.
func Compare(a, b HandResult) int {
	if a.Rank != b.Rank {
		if a.Rank > b.Rank {
			return 1
		}
		return -1
	}
	for i := 0; i < len(a.Tiebreak) && i < len(b.Tiebreak); i++ {
		if a.Tiebreak[i] != b.Tiebreak[i] {
			if a.Tiebreak[i] > b.Tiebreak[i] {
				return 1
			}
			return -1
		}
	}
	return 0
}

func isFlush(hand []Card) bool {
	s := hand[0].Suit
	for _, c := range hand[1:] {
		if c.Suit != s {
			return false
		}
	}
	return true
}

// isStraight expects ranks sorted descending. The loop handles every
// run of consecutive ranks (e.g. 9-10-J-Q-K → King, 10-J-Q-K-A → Ace),
// and the wheel A-2-3-4-5 is a special case where Ace plays low and the
// straight's high card is Five.
func isStraight(ranks []Rank) (bool, Rank) {
	if isWheel(ranks) {
		return true, Five
	}
	for i := range len(ranks) - 1 {
		if ranks[i]-1 != ranks[i+1] {
			return false, 0
		}
	}
	return true, ranks[0]
}

func isWheel(ranks []Rank) bool {
	if ranks[0] != Ace {
		return false
	}
	for i := 1; i < len(ranks); i++ {
		want := Rank(len(ranks) - i + 1)
		if ranks[i] != want {
			return false
		}
	}
	return true
}

func rankCounts(ranks []Rank) map[Rank]int {
	m := make(map[Rank]int, 5)
	for _, r := range ranks {
		m[r]++
	}
	return m
}

// groupByCount returns a map from count -> ranks with that count, each rank
// list sorted descending so the strongest group sits in front.
func groupByCount(counts map[Rank]int) map[int][]Rank {
	out := make(map[int][]Rank)
	for r, c := range counts {
		out[c] = append(out[c], r)
	}
	for _, list := range out {
		sort.Slice(list, func(i, j int) bool { return list[i] > list[j] })
	}
	return out
}
