package main

import (
	"math/rand"
	"time"

	"github.com/herry13/poker-go/pkg/poker"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	table := poker.Table{
		Dealer: &poker.Dealer{},
		Users: []*poker.User{
			&poker.User{Name: "player1"},
			&poker.User{Name: "player2", IsBot: true},
			&poker.User{Name: "player3", IsBot: true},
		},
	}
	table.Play()
}
