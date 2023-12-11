package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peifengll/SpaceRepetition/internal/service"
)

type KnowledgeHandler struct {
	*Handler
	knowledgeService service.KnowledgeService
}

func NewKnowledgeHandler(handler *Handler, knowledgeService service.KnowledgeService) *KnowledgeHandler {
	return &KnowledgeHandler{
		Handler:      handler,
		knowledgeService: knowledgeService,
	}
}

func (h *KnowledgeHandler) GetKnowledge(ctx *gin.Context) {

}
