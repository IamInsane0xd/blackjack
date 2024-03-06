package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Game struct {
	Cards       []string
	DealerCards []string
	PlayerCards []string
	PlayerMoney int64
	Bet         int64
}

func NewGame() *Game {
	var cards []string

	for i := 0; i < 13; i++ {
		value := ""

		switch i {
		case 0:
			value += "A"
			break

		case 10:
			value += "J"
			break

		case 11:
			value += "Q"
			break

		case 12:
			value += "K"
			break

		default:
			value += strconv.FormatInt(int64(i+1), 32)
			break
		}

		for j := 0; j < 4; j++ {
			suite := ""

			switch j {
			case 0:
				suite += "C"
				break

			case 1:
				suite += "D"
				break

			case 2:
				suite += "H"
				break

			case 3:
				suite += "S"
				break
			}

			cards = append(cards, fmt.Sprintf("%s-%s", value, suite))
		}
	}

	return &Game{
		cards,
		nil,
		nil,
		0,
		0,
	}
}

func (g *Game) DealCards() {
	toPlayer := true

	for i := 0; i < 4; i++ {
		cardIndex := rand.Intn(len(g.Cards) + 1)
		card := g.Cards[cardIndex]
		g.Cards = append(g.Cards[:cardIndex], g.Cards[cardIndex+1:]...)

		if toPlayer {
			g.PlayerCards = append(g.PlayerCards, card)
		} else {
			g.DealerCards = append(g.DealerCards, card)
		}

		toPlayer = !toPlayer
	}
}
