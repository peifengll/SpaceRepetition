package service

import (
	"context"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type AdminService interface {
	GetAdmin(ctx context.Context, id int) (*model.Admin, error)
	Register(ctx context.Context, req *v1.AdminRegisterReq, alter int) error
	Login(ctx context.Context, req *v1.LoginAdminReq) (string, error)
}

func NewAdminService(service *Service, adminRepository repository.AdminRepository) AdminService {
	return &adminService{
		Service:         service,
		adminRepository: adminRepository,
	}
}

type adminService struct {
	*Service
	adminRepository repository.AdminRepository
}

func (s *adminService) GetAdmin(ctx context.Context, id int) (*model.Admin, error) {
	return s.adminRepository.GetAdmin(ctx, id)
}

func (s *adminService) Register(ctx context.Context, req *v1.AdminRegisterReq, alter int) error {
	// check Privileges
	admin, err := s.GetAdmin(ctx, alter)
	if err != nil {
		return err
	}
	if admin == nil {
		return v1.ErrUnauthorized
	}
	switch req.Privileges {
	case 1:
		if admin.Privileges == 0 || admin.Privileges == 1 {
		} else {
			return v1.ErrPrivileges
		}
	case 2:
		if admin.Privileges >= 2 {
			return v1.ErrPrivileges
		}
	default:
		return v1.ErrPrivileges
	}
	// check username
	if user, err := s.adminRepository.GetAdminByUserName(ctx, req.Username); err == nil && user != nil {
		return v1.ErrUsernameAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	info := &model.Admin{
		Username:   req.Username,
		Password:   string(hashedPassword),
		Email:      req.Email,
		Name:       req.Name,
		Phone:      req.Phone,
		Privileges: req.Privileges,
	}

	return s.adminRepository.Create(ctx, info)
}

func (s *adminService) Login(ctx context.Context, req *v1.LoginAdminReq) (string, error) {
	user, err := s.adminRepository.GetAdminByUserName(ctx, req.Username)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(strconv.FormatInt(user.ID, 10), time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}
