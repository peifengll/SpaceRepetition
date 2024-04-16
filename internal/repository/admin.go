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
	DelById(ctx context.Context, id int64) error
	ShowAll(ctx context.Context) ([]model.Admin, error)
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

func (r *adminRepository) ShowAll(ctx context.Context) ([]model.Admin, error) {
	var res []model.Admin
	sql := `
	select id,
       username,
       email,
       name,
       phone,
       privileges
	from admins
	`
	if err := r.DB(ctx).Raw(sql).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *adminRepository) Update(ctx context.Context, admin *model.Admin) error {
	if err := r.DB(ctx).Updates(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) DelById(ctx context.Context, id int64) error {
	if err := r.DB(ctx).Delete(&model.Admin{}, id).Error; err != nil {
		return err
	}
	return nil
}
