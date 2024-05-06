package repository

import (
	"context"
	"fmt"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/global_"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/pkg/helper/convert"
	"github.com/peifengll/SpaceRepetition/pkg/helper/fsrs"
	"strconv"
)

type KnowledgeRepository interface {
	FirstById(id int64) (*model.Knowledge, error)
	GetAllReviewCard(userid string) ([]v1.DeckCardReviewResp, error)
	AddRecordToRedis(userid string, cardId int64, rate fsrs.Rating) error
	LoadParms(userid string) (*fsrs.Parameters, bool, error)
}

func NewKnowledgeRepository(repository *Repository) KnowledgeRepository {
	return &knowledgeRepository{
		Repository: repository,
	}
}

type knowledgeRepository struct {
	*Repository
}

func (r *knowledgeRepository) FirstById(id int64) (*model.Knowledge, error) {
	var knowledge model.Knowledge
	// TODO: query db
	return &knowledge, nil
}

// 只要要复习的卡片，卡片本身的内容，和卡片属于的牌组
func (r *knowledgeRepository) GetAllReviewCard(userid string) ([]v1.DeckCardReviewResp, error) {
	decks := &[]model.Deck{}
	err := r.db.Model(&model.Deck{}).Where("user_id = ?", userid).Find(decks).Error
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	sql := `
	SELECT k.id,
		   r.id AS record_id,
		   k.font,
		   k.originfont,
		   k.back,
		   k.deckid,
		   k.typeof
	FROM knowledge k
			 LEFT JOIN record r ON k.id = r.knowledge_id
	WHERE DATE(r.Due) <= CURDATE()
	  and r.on = 1
	  and k.user_id = ?
	ORDER BY k.deckid
`
	var cards []*v1.CardReviewResp
	err = r.db.Raw(sql, userid).Find(&cards).Error
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	if len(cards) == 0 {
		// 今天学完了
		return nil, nil
	}
	reviews := make([]v1.DeckCardReviewResp, 0)
	i := 0
	fmt.Println("+++++++++++++++++++++++")
	for i := 0; i < len(cards); i++ {
		fmt.Printf("%#v\n", cards[i])
	}
	fmt.Println("+++++++++++++++++++++++")

	for i < len(cards) {
		if i+1 == len(cards) {
			c := model.Deck{}
			err = r.db.Model(&model.Deck{}).Select("name").Where("id=?", cards[i].DeckID).First(&c).Error
			if err != nil {
				r.logger.Warn(err.Error())
			}
			reviews = append(reviews, v1.DeckCardReviewResp{
				ID:    cards[i].DeckID,
				Name:  c.Name,
				Cards: cards[i:],
			})
			break
		}
		allIs := true
		for j := i + 1; j < len(cards); j++ {
			if cards[j].DeckID != cards[i].DeckID {
				allIs = false
				c := model.Deck{}
				err = r.db.Model(&model.Deck{}).Select("name").Where("id=?", cards[i].DeckID).First(&c).Error
				if err != nil {
					r.logger.Warn(err.Error())
				}
				reviews = append(reviews, v1.DeckCardReviewResp{
					ID:    cards[i].DeckID,
					Name:  c.Name,
					Cards: cards[i:j],
				})
				i = j
			}
		}
		if allIs {
			c := model.Deck{}
			err = r.db.Model(&model.Deck{}).Select("name").Where("id=?", cards[i].DeckID).First(&c).Error
			if err != nil {
				r.logger.Warn(err.Error())
			}
			reviews = append(reviews, v1.DeckCardReviewResp{
				ID:    cards[i].DeckID,
				Name:  c.Name,
				Cards: cards[i:],
			})
			break
		}
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("%#v\n", reviews)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	return reviews, nil
}

func (r *knowledgeRepository) AddRecordToRedis(userid string, cardId int64, rate fsrs.Rating) error {

	ctx := context.Background()
	// 加入复习用户
	err := r.rdb.SAdd(ctx, global_.GetReviewersTodayKey(global_.Schedule), userid).Err()
	if err != nil {
		return err
	}
	// 当前用户的复习卡片数
	err = r.rdb.SAdd(
		ctx,
		global_.GetReviewCardsSetKey(userid, global_.Schedule),
		cardId,
	).Err()
	if err != nil {
		return err
	}
	// 复习操作记录
	err = r.rdb.HIncrBy(ctx, global_.GetReviewOpNumberKey(userid, global_.Schedule), rate.String(), 1).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *knowledgeRepository) LoadParms(userid string) (*fsrs.Parameters, bool, error) {
	// 获取参数，要是没得就用default是
	ctx := context.Background()
	flag := r.rdb.Exists(ctx, global_.GetFsrsParmsKey(userid)).Val()
	parms := fsrs.Parameters{
		RequestRetention: 0,
		MaximumInterval:  0,
		W:                fsrs.Weights{},
	}
	if flag == 0 {
		var po model.User
		err := r.db.Model(&model.User{}).Where("user_id = ?", userid).
			Select("max_interval", "weights", "request_retention").
			First(&po).Error
		if err != nil {
			return nil, false, err
		}
		r.rdb.HSet(ctx, global_.GetFsrsParmsKey(userid),
			"max_interval", po.MaxInterval,
			"weights", po.Weights,
			"request_retention", po.RequestRetention)
	}
	// 存在 ，就去拿
	dic := r.rdb.HGetAll(ctx, global_.GetFsrsParmsKey(userid)).Val()
	float1, err := strconv.ParseFloat(dic["max_interval"], 10)
	if err != nil {
		return nil, false, err
	}
	parms.MaximumInterval = float1
	parms.RequestRetention, err = strconv.ParseFloat(dic["request_retention"], 10)
	if err != nil {
		return nil, false, err
	}
	if dic["weights"] == "" {
		return &parms, false, err
	} else {
		wes, err := convert.ParseStringToFloatSlice(dic["weights"])
		if err != nil {
			return nil, false, err
		}
		if len(wes) < 17 {
			return &parms, false, err
		}
		for ii := 0; ii < 17; ii++ {
			parms.W[ii] = wes[ii]
		}
		return &parms, true, nil
	}
}
