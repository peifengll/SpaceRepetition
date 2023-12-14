package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"net/http"
	"strconv"
)

type SectionHandler struct {
	*Handler
	sectionService service.SectionService
}

func NewSectionHandler(handler *Handler, sectionService service.SectionService) *SectionHandler {
	return &SectionHandler{
		Handler:        handler,
		sectionService: sectionService,
	}
}

func (h *SectionHandler) GetSection(ctx *gin.Context) {

}

func (h *SectionHandler) AddSection(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var s v1.SectionReq
	err = ctx.ShouldBindJSON(&s)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	err = h.sectionService.AddSection(&s, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *SectionHandler) DeleteSection(ctx *gin.Context) {
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
	err = h.sectionService.DeleteSectionById(id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *SectionHandler) UpdateSection(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var err error
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var s v1.SectionReq
	err = ctx.ShouldBindJSON(&s)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	err = h.sectionService.UpdateSectionName(s.ID, s.Name)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
