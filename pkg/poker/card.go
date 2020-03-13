package poker

import (
	"bytes"
	"strconv"
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
