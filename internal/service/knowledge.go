package service

import (
	"fmt"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"github.com/peifengll/SpaceRepetition/pkg/helper/fsrs"
	"log"
	"time"
)

type KnowledgeService interface {
	GetKnowledge(id int64) (*model.Knowledge, error)
	AddCard(v *v1.CardRequest, userid string) error
	GetKnowledgeById(id int64) (*v1.CardResp, error)
	UpdateCard(v *v1.CardRequest) error
	DeleteCard(id int64) error
	SearchCards(content string, userId string) ([]*v1.CardResp, error)
	ChooseToReview(ids int64, userId string) error
	GetAllReview(id string) ([]v1.DeckCardReviewResp, error)
	CheckReview(id int64) (bool, error)
	ReviewOp(t *v1.CardReviewOptReq, userid string) (int64, error)
	GetReviewStatics(userId string) ([]*model.DayReviewStatistic, error)
}

func NewKnowledgeService(que *query.Query, service *Service, knowledgeRepository repository.KnowledgeRepository) KnowledgeService {
	return &knowledgeService{
		Service:             service,
		query:               que,
		knowledgeRepository: knowledgeRepository,
	}
}

type knowledgeService struct {
	*Service
	query               *query.Query
	knowledgeRepository repository.KnowledgeRepository
}

func (s *knowledgeService) GetKnowledge(id int64) (*model.Knowledge, error) {
	return s.knowledgeRepository.FirstById(id)
}

