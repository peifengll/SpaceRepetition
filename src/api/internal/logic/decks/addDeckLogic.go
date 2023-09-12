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

type AddDeckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDeckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeckLogic {
	return &AddDeckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDeckLogic) AddDeck(req *types.DeckAddReq) (resp *types.CommonResply, err error) {
	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	de := &model.Deck{
		FloderId: req.FloderID,
		Name:     req.Name,
		Cardnum:  0,
	}
	err = model.DeckNsp.AddDeck(l.svcCtx.DbEngine, model.DeckNsp.PersonTableName(strconv.FormatInt(id, 10)), de)
	if err != nil {
		return nil, err
	}

	return &types.CommonResply{
		Code:    251,
		Message: "ok",
	}, nil
}
