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

type AddFloderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFloderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFloderLogic {
	return &AddFloderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFloderLogic) AddFloder(req *types.FloderAddReq) (resp *types.CommonResply, err error) {
	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	err = model.FloderNsp.CreateFloder(l.svcCtx.DbEngine, model.FloderNsp.PersonTableName(strconv.FormatInt(id, 10)), req.Name)
	if err != nil {
		return nil, err
	}
	// 新建一个文件夹
	return &types.CommonResply{
		Code:    250,
		Message: "添加成功",
	}, nil
}
