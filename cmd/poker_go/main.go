package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type cardSuit int

const (
	clubs cardSuit = iota
	diamonds
	hearts
	spades
)

type cardRank int

const (
	jack cardRank = iota + 11
	queens
	kings
	ace
)

type card struct {
	suit cardSuit
	rank cardRank
}

func (c card) String() string {
	var buffer bytes.Buffer

	switch c.suit {
	case clubs:
		buffer.WriteString("Clubs")
	case diamonds:
		buffer.WriteString("Diamonds")
	case hearts:
		buffer.WriteString("Hearts")
	case spades:
		buffer.WriteString("Spades")
	}
	buffer.WriteByte(':')
	switch c.rank {
	case jack:
		buffer.WriteString("Jack")
	case queens:
		buffer.WriteString("Queens")
	case kings:
		buffer.WriteString("Kings")
	case ace:
		buffer.WriteString("Ace")
	default:
		buffer.WriteString(strconv.Itoa(int(c.rank)))
	}

	return buffer.String()
}

type player interface {
	nextMove() bool // return false if forfeit, otherwise true
}

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

type user struct {
	name    string
	isBot   bool
	timeout int // in seconds (0 means no timeout)
	cards   [2]card
}

func (u *user) getUserInput() string {
	timeout := make(chan bool, 1)
	if u.timeout > 0 {
		go func() {
			time.Sleep(3 * time.Second) // waiting user input
			timeout <- true
		}()
	}

	userInput := make(chan string, 1)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%v: next move? ", u.name)
		char, _, err := reader.ReadRune()
		if err != nil {
			panic(fmt.Sprintf("%v\n", err))
		}
		switch char {
		case 'f', 'F':
			userInput <- "forfeit"
		case 'i', 'I':
			userInput <- "increase"
		default:
			userInput <- "bet"
		}
	}()

	select {
	case input := <-userInput:
		return input
	case <-timeout:
		fmt.Println()
		return "forfeit"
	}
}

func (u *user) getBotInput() string {
	moves := []string{
		"forfeit",
		"bet",
		"bet",
		"bet-increase",
		"bet-increase",
	}
	return moves[rand.Intn(len(moves))]
}

func (u *user) nextMove() bool {
	var move string
	if u.isBot {
		move = u.getBotInput()
	} else {
		move = u.getUserInput()
	}
	fmt.Printf("%s: %s\n", u.name, move)
	return move != "forfeit"
}

type table struct {
	dealer *dealer
	users  []*user
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

func main() {
	rand.Seed(time.Now().UnixNano())
	table := table{
		dealer: &dealer{},
		users: []*user{
			&user{name: "player1"},
			&user{name: "player2", isBot: true},
			&user{name: "player3", isBot: true},
		},
	}
	table.play()
}
