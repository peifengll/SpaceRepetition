package decks

import (
	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"
	"SpaceRepetition/src/internal/model"
	"context"
	"encoding/json"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllFloderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllFloderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllFloderLogic {
	return &GetAllFloderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllFloderLogic) GetAllFloder(req *types.FloderAllReq) (resp *types.FloderAllResp, err error) {
	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	floderList, err := model.FloderNsp.GetAllFloder(l.svcCtx.DbEngine, model.FloderNsp.PersonTableName(strconv.FormatInt(id, 10)))
	if err != nil {
		return nil, err
	}
	flresp := make([]types.Floderinfo, len(floderList))
	for i := 0; i < len(floderList); i++ {
		flresp[i].Name = floderList[i].Name
		flresp[i].DeckNum = floderList[i].DeckNum
		flresp[i].Id = int(floderList[i].ID)
	}
	//

	return &types.FloderAllResp{
		Data: flresp,
	}, nil
}
