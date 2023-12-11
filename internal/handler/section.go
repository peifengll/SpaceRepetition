package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peifengll/SpaceRepetition/internal/service"
)

type SectionHandler struct {
	*Handler
	sectionService service.SectionService
}

func NewSectionHandler(handler *Handler, sectionService service.SectionService) *SectionHandler {
	return &SectionHandler{
		Handler:      handler,
		sectionService: sectionService,
	}
}

func (h *SectionHandler) GetSection(ctx *gin.Context) {

}
