package knowledge

import (
	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"
	"SpaceRepetition/src/internal/model"
	"SpaceRepetition/src/utils/go-fsrs-main"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewLogic {
	return &ReviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewLogic) Review(req *types.ReviewReq) (resp *types.ReviewResp, err error) {
	// 首先就查出旧的记录
	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	oldRe, err := model.RecordNsp.RepealRecordByKid(
		l.svcCtx.DbEngine,
		model.RecordNsp.PersonTableName(strconv.FormatInt(id, 10)),
		int(req.KnowledgeId))
	if err != nil {
		return nil, err
	}
	p := fsrs.DefaultParam()
	oldCard := fsrs.Card{
		Due:           oldRe.Due,
		Stability:     oldRe.Stability,
		Difficulty:    oldRe.Difficulty,
		ElapsedDays:   oldRe.ElapsedDays,
		ScheduledDays: oldRe.ScheduledDays,
		Reps:          oldRe.Reps,
		Lapses:        oldRe.Lapses,
		State:         oldRe.State,
		LastReview:    oldRe.LastReview,
	}
	now := time.Now()
	// 复习调度信息，供选择
	schedulingCards := p.Repeat(oldCard, now)
	// 输出复习信息
	fmt.Println("=============================")
	fmt.Printf("%+v", schedulingCards[fsrs.Rating(req.Op)])
	fmt.Println("=============================")

	// 原本记录失效，
	err = model.RecordNsp.UpdateOnById(
		l.svcCtx.DbEngine,
		model.RecordNsp.PersonTableName(strconv.FormatInt(id, 10)),
		int(oldRe.ID),
	)
	if err != nil {
		return nil, err
	}
	newRe := &model.Record{
		KnowledgeId:   oldRe.ID,
		Due:           schedulingCards[fsrs.Rating(req.Op)].Card.Due,
		Stability:     schedulingCards[fsrs.Rating(req.Op)].Card.Stability,
		Difficulty:    schedulingCards[fsrs.Rating(req.Op)].Card.Difficulty,
		ElapsedDays:   schedulingCards[fsrs.Rating(req.Op)].Card.ElapsedDays,
		ScheduledDays: schedulingCards[fsrs.Rating(req.Op)].Card.ScheduledDays,
		Reps:          schedulingCards[fsrs.Rating(req.Op)].Card.Reps,
		Lapses:        schedulingCards[fsrs.Rating(req.Op)].Card.Lapses,
		State:         schedulingCards[fsrs.Rating(req.Op)].Card.State,
		On:            false,
		LastOp:        req.Op,
		LastReview:    schedulingCards[fsrs.Rating(req.Op)].Card.LastReview,
	}
	//留下最新的记录
	err = model.RecordNsp.AddNewRecord(
		l.svcCtx.DbEngine,
		model.RecordNsp.PersonTableName(strconv.FormatInt(id, 10)),
		newRe,
	)
	if err != nil {
		return nil, err
	}
	// 到这里就看要不要再复习了,
	today := time.Date(now.Year(), now.Month(), now.Day(), 24, 0, 0, 0, now.Location())
	if newRe.Due.Before(today) {
		// 就还要复习，
		return &types.ReviewResp{
			CommonResply: types.CommonResply{
				Code:    207,
				Message: "正常",
			},
			Data: types.KnowledgeInfo{
				Font: "为什么没有放id属性",
				Back: "这里返回的信息没有啥用，后期再改",
			},
		}, nil
	}
	return &types.ReviewResp{
		CommonResply: types.CommonResply{
			Code:    207,
			Message: "正常",
		},
		Data: types.KnowledgeInfo{},
	}, nil
}
