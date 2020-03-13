package poker

import "fmt"

type table struct {
	dealer *dealer
	users  []*user
}

type player interface {
	nextMove() bool // return false if forfeit, otherwise true
}

func (t *table) play() {
	t.dealer.reset()
	t.dealer.dealCards(t.users)

	players := []player{}
	for _, u := range t.users {
		players = append(players, u)
		fmt.Printf("%s: cards -> %s %s\n", u.name, u.cards[0], u.cards[1])
	}
	players = append(players, t.dealer)

	// open the first 3 cards
	for i := 0; i < 3; i++ {
		t.dealer.openCard()
	}

	for len(players) > 2 && !t.dealer.finished {
		player := players[0]
		players = players[1:]
		if player.nextMove() {
			players = append(players, player)
		}
	}

	if len(players) > 2 {
		fmt.Println("It's a DRAW!")
	} else {
		for _, player := range players {
			if p, ok := player.(*user); ok {
				fmt.Printf("%v is the winner!\n", p.name)
			}
		}
	}
}
