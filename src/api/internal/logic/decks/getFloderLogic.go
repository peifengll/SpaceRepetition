package decks

import (
	"SpaceRepetition/src/internal/model"
	"context"
	"encoding/json"
	"strconv"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFloderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFloderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFloderLogic {
	return &GetFloderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFloderLogic) GetFloder(req *types.FloderReq) (resp *types.FloderResp, err error) {
	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	decks, err := model.DeckNsp.GetDecksByFloder(l.svcCtx.DbEngine, model.DeckNsp.PersonTableName(strconv.FormatInt(id, 10)), req.Id)
	if err != nil {
		return nil, err
	}
	flresp := make([]types.Deckinfo, len(decks))
	for i := 0; i < len(decks); i++ {
		flresp[i].Name = decks[i].Name
		flresp[i].Cardnum = decks[i].Cardnum
		flresp[i].Update = 123
		flresp[i].Id = decks[i].FloderId
	}
	return &types.FloderResp{
		Code:    203,
		Message: "没啥问题",
		Data:    flresp,
	}, nil
}