func (s *knowledgeService) AddCard(v *v1.CardRequest, userid string) error {
	cc := &model.Knowledge{
		Font:       v.Font,
		Originfont: v.Originfont,
		Back:       v.Back,
		Typeof:     v.Typeof,
		Deckid:     v.Deckid,
		Sectionid:  v.Sectionid,
		UserID:     userid,
	}
	tx := s.query.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	err := tx.Knowledge.Create(cc)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Deck.Where(tx.Deck.ID.Eq(cc.Deckid)).UpdateSimple(tx.Deck.Cardnum.Add(1))
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Section.Where(tx.Section.ID.Eq(cc.Sectionid)).UpdateSimple(tx.Section.CardNum.Add(1))
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *knowledgeService) GetKnowledgeById(id int64) (*v1.CardResp, error) {
	first, err := s.query.Knowledge.Where(s.query.Knowledge.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	res := &v1.CardResp{
		ID:         first.ID,
		Font:       first.Back,
		Originfont: first.Originfont,
		Back:       first.Back,
		Onlearning: first.Onlearning,
		Typeof:     first.Typeof,
		Deckid:     first.Deckid,
		Skilled:    first.Skilled,
		Sectionid:  first.Sectionid,
	}
	return res, nil
}

func (s *knowledgeService) UpdateCard(v *v1.CardRequest) error {
	k := &model.Knowledge{
		Font:       v.Font,
		Originfont: v.Originfont,
		Back:       v.Back,
		Typeof:     v.Typeof,
		Deckid:     v.Deckid,
		Sectionid:  v.Sectionid,
	}
	_, err := s.query.Knowledge.Where(s.query.Knowledge.ID.Eq(v.ID)).Updates(k)
	if err != nil {
		return err
	}
	return nil
}

func (s *knowledgeService) DeleteCard(id int64) error {
	deck, err := s.query.Knowledge.Where(s.query.Knowledge.ID.Eq(id)).First()
	if err != nil {
		return err
	}
	tx := s.query.Begin()
	_, err = tx.Section.Where(tx.Section.ID.Eq(deck.Sectionid)).UpdateSimple(tx.Section.CardNum.Sub(1))
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Deck.Where(tx.Deck.ID.Eq(deck.Deckid)).UpdateSimple(tx.Deck.Cardnum.Sub(1))
	if err != nil {
		tx.Rollback()
		return err
	}
	if deck.Onlearning == 1 {
		_, err = tx.Deck.Where(tx.Deck.ID.Eq(deck.Deckid)).UpdateSimple(tx.Deck.Learnnumber.Sub(1))
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// 删除记录
	_, err = tx.Record.Where(tx.Record.KnowledgeID.Eq(id)).Delete()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = s.query.Knowledge.Where(s.query.Knowledge.ID.Eq(id)).Delete()
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// 这个就默认是搜索这个用户所有的了
func (s *knowledgeService) SearchCards(content string, userId string) ([]*v1.CardResp, error) {
	// todo  待测试这个能不能用
	// 直接插origin，查属于这个牌组的
	con := fmt.Sprintf("%%%s%%", content)
	knows, err := s.query.Knowledge.Where(s.query.Knowledge.UserID.Eq(userId), s.query.Knowledge.Originfont.Like(con)).Find()
	if err != nil {
		return nil, err
	}
	cards := make([]*v1.CardResp, len(knows))
	for i := 0; i < len(cards); i++ {
		cards[i] = &v1.CardResp{
			ID:         knows[i].ID,
			Font:       knows[i].Font,
			Originfont: knows[i].Originfont,
			Back:       knows[i].Back,
			Onlearning: knows[i].Onlearning,
			Typeof:     knows[i].Typeof,
			Deckid:     knows[i].Deckid,
			Skilled:    knows[i].Skilled,
			Sectionid:  knows[i].Sectionid,
		}
	}
	return cards, nil
}
func (s *knowledgeService) CheckReview(id int64) (bool, error) {
	kno := s.query.Knowledge
	first, err := kno.Where(kno.ID.Eq(id)).First()
	if err != nil {
		return false, err
	}
	return first.Onlearning == 1, nil
}

// 先获取id列表
// 将所有的id对应的card的onlearn更新为在学习，
// 向复习记录表中插入一些新的record，也就是用来进行复习调度用的数据，应开始复习的时间就设置为今天
//
// 对了，应该不需要用事务的方式去做...?因为只要加入了复习，其实也不影响，。。。。
// 不对，还是要以事务的方式去做，因为 设计两个表，一个更新了，另一个却没更新，这是不得行的
func (s *knowledgeService) ChooseToReview(id int64, userId string) error {
	//	这里要用事务来做，全部都开始复习
	q := s.query
	q.Transaction(func(tx *query.Query) error {
		po, err := tx.Knowledge.Where(tx.Knowledge.ID.Eq(id)).First()
		if err != nil {
			return err
		}
		// 第一步 更新为onlearning
		_, err = tx.Knowledge.Where(tx.Knowledge.ID.Eq(id)).Update(tx.Knowledge.Onlearning, 1)
		if err != nil {
			return err
		}

		_, err = tx.Deck.Where(tx.Deck.ID.Eq(po.Deckid)).UpdateSimple(tx.Deck.Learnnumber.Add(1))
		if err != nil {
			return err
		}
		//_, err := tx.Knowledge.Where(tx.Knowledge.ID.Eq(id)).Update(tx.Knowledge.Onlearning, 1)
		// 第二步 增加记录到复习里边去
		records := &model.Record{
			KnowledgeID:   id,
			Due:           time.Now(),
			Stability:     0,
			Difficulty:    0,
			ElapsedDays:   0,
			ScheduledDays: 0,
			Reps:          0,
			Lapses:        0,
			State:         int64(fsrs.New),
			On:            1,
			LastReview:    time.Now(),
			UserID:        userId,
		}
		err = tx.Record.Create(records)
		if err != nil {
			return err
		}

		return nil
	})
	return nil
}

func (s *knowledgeService) GetAllReview(id string) ([]v1.DeckCardReviewResp, error) {
	reviews, err := s.knowledgeRepository.GetAllReviewCard(id)
	if err == nil && reviews == nil {
		return make([]v1.DeckCardReviewResp, 0), nil
	}
	return reviews, err
}

// 进行一次复习计算，算出时间间隔，并更新到数据库
func (s *knowledgeService) ReviewOp(t *v1.CardReviewOptReq, userid string) (int64, error) {
	// 每个用户都应该有自己的 回忆成功概率， 80 到 90 之间最好在
	var p fsrs.Parameters
	parms, b, err := s.knowledgeRepository.LoadParms(userid)
	log.Println(parms)
	log.Println(b)
	if err != nil {
		return 0, err
	}
	p = fsrs.DefaultParam()
	if parms != nil {
		if b {
			p.W = parms.W
		}
		if parms.MaximumInterval != 0 {
			p.MaximumInterval = parms.MaximumInterval
		}
		if parms.RequestRetention != 0 {
			p.RequestRetention = parms.RequestRetention
		}
	}
	// 先就用这个默认的
	q := s.query
	record, err := q.Record.Where(q.Record.KnowledgeID.Eq(t.ID), q.Record.On.Eq(1)).First()
	//fmt.Printf("%#v\n", record)
	if err != nil {
		return 0, err
	}
	q.Knowledge.Where(q.Knowledge.ID.Eq(t.ID))
	card := fsrs.Card{
		Due:           record.Due,
		Stability:     record.Stability,
		Difficulty:    record.Difficulty,
		ElapsedDays:   uint64(record.ElapsedDays),
		ScheduledDays: uint64(record.ScheduledDays),
		Reps:          uint64(record.Reps),
		Lapses:        uint64(record.Lapses),
		State:         fsrs.State(record.State),
		LastReview:    record.LastReview,
	}
	schedulingCards := p.Repeat(card, time.Now())
	// 新纪录
	card = schedulingCards[t.Opt].Card
	newRecord := model.Record{
		KnowledgeID:   t.ID,
		Due:           card.Due,
		Stability:     card.Stability,
		Difficulty:    card.Difficulty,
		ElapsedDays:   int64(card.ElapsedDays),
		ScheduledDays: int64(card.ScheduledDays),
		Reps:          int64(card.Reps),
		Lapses:        int64(card.Lapses),
		State:         int64(card.State),
		On:            1,
		Lastop:        int64(t.Opt),
		LastReview:    time.Now(),
		UserID:        userid,
	}
	err = q.Transaction(func(tx *query.Query) error {
		// 将到期的这条数据的评分给他
		_, err = tx.Record.Where(q.Record.ID.Eq(t.RecordID)).Update(q.Record.Rate, t.Opt)
		if err != nil {
			return err
		}
		// 新的record插入进去，旧的record进行抛掉
		_, err = tx.Record.Where(q.Record.ID.Eq(record.ID)).Update(q.Record.On, 0)
		if err != nil {
			return err
		}
		err = tx.Record.Create(&newRecord)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	err = s.knowledgeRepository.AddRecordToRedis(userid, t.ID, t.Opt)
	if err != nil {
		return 0, err
	}
	if isToday(newRecord.Due) {
		return newRecord.ID, nil
	}
	return 0, nil
}

func isToday(t time.Time) bool {

	// 获取当前时间的年、月、日
	currentYear, currentMonth, currentDay := time.Now().Date()

	// 获取给定时间的年、月、日
	year, month, day := t.Date()

	// 比较日期部分是否相等
	return year == currentYear && month == currentMonth && day == currentDay
}

func (s *knowledgeService) GetReviewStatics(userId string) ([]*model.DayReviewStatistic, error) {
	finds, err := s.query.DayReviewStatistic.Where(s.query.DayReviewStatistic.UserID.Eq(userId)).Find()
	if err != nil {
		return nil, err
	}
	return finds, nil
}
