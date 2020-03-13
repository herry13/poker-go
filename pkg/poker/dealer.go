package poker

import (
	"fmt"
	"math/rand"
)

type dealer struct {
	cards       []card
	openedCards []card
	finished    bool
}

func (d *dealer) shuffleCards() {
	d.cards = []card{}
	for suit := clubs; suit <= spades; suit++ {
		for rank := cardRank(1); rank <= ace; rank++ {
			d.cards = append(d.cards, card{suit: suit, rank: rank})
		}
	}
	rand.Shuffle(
		len(d.cards),
		func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] },
	)
}

func (d *dealer) reset() {
	d.openedCards = []card{}
	d.finished = false
}

func (d *dealer) dealCards(users []*user) {
	fmt.Println("dealer: deal cards")
	d.shuffleCards()
	for i := 0; i < 2; i++ {
		for _, u := range users {
			u.cards[i], d.cards = d.cards[0], d.cards[1:]
		}
	}
}

func (d *dealer) openCard() {
	// open a card
	card := d.cards[0]
	d.cards = d.cards[1:]
	d.openedCards = append(d.openedCards, card)
	fmt.Printf("dealer: open card %s - round %v\n", card.String(), len(d.openedCards))
}

func (d *dealer) nextMove() bool {
	if rounds := len(d.openedCards); rounds < 5 {
		d.openCard()
	} else if rounds == 5 {
		d.finished = true
	}
	return true
}
