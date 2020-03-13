package poker

import (
	"fmt"
	"math/rand"
)

// Dealer represents a Poker dealer
type Dealer struct {
	cards       []card
	openedCards []card
	finished    bool
}

func (d *Dealer) shuffleCards() {
	d.cards = []card{}
	for suit := clubs; suit <= spades; suit++ {
		for rank := CardRank(1); rank <= ace; rank++ {
			d.cards = append(d.cards, card{Suit: suit, Rank: rank})
		}
	}
	rand.Shuffle(
		len(d.cards),
		func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] },
	)
}

func (d *Dealer) reset() {
	d.openedCards = []card{}
	d.finished = false
}

// HasFinished returns true if the dealer cannot make any move, otherwise false
func (d *Dealer) HasFinished() bool {
	return d.finished
}

func (d *Dealer) dealCards(users []*User) {
	fmt.Println("dealer: deal cards")
	d.shuffleCards()
	for i := 0; i < 2; i++ {
		for _, u := range users {
			u.Cards[i], d.cards = d.cards[0], d.cards[1:]
		}
	}
}

// OpenCard opens a card on the table
func (d *Dealer) OpenCard() {
	card := d.cards[0]
	d.cards = d.cards[1:]
	d.openedCards = append(d.openedCards, card)
	fmt.Printf("dealer: open card %s - round %v\n", card.String(), len(d.openedCards))
}

// NextMove returns the next move of the dealer
func (d *Dealer) NextMove() bool {
	if rounds := len(d.openedCards); rounds < 5 {
		d.OpenCard()
	} else if rounds == 5 {
		d.finished = true
	}
	return true
}
