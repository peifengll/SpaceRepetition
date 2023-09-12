package deckopen

import (
	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"
	"SpaceRepetition/src/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type OpenDeckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenDeckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenDeckLogic {
	return &OpenDeckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenDeckLogic) OpenDeck(req *types.DeckAllReq) (resp *types.DeckAllResp, err error) {

	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	// 先查有几个章节属于这个表
	sectionList, err := model.SectionNsp.FindSectionListByDeckId(l.svcCtx.DbEngine,
		model.SectionNsp.PersonTableName(strconv.FormatInt(id, 10)),
		req.DeckId)
	if err != nil {
		return nil, err
	}
	dic := make(map[string][]types.Knowledgeinfo)
	//再查  有几个知识点属于这个章节
	for _, section := range sectionList {
		newname := fmt.Sprintf("%d_%s", section.ID, section.Name)
		klist, err := model.KnowledgeNsp.FindKnowLedgeListBySectionId(l.svcCtx.DbEngine,
			model.KnowledgeNsp.PersonTableName(strconv.FormatInt(id, 10)),
			int(section.ID),
		)
		if err != nil {
			return nil, err
		}
		dic[newname] = make([]types.Knowledgeinfo, len(klist))
		for i := 0; i < len(klist); i++ {
			dic[newname][i] = types.Knowledgeinfo{
				Id:     int(klist[i].ID),
				Font:   klist[i].Font,
				Back:   klist[i].Back,
				Typeof: klist[i].Typeof,
				Sectioninfo: types.Sectioninfo{
					Id:   int(section.ID),
					Name: section.Name,
				},
			}
		}
	}
	return &types.DeckAllResp{
		CommonResply: types.CommonResply{
			Code:    209,
			Message: "没啥大问题",
		},
		Data: dic,
	}, nil
}
