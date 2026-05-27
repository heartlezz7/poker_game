# Poker Game

A 5-card poker hand evaluator written in Go. Builds a shuffled 52-card deck, deals to N players, evaluates each hand, and picks the winner (or detects a tie).

## Requirements

- Go 1.22 or newer (uses `for i := range N` integer iteration).

## Run

```bash
go run .
```

Example output:

```
Player 1: [A Spades] [10 Spades] [5 Spades] [2 Spades] [K Spades] -> Flush
Player 2: [J Diamonds] [J Clubs] [4 Hearts] [9 Spades] [7 Clubs] -> One Pair
Player 3: [Q Hearts] [8 Spades] [2 Clubs] [6 Diamonds] [A Clubs] -> High Card
Player 4: [K Hearts] [K Diamonds] [K Clubs] [10 Spades] [10 Hearts] -> Full House

Cards left in deck: 32

*** Winner is Player 4 with Full House! ***
```

## Test

```bash
go test ./...
```

The `player` package ships ~50 subtests covering every hand rank, tiebreak ordering, and error paths.

## Project layout

```
poker_game/
├── main.go              # demo: deal to N players, evaluate, pick winner
├── card/
│   ├── card.go          # Card, Suit, Rank types and Stringer methods
│   └── deck.go          # Deck: NewDeck, Shuffle, Draw, DrawN, Remaining
└── player/
    ├── player.go        # Player struct
    ├── hand.go          # Evaluate + Compare + hand-classification helpers
    └── hand_test.go     # comprehensive test suite
```

## API overview

### `card` package

```go
deck := card.NewDeck()        // fresh 52-card deck (Clubs..Spades x Two..Ace)
deck.Shuffle()                // in-place Fisher-Yates via math/rand
c, err := deck.Draw()         // draw one card
cards, err := deck.DrawN(5)   // draw n cards
deck.Remaining()              // cards left in deck
```

`Card` prints as `[Rank Suit]`, e.g. `[K Hearts]`, `[10 Spades]`.

### `player` package

```go
res, err := player.Evaluate(hand)   // hand must be exactly 5 cards
cmp := player.Compare(a, b)         // 1 if a wins, -1 if b wins, 0 if tied
```

`HandResult` carries a `Rank` (one of the 10 standard poker hands) plus a `Tiebreak` slice of `card.Rank` values, ordered so a lexicographic compare resolves ties correctly.

## Hand rankings (lowest to highest)

| Rank | Description | Tiebreak order |
|------|-------------|----------------|
| High Card | No matches | All 5 ranks, descending |
| One Pair | 2 of a kind | Pair rank, then 3 kickers desc |
| Two Pair | 2 + 2 | Top pair, bottom pair, kicker |
| Three of a Kind | 3 of a kind | Trip rank, 2 kickers desc |
| Straight | 5 sequential | High card of the run |
| Flush | 5 of a suit | All 5 ranks, descending |
| Full House | 3 + 2 | Trip rank, pair rank |
| Four of a Kind | 4 of a kind | Quad rank, kicker |
| Straight Flush | Sequential + same suit | High card of the run |
| Royal Flush | 10-J-Q-K-A same suit | Ace |

### Special cases

- **Wheel (A-2-3-4-5):** counted as a straight with Five as the high card, so it loses to 2-3-4-5-6 but beats every non-straight.
- **Steel wheel (A-2-3-4-5 same suit):** treated as a Straight Flush with Five high — *not* a Royal Flush.

## Customizing

Change the number of players or hand size in `main.go`:

```go
const (
    cardsPerHand = 5
    numPlayers   = 4
)
```

With a 52-card deck and 5 cards per hand, up to 10 players are supported.
