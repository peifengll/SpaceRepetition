package deckopen

import (
	"context"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CardOnLearningLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCardOnLearningLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CardOnLearningLogic {
	return &CardOnLearningLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CardOnLearningLogic) CardOnLearning(req *types.CardLearningReq) (resp *types.CardLearningResp, err error) {
	// todo: add your logic here and delete this line

	return
}
