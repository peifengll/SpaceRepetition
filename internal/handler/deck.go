package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peifengll/SpaceRepetition/internal/service"
)

type DeckHandler struct {
	*Handler
	deckService service.DeckService
}

func NewDeckHandler(handler *Handler, deckService service.DeckService) *DeckHandler {
	return &DeckHandler{
		Handler:      handler,
		deckService: deckService,
	}
}

func (h *DeckHandler) GetDeck(ctx *gin.Context) {

}
