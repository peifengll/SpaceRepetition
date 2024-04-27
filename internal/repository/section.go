package repository

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type SectionRepository interface {
	FirstById(id int64) (*model.Section, error)
	GetOnLearnNumberBySectionId(secid int64) (int64, error)
}

func NewSectionRepository(repository *Repository) SectionRepository {
	return &sectionRepository{
		Repository: repository,
	}
}

type sectionRepository struct {
	*Repository
}

func (r *sectionRepository) FirstById(id int64) (*model.Section, error) {
	var section model.Section
	// TODO: query db
	return &section, nil
}

// 返回这个下边的正在学习卡片的数量
func (r *sectionRepository) GetOnLearnNumberBySectionId(secid int64) (int64, error) {
	var count int64
	err := r.db.Model(&model.Knowledge{}).Where("sectionid = ?  and onlearning = 1", secid).Count(&count).Error
	return count, err
}
