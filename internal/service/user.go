package service

import (
	"context"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
	UpdateUserInfos(ctx context.Context, userId string, req *v1.UserUpdateReq) error
	GetTenUserInfos() ([]v1.UserWithNum, error)
	OnLearningErNum() (int64, error)
}

func NewUserService(que *query.Query, service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		query:    que,

		Service: service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	query    *query.Query
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check username
	if user, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil && user != nil {
		return v1.ErrEmailAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user := &model.User{
		UserID:   userId,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(user.UserID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.UserID,
		Nickname: user.Nickname,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUserInfos(ctx context.Context, userId string, req *v1.UserUpdateReq) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}
	if req.Gender != 0 {
		user.Gender = req.Gender
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Username != "" {
		user.UserID = req.Username
	}
	if req.Introduction != "" {
		user.Introduction = req.Introduction
	}
	if req.MaxInterval != 0 {
		user.MaxInterval = req.MaxInterval
	}
	if req.Weights != "" {
		user.Weights = req.Weights
	}
	if req.RequestRetention != 0 {
		user.RequestRetention = req.RequestRetention
	}
	if req.Weights != "" || req.RequestRetention != 0 || req.MaxInterval != 0 {
		err = s.userRepo.DelFsrsParmKey(userId)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetTenUserInfos() ([]v1.UserWithNum, error) {
	numids, err := s.userRepo.Limit10UserIds()
	if err != nil {
		return nil, err
	}
	dc := make([]v1.UserWithNum, 0)
	u := s.query.User
	for _, item := range numids {
		first, err := u.Where(u.UserID.Eq(item.UserId)).First()
		if err != nil {
			return nil, err
		}
		dc = append(dc, v1.UserWithNum{
			User: first,
			Num:  item.Num,
		})
	}
	return dc, nil
}

func (s *userService) OnLearningErNum() (int64, error) {
	return s.userRepo.LearningToDayNum()
}
