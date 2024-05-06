package repository

import (
	"context"
	"errors"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/global_"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	DelFsrsParmKey(userid string) error
	Limit10UserIds() ([]v1.CalDeckNum, error)
	LearningToDayNum() (int64, error)
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DelFsrsParmKey(userid string) error {
	return r.rdb.Del(context.Background(), global_.GetFsrsParmsKey(userid)).Err()
}

func (r *userRepository) Limit10UserIds() ([]v1.CalDeckNum, error) {
	var datas []v1.CalDeckNum
	sql := `
	select count(*) as num, user_id
	from decks
	group by user_id
	order by num desc 
	limit 10
	`
	err := r.db.Raw(sql).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func (r *userRepository) LearningToDayNum() (int64, error) {
	result, err := r.rdb.SCard(context.Background(), global_.GetReviewersTodayKey(global_.Schedule)).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
