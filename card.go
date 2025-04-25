package playingcards

// NewCard creates a new card with the given rank, suit, and hidden state
func NewCard(rank uint8, suit uint8, hidden bool) Card {
	return Card{
		Rank:   rank,
		Suit:   suit,
		Hidden: hidden,
	}
}

// Unique value for identifing suit of a card
const (
	Spades uint8 = iota
	Clubs
	Hearts
	Diamonds
)

// Card represents a playing card with rank, suit, and hidden state
type Card struct {
	Rank   uint8
	Suit   uint8
	Hidden bool
}

// String returns a string representation of the card
func (c Card) String() string {
	if c.Hidden {
		return "🂠"
	}
	if c.Rank > 13 || c.Rank < 1 || c.Suit > Diamonds {
		return "?"
	}

	var ace rune

	switch c.Suit {
	case Spades:
		ace = '🂡'
	case Clubs:
		ace = '🃑'
	case Hearts:
		ace = '🂱'
	case Diamonds:
		ace = '🃁'
	}

	// dont include knights
	if c.Rank > 11 {
		c.Rank++
	}

	return string(ace + rune(c.Rank-1))
}
