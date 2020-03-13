package poker

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// User represents a poker user
type User struct {
	Name  string
	IsBot bool
	Cards [2]card

	timeout int // in seconds (0 means no timeout)
}

func (u *User) getUserInput() string {
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
		fmt.Printf("%v: next move? ", u.Name)
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

func (u *User) getBotInput() string {
	moves := []string{
		"forfeit",
		"bet",
		"bet",
		"bet-increase",
		"bet-increase",
	}
	return moves[rand.Intn(len(moves))]
}

// NextMove gets the next move
func (u *User) NextMove() bool {
	var move string
	if u.IsBot {
		move = u.getBotInput()
	} else {
		move = u.getUserInput()
	}
	fmt.Printf("%s: %s\n", u.Name, move)
	return move != "forfeit"
}
