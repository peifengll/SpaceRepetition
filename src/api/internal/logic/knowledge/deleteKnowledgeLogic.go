package knowledge

import (
	"context"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteKnowledgeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteKnowledgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteKnowledgeLogic {
	return &DeleteKnowledgeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteKnowledgeLogic) DeleteKnowledge(req *types.KnowledgeDelReq) (resp *types.KnowledgeDelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
