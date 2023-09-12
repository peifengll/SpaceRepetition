package user

import (
	"SpaceRepetition/src/internal/model"
	"SpaceRepetition/src/internal/service"
	"context"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.CommonResply, err error) {
	// 先看有没有
	exist, err := model.UserNsp.FindByUsername(l.svcCtx.DbEngine, req.UserName)
	if err != nil {
		return nil, err
	}
	if exist {
		return &types.CommonResply{
			Code: 502, Message: "用户名重复",
		}, nil
	}
	u := &model.User{}
	u.UserName = req.UserName
	u.PassWord = req.PassWord
	u.Email = req.Email
	u.Gender = req.Gender
	u.Age = req.Age

	err = model.UserNsp.AddUser(l.svcCtx.DbEngine, u)
	id, err := model.UserNsp.FindIdByUsername(l.svcCtx.DbEngine, u.UserName)
	if err != nil {
		return nil, err
	}
	// 帮这个用户把所有表注册好
	// 注册所有的表
	service.HelpUserRegisterStroge(l.svcCtx.DbEngine, id)

	if err != nil {
		return &types.CommonResply{
			Code: 503, Message: "注册失败，稍后重试",
		}, err
	}
	return &types.CommonResply{
		Code: 205, Message: "欢迎新用户~",
	}, nil
}
