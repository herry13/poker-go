package poker

import "fmt"

// Table is a poker table
type Table struct {
	Dealer *Dealer
	Users  []*User
}

// Player is a poker player who gets turns to play/deal cards.
type Player interface {
	NextMove() bool // return false if forfeit, otherwise true
}

// Play starts a poker game
func (t *Table) Play() {
	t.Dealer.reset()
	t.Dealer.dealCards(t.Users)

	players := []Player{}
	for _, u := range t.Users {
		players = append(players, u)
		fmt.Printf("%s: cards -> %s %s\n", u.Name, u.Cards[0], u.Cards[1])
	}
	players = append(players, t.Dealer)

	// open the first 3 cards
	for i := 0; i < 3; i++ {
		t.Dealer.OpenCard()
	}

	for len(players) > 2 && !t.Dealer.HasFinished() {
		player := players[0]
		players = players[1:]
		if player.NextMove() {
			players = append(players, player)
		}
	}

	if len(players) > 2 {
		fmt.Println("It's a DRAW!")
	} else {
		for _, player := range players {
			if p, ok := player.(*User); ok {
				fmt.Printf("%v is the winner!\n", p.Name)
			}
		}
	}
}
