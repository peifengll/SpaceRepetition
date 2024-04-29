package repository

import (
	"fmt"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type KnowledgeRepository interface {
	FirstById(id int64) (*model.Knowledge, error)
	GetAllReviewCard(userid string) ([]v1.DeckCardReviewResp, error)
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
