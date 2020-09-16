package usecase

import (
	"context"
	"time"

	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

type PositionUseCase interface {
	Post(ctx context.Context, atitude int64, longitude int64) *model.UserPosition
}

type positionUseCase struct {
	userPositionRepo repository.UserPosition
}

func NewPositionUseCase(repo repository.UserPosition) PositionUseCase {
	return &positionUseCase{
		userPositionRepo: repo,
	}
}

func (u *positionUseCase) Post(ctx context.Context, latitude int64, longitude int64) *model.UserPosition {
	userPositionModel := &model.UserPosition{
		Latitude:  latitude,
		Longitude: longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	u.userPositionRepo.Create(ctx, userPositionModel)

	return userPositionModel
}
