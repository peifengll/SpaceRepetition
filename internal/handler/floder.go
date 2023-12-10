package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"go.uber.org/zap"
	"strconv"
)

type FloderHandler struct {
	*Handler
	floderService service.FloderService
}

func NewFloderHandler(handler *Handler, floderService service.FloderService) *FloderHandler {
	return &FloderHandler{
		Handler:       handler,
		floderService: floderService,
	}
}

func (h *FloderHandler) GetFloder(ctx *gin.Context) {
	userid := ctx.Query("userid")
	id, err := strconv.ParseInt(userid, 10, 64)
	//log.Fatalf("useriddd: %s", userid)
	if err != nil {
		h.logger.Error("GetFloder", zap.Any("get userid", err))
	}
	floderlist, err := h.floderService.FindByUserId(id)
	if err != nil {
		h.logger.Error("GetFloder", zap.Any("FindByUserId", err))
	}
	v1.HandleSuccess(ctx, floderlist)
}
