package playingcards

import (
	"math/rand/v2"
	"strings"
)

// NewDeck creates a new face-up, 52-card deck
func NewDeck() Deck {
	var d Deck

	for suit := range 4 {
		for rank := range 13 {
			rank++
			d = append(d, Card{Rank: uint8(rank), Suit: uint8(suit), Hidden: false})
		}
	}

	return d
}

type Deck []Card

// Copy creates a copy of the deck
func (d Deck) Copy() Deck {
	return d
}

// Shuffle shuffles the deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

// Flip flips the visibility of all cards in the deck
func (d *Deck) Flip() {
	for i := range *d {
		(*d)[i].Hidden = !(*d)[i].Hidden
	}
}

func (d Deck) String() string {
	var sb strings.Builder
	for _, card := range d {
		sb.WriteString(card.String())
	}
	return sb.String()
}
