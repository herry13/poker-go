package poker

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

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
