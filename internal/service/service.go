package service

import (
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"github.com/peifengll/SpaceRepetition/pkg/helper/sid"
	"github.com/peifengll/SpaceRepetition/pkg/jwt"
	"github.com/peifengll/SpaceRepetition/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
