package deckopen

import (
	"SpaceRepetition/src/internal/model"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddToReviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddToReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddToReviewLogic {
	return &AddToReviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 选中后一批进行调度复习
func (l *AddToReviewLogic) AddToReview(req *types.CardToLearningReq) (resp *types.CommonResply, err error) {

	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	for _, kid := range req.IdList {
		err := model.KnowledgeNsp.UpdateOnLearning(l.svcCtx.DbEngine,
			model.KnowledgeNsp.PersonTableName(strconv.FormatInt(id, 10)),
			kid)
		if err != nil {
			return nil, err
		}
		// 新更新,
		re := &model.Record{
			KnowledgeId:   0,
			Due:           time.Now(),
			Stability:     0,
			Difficulty:    0,
			ElapsedDays:   0,
			ScheduledDays: 0,
			Reps:          0,
			Lapses:        0,
			State:         0,
			On:            false,
			LastOp:        0,
			LastReview:    time.Now(),
		}
		err = model.RecordNsp.AddNewRecord(l.svcCtx.DbEngine,
			model.RecordNsp.PersonTableName(strconv.FormatInt(id, 10)),
			re,
		)
		if err != nil {
			return nil, err
		}
	}
	return &types.CommonResply{
		Code:    206,
		Message: "加入复习",
	}, nil
}
