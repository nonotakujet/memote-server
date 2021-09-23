package usecase

import (
	"context"
	"log"
	"time"

	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
	"github.com/nonotakujet/memote-server/domain/viewmodel"
)

type PositionUseCase interface {
	Post(ctx context.Context, locationViewModels []viewmodel.LocationViewModel) *model.UserPosition
}

type positionUseCase struct {
	userPositionRepo repository.UserPosition
}

func NewPositionUseCase(repo repository.UserPosition) PositionUseCase {
	return &positionUseCase{
		userPositionRepo: repo,
	}
}

func (u *positionUseCase) Post(ctx context.Context, locationViewModels []viewmodel.LocationViewModel) *model.UserPosition {

	uid, err := model.UserFromContext(ctx)
	if err != nil {
		log.Fatalf("Failed get uid from context: %v", err)
	}

	userPositionModel := &model.UserPosition{
		Latitude:  0,
		Longitude: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	u.userPositionRepo.Create(ctx, uid, userPositionModel)

	return userPositionModel
}
