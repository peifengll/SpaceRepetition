package deckopen

import (
	"context"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CardNoLearningLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCardNoLearningLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CardNoLearningLogic {
	return &CardNoLearningLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CardNoLearningLogic) CardNoLearning(req *types.CardLearningReq) (resp *types.CardLearningResp, err error) {
	// todo: add your logic here and delete this line

	return
}
