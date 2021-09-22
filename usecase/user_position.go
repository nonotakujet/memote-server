package usecase

import (
	"context"
	"log"
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
	uid, err := model.UserFromContext(ctx)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	userPositionModel := &model.UserPosition{
		Latitude:  latitude,
		Longitude: longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	u.userPositionRepo.Create(ctx, uid, userPositionModel)

	return userPositionModel
}
