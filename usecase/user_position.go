package usecase

import (
	"time"

	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

type PositionUseCase interface {
	Post(latitude int64, longitude int64) *model.UserPosition
}

type positionUseCase struct {
	userPositionRepo repository.UserPosition
}

func NewPositionUseCase(repo repository.UserPosition) PositionUseCase {
	return &positionUseCase{
		userPositionRepo: repo,
	}
}

func (u *positionUseCase) Post(latitude int64, longitude int64) *model.UserPosition {
	userPositionModel := &model.UserPosition{
		Latitude:  latitude,
		Longitude: longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	u.userPositionRepo.Create(userPositionModel)

	return userPositionModel
}
