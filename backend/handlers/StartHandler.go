package handlers

import (
	"encoding/json"
	"net/http"

	"backend/models"
)

type StartHandler struct {
	Game *models.Game
}

func NewStartHandler(game *models.Game) *StartHandler {
	return &StartHandler{game}
}

func (h StartHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		ForbiddenHandler(writer)
		return
	}

	h.Game.DealCards()

	resp := models.NewCardsUpdateResponse(len(h.Game.DealerCards), []string{h.Game.DealerCards[0]},
		len(h.Game.PlayerCards), h.Game.PlayerCards)
	jsonBytes, err := json.Marshal(resp)

	if err != nil {
		InternalServerErrorHandler(writer)
		return
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if _, err = writer.Write(jsonBytes); err != nil {
		InternalServerErrorHandler(writer)
		return
	}
}
