package knowledge

import (
	"context"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReviewingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReviewingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReviewingListLogic {
	return &GetReviewingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReviewingListLogic) GetReviewingList() (resp *types.ReviewingListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
