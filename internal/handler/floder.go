package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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

// 展示所有文件夹，以及第一个文件夹的情况
func (h *FloderHandler) GetFloder(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	//userid := ctx.Query("userid")
	//log.Fatalf("useriddd: %s", userid)
	floderlist, err := h.floderService.FindByUserId(userId)
	if err != nil {
		h.logger.Error("GetFloder", zap.Any("FindByUserId", err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)

	}
	floderDeckResps := make([]v1.FloderDeckResp, len(floderlist))
	for i := 0; i < len(floderlist); i++ {
		decks, err := h.deckService.GetDecksByFloderId(floderlist[i].ID)
		if err != nil {
			h.logger.Error("GetFloder", zap.Any("GetDecksByFloderId", err))
			v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
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
			v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		}
	}
	v1.HandleSuccess(ctx, floderDeckResps)
}

func (h *FloderHandler) AddFloder(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var f v1.FloderRequest
	if err := ctx.ShouldBindJSON(&f); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if err := h.floderService.AddFloder(userId, f.Name); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *FloderHandler) DeleteFloder(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	id := ctx.Param("id")
	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	var isAccess bool
	if isAccess, err = h.floderService.CheckAccess(iid, userId); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if !isAccess {
		v1.HandleError(ctx, http.StatusForbidden, v1.ErrNoAccessDeleteFloder, nil)
		return
	}

	if err := h.floderService.DeleteFloder(iid); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *FloderHandler) GetFloderById(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	id := ctx.Param("id")
	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	h.floderService.GetFloderById(iid)

}

func (h *FloderHandler) UpdateFloder(c *gin.Context) {
	userId := GetUserIdFromCtx(c)
	if userId == "" {
		v1.HandleError(c, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var err error
	var f v1.FloderRequest
	if err := c.ShouldBindJSON(&f); err != nil {
		v1.HandleError(c, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	var isAccess bool
	if isAccess, err = h.floderService.CheckAccess(f.ID, userId); err != nil {
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	if !isAccess {
		v1.HandleError(c, http.StatusForbidden, v1.ErrNoAccessDeleteFloder, nil)
		return
	}
	ff := &model.Floder{ID: f.ID, Name: f.Name}
	if err := h.floderService.UpdateFloder(ff); err != nil {
		v1.HandleError(c, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(c, nil)
}
