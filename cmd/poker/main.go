package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//poker.dealer{}
	/*table := poker.table{
		dealer: &dealer{},
		users: []*user{
			&user{name: "player1"},
			&user{name: "player2", isBot: true},
			&user{name: "player3", isBot: true},
		},
	}
	table.play()*/
}
