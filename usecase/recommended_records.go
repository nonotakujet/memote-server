package usecase

import (
	"context"
	"math/rand"

	"github.com/DeNA/aelog"
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

	fixedRecordModels, err := u.userFixedRecordRepo.GetAll(ctx, uid)
	if len(fixedRecordModels) == 0 {
		return []*model.UserFixedRecord{}, nil
	}

	// 重み付き抽選を行う
	// 重みはレコードに設定されている写真数とする.
	totalWeight := 0
	var picked *model.UserFixedRecord

	for i := 0; i < len(fixedRecordModels); i++ {
		for _, location := range fixedRecordModels[i].Locations {
			totalWeight += len(location.Pictures)
		}
	}

	rnd := rand.Intn(100000) % totalWeight

	for i := 0; i < len(fixedRecordModels); i++ {
		locationNum := len(fixedRecordModels[i].Locations)
		if rnd < locationNum {
			// Hit.
			picked = fixedRecordModels[i]
			break
		}
		rnd -= locationNum
	}

	if picked == nil {
		return []*model.UserFixedRecord{}, nil
	}

	return []*model.UserFixedRecord{picked}, err
}
