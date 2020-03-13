package poker

import (
	"bytes"
	"strconv"
)

// CardSuit is an enum of the suit of the card
type CardSuit int

const (
	clubs CardSuit = iota
	diamonds
	hearts
	spades
)

// CardRank is an enum of the suite of the card
type CardRank int

const (
	jack CardRank = iota + 11
	queens
	kings
	ace
)

type card struct {
	Suit CardSuit
	Rank CardRank
}

func (c card) String() string {
	var buffer bytes.Buffer

	switch c.Suit {
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
	switch c.Rank {
	case jack:
		buffer.WriteString("Jack")
	case queens:
		buffer.WriteString("Queens")
	case kings:
		buffer.WriteString("Kings")
	case ace:
		buffer.WriteString("Ace")
	default:
		buffer.WriteString(strconv.Itoa(int(c.Rank)))
	}

	return buffer.String()
}
