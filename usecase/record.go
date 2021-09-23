package usecase

import (
	"context"
	"log"

	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
	"github.com/nonotakujet/memote-server/domain/viewmodel"
	"github.com/thoas/go-funk"
)

type RecordUseCase interface {
	Post(ctx context.Context, recordViewModel *viewmodel.RecordViewModel) *model.UserRecord
}

type recordUseCase struct {
	userRecordRepo repository.UserRecord
}

func NewRecordUseCase(repo repository.UserRecord) RecordUseCase {
	return &recordUseCase{
		userRecordRepo: repo,
	}
}

func (u *recordUseCase) Post(ctx context.Context, recordViewModel *viewmodel.RecordViewModel) *model.UserRecord {
	uid, err := model.UserFromContext(ctx)
	if err != nil {
		log.Fatalf("Failed get uid from context: %v", err)
	}

	userRecordModel := &model.UserRecord{
		Id: recordViewModel.Id,
		Locations: funk.Map(recordViewModel.Locations, func(location viewmodel.LocationViewModel) model.UserLocation {
			return model.UserLocation{
				Latitude:  location.Lat,
				Longitude: location.Long,
				Time:      location.Time,
			}
		}).([]model.UserLocation),
		CreatedAt: recordViewModel.CreatedAt,
	}
	u.userRecordRepo.Create(ctx, uid, userRecordModel)

	return userRecordModel
}
