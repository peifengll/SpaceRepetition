package repository

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAdmin(ctx context.Context, id int) (*model.Admin, error)
	Create(ctx context.Context, admin *model.Admin) error
	Update(ctx context.Context, admin *model.Admin) error
	GetAdminByUserName(ctx context.Context, username string) (*model.Admin, error)
}

func NewAdminRepository(
	repository *Repository,
) AdminRepository {
	return &adminRepository{
		Repository: repository,
	}
}

type adminRepository struct {
	*Repository
}

func (r *adminRepository) GetAdmin(ctx context.Context, id int) (*model.Admin, error) {
	var admin model.Admin
	err := r.DB(ctx).First(&admin, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, err
}
func (r *adminRepository) GetAdminByUserName(ctx context.Context, username string) (*model.Admin, error) {
	var admin model.Admin
	sql := `
	select id,
		   username,
		   password,
		   email,
		   name,
		   phone,
		   privileges
	from admins
	where username = ?
	`
	err := r.DB(ctx).Raw(sql, username).First(&admin).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, nil
}

func (r *adminRepository) Create(ctx context.Context, admin *model.Admin) error {
	if err := r.DB(ctx).Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) Update(ctx context.Context, admin *model.Admin) error {
	if err := r.DB(ctx).Save(admin).Error; err != nil {
		return err
	}
	return nil
}
