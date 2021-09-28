package usecase

import (
	"context"
	"math/rand"
	"sort"

	"github.com/DeNA/aelog"
	"github.com/mmcloughlin/geohash"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

type RecommendedRecordsUseCase interface {
	Get(ctx context.Context, lat float64, lng float64) ([]*model.UserFixedRecord, error)
}

type recommendedRecordsUseCase struct {
	userFixedRecordRepo repository.UserFixedRecord
	userLocationRepo    repository.UserLocation
}

func NewRecommendedRecordUseCase(userFixedRecordRepo repository.UserFixedRecord, userLocationRepo repository.UserLocation) RecommendedRecordsUseCase {
	return &recommendedRecordsUseCase{
		userFixedRecordRepo: userFixedRecordRepo,
		userLocationRepo:    userLocationRepo,
	}
}

func (u *recommendedRecordsUseCase) Get(ctx context.Context, lat float64, lng float64) ([]*model.UserFixedRecord, error) {
	uid, err := model.UserFromContext(ctx)
	if err != nil {
		aelog.Errorf(ctx, "Failed get uid from context: %v", err)
		return nil, err
	}

	// lat,lngからgeohashを生成.
	geohash := geohash.EncodeWithPrecision(lat, lng, 10)

	aelog.Infof(ctx, "geohash -> %s", geohash)

	// geohashからlocationを取得.
	locationModels, err := u.userLocationRepo.GetNearBy(ctx, uid, geohash)

	if err != nil {
		aelog.Errorf(ctx, "error: %v\n", err)
		return nil, err
	}

	if len(locationModels) == 0 {
		return []*model.UserFixedRecord{}, nil
	}

	// recordの少ない順にソート.
	sort.Slice(locationModels, func(i, j int) bool { return len(locationModels[i].RecordIds) < len(locationModels[j].RecordIds) })

	// 先頭のrecordIdsを取得.
	recordIds := locationModels[0].RecordIds

	// recordIdsをrandomして先頭を返す.
	for i := range recordIds {
		j := rand.Intn(i + 1)
		recordIds[i], recordIds[j] = recordIds[j], recordIds[i]
	}
	recordId := recordIds[0]

	fixedRecordModel, err := u.userFixedRecordRepo.GetById(ctx, uid, recordId)

	return []*model.UserFixedRecord{fixedRecordModel}, err
}
