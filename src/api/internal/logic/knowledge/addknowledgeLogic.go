package knowledge

import (
	"SpaceRepetition/src/internal/model"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"strconv"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddknowledgeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddknowledgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddknowledgeLogic {
	return &AddknowledgeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddknowledgeLogic) Addknowledge(req *types.KnowledgeAddReq) (resp *types.KnowledgeAddResp, err error) {

	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	k := &model.Knowledge{
		Model:      gorm.Model{},
		Font:       req.Font,
		OriginFont: req.Font,
		Back:       req.Back,
		OnLearning: false,
		Typeof:     0,
		DeckId:     req.DeckId,
		Skilled:    0,
		SectionId:  req.SectionId,
	}
	err = model.KnowledgeNsp.AddKnowledge(l.svcCtx.DbEngine,
		model.KnowledgeNsp.PersonTableName(strconv.FormatInt(id, 10)),
		k)
	if err != nil {
		return nil, err
	}
	return &types.KnowledgeAddResp{
		CommonResply: types.CommonResply{
			Code:    209,
			Message: "没啥",
		},
	}, nil
}
