package deckopen

import (
	"SpaceRepetition/src/internal/model"
	"context"
	"encoding/json"
	"strconv"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSectionLogic {
	return &AddSectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSectionLogic) AddSection(req *types.SectionAddReq) (resp *types.CommonResply, err error) {

	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	err = model.SectionNsp.AddSection(l.svcCtx.DbEngine,
		model.SectionNsp.PersonTableName(strconv.FormatInt(id, 10)),
		req.Name,
		req.DeckId)
	if err != nil {
		return nil, err
	}

	return &types.CommonResply{
		Code:    206,
		Message: "正常",
	}, nil
}
