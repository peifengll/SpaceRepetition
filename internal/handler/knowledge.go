package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"net/http"
	"strconv"
)

type KnowledgeHandler struct {
	*Handler
	knowledgeService service.KnowledgeService
}

func NewKnowledgeHandler(handler *Handler, knowledgeService service.KnowledgeService) *KnowledgeHandler {
	return &KnowledgeHandler{
		Handler:          handler,
		knowledgeService: knowledgeService,
	}
}

func (h *KnowledgeHandler) GetKnowledge(ctx *gin.Context) {

}

func (h *KnowledgeHandler) AddCard(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	var c v1.CardRequest
	err = ctx.ShouldBindJSON(&c)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	err = h.knowledgeService.AddCard(&c, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *KnowledgeHandler) GetCard(ctx *gin.Context) {
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

	card, err := h.knowledgeService.GetKnowledgeById(id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, card)
}

func (h *KnowledgeHandler) UpdateCard(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var c v1.CardRequest
	err = ctx.ShouldBindJSON(&c)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	err = h.knowledgeService.UpdateCard(&c)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *KnowledgeHandler) DeleteCard(ctx *gin.Context) {
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
	err = h.knowledgeService.DeleteCard(id)
	v1.HandleSuccess(ctx, nil)
}

func (h *KnowledgeHandler) SearchCards(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	content := ctx.Param("content")
	cards, err := h.knowledgeService.SearchCards(content, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, cards)
}

func (h *KnowledgeHandler) ChooseToReview(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	var id v1.CardIdReq
	err = ctx.BindJSON(&id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	// 如果已经复习了，复习个蛋
	review, err := h.knowledgeService.CheckReview(id.Id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if review {
		v1.HandleSuccess(ctx, "这条已经加入复习了")
		return
	}
	err = h.knowledgeService.ChooseToReview(id.Id, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func (h *KnowledgeHandler) GetAllReview(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	review, err := h.knowledgeService.GetAllReview(userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, review)
}

func (h *KnowledgeHandler) ReviewOpt(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var t v1.CardReviewOptReq
	err := ctx.BindJSON(&t)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
	}
	recordId, err := h.knowledgeService.ReviewOp(&t, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, map[string]any{"finished": recordId == 0, "rid": recordId})
}
