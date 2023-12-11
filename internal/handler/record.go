package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peifengll/SpaceRepetition/internal/service"
)

type RecordHandler struct {
	*Handler
	recordService service.RecordService
}

func NewRecordHandler(handler *Handler, recordService service.RecordService) *RecordHandler {
	return &RecordHandler{
		Handler:      handler,
		recordService: recordService,
	}
}

func (h *RecordHandler) GetRecord(ctx *gin.Context) {

}
