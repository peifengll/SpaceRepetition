package repository

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type KnowledgeRepository interface {
	FirstById(id int64) (*model.Knowledge, error)
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

// 返回的是一张待复习卡片，record相关的参数是要的，
// 然后卡片本身的内容
func (r *knowledgeRepository) GetAllReviewCard() {
	r.db.Raw("")
}
