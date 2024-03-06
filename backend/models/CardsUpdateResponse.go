package models

type CardsUpdateResponse struct {
	DealerCardCount int      "json:\"dealerCardCount\""
	DealerCards     []string "json:\"dealerCards\""
	PlayerCardCount int      "json:\"playerCardCount\""
	PlayerCards     []string "json:\"playerCards\""
}

func NewCardsUpdateResponse(dealerCardCount int, dealerCards []string, playerCardCount int,
	playerCards []string) *CardsUpdateResponse {
	return &CardsUpdateResponse{
		dealerCardCount,
		dealerCards,
		playerCardCount,
		playerCards,
	}
}
