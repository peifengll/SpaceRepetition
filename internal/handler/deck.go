package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"net/http"
	"strconv"
)

type DeckHandler struct {
	*Handler
	deckService service.DeckService
}

func NewDeckHandler(handler *Handler, deckService service.DeckService) *DeckHandler {
	return &DeckHandler{
		Handler:     handler,
		deckService: deckService,
	}
}

func (h *DeckHandler) GetDeck(ctx *gin.Context) {

}

func (h *DeckHandler) AddDeck(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var d v1.DeckRequest
	if err = ctx.ShouldBindJSON(&d); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	isAccess, err := h.deckService.CheckAddDeckAccess(d.Floderid, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if !isAccess {
		v1.HandleError(ctx, http.StatusForbidden, v1.ErrNoAccessAddDeck, nil)
		return
	}
	de := &model.Deck{
		Name:         d.Name,
		Cardnum:      0,
		Learnnumber:  0,
		Introduction: d.Introduction,
		Floderid:     d.Floderid,
		UserID:       userId,
	}
	err = h.deckService.AddDeck(de)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *DeckHandler) GetDeckById(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	id := ctx.Param("id")
	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	isAccess, err := h.deckService.CheckDeckAccess(iid, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if !isAccess {
		v1.HandleError(ctx, http.StatusForbidden, v1.ErrNoAccessDeleteFloder, nil)
		return
	}

	all, err := h.deckService.GetDeckAll(iid)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, all)
}

func (h *DeckHandler) UpdateDeck(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var d v1.DeckRequest
	err = ctx.ShouldBindJSON(&d)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, err)
		return
	}
	isAccess, err := h.deckService.CheckDeckAccess(d.ID, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if !isAccess {
		v1.HandleError(ctx, http.StatusForbidden, v1.ErrNoAccessDeleteFloder, nil)
		return
	}
	dec := &model.Deck{
		ID:           d.ID,
		Name:         d.Name,
		Introduction: d.Introduction,
		Floderid:     d.Floderid,
	}
	err = h.deckService.UpdateDeck(dec)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *DeckHandler) DeleteDeck(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	idstr := ctx.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, err)
		return
	}

	isAccess, err := h.deckService.CheckDeckAccess(id, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if !isAccess {
		v1.HandleError(ctx, http.StatusForbidden, v1.ErrNoAccessDeleteDeck, nil)
		return
	}
	err = h.deckService.DeleteDeck(id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
