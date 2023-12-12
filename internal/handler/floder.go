package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"go.uber.org/zap"
)

type FloderHandler struct {
	*Handler
	floderService service.FloderService
	deckService   service.DeckService
}

func NewFloderHandler(handler *Handler, deckService service.DeckService, floderService service.FloderService) *FloderHandler {
	return &FloderHandler{
		Handler:       handler,
		floderService: floderService,
		deckService:   deckService,
	}
}

func (h *FloderHandler) GetFloder(ctx *gin.Context) {
	userid := ctx.Query("userid")
	//log.Fatalf("useriddd: %s", userid)
	floderlist, err := h.floderService.FindByUserId(userid)
	if err != nil {
		h.logger.Error("GetFloder", zap.Any("FindByUserId", err))
	}
	floderDeckResps := make([]v1.FloderDeckResp, len(floderlist))
	for i := 0; i < len(floderlist); i++ {
		decks, err := h.deckService.GetDecksByFloderId(floderlist[i].ID)
		if err != nil {
			h.logger.Error("GetFloder", zap.Any("GetDecksByFloderId", err))
		}
		deckrs := make([]v1.DeckResp, len(decks))
		for j := 0; j < len(decks); j++ {
			deckrs[j] = v1.DeckResp{
				ID:           decks[j].ID,
				Name:         decks[j].Name,
				Cardnum:      decks[j].Cardnum,
				Learnnumber:  decks[j].Learnnumber,
				Introduction: decks[j].Introduction,
				Floderid:     decks[j].Floderid,
			}
		}

		floderDeckResps[i] = v1.FloderDeckResp{
			ID:      floderlist[i].ID,
			Name:    floderlist[i].Name,
			Decknum: floderlist[i].Decknum,
			Decks:   deckrs,
		}
		if err != nil {
			return
		}
	}
	v1.HandleSuccess(ctx, floderDeckResps)
}
