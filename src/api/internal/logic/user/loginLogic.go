package user

import (
	"SpaceRepetition/src/internal/model"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResply, err error) {
	user, err := model.UserNsp.FindByUsernameAndPassword(l.svcCtx.DbEngine, req.UserName, req.PassWord)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if user == nil {
		return &types.LoginResply{
			CommonResply: types.CommonResply{
				Code: 404, Message: "用户名不存在或密码不正确",
			},
		}, nil
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(user.ID))

	return &types.LoginResply{
		CommonResply: types.CommonResply{
			Code: 203, Message: "一切正常",
		},
		Data: types.Userinfo{
			Gender:   user.Gender,
			Age:      user.Age,
			UserName: user.UserName,
			Email:    user.Email,
		},
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
