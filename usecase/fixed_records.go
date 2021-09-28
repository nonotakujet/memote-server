package usecase

import (
	"context"

	"github.com/DeNA/aelog"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

type FixedRecordsUseCase interface {
	GetByRecordId(ctx context.Context, recordId string) (*model.UserFixedRecord, error)
	GetAllByPictureFecthedFlag(ctx context.Context, isPictureFetched bool) ([]*model.UserFixedRecord, error)
	Update(ctx context.Context, recordId string, userFixedRecordModel *model.UserFixedRecord) (*model.UserFixedRecord, error)
}

type fixedRecordsUseCase struct {
	userFixedRecordRepo repository.UserFixedRecord
}

func NewFixedRecordUseCase(userFixedRecordRepo repository.UserFixedRecord) FixedRecordsUseCase {
	return &fixedRecordsUseCase{
		userFixedRecordRepo: userFixedRecordRepo,
	}
}

func (u *fixedRecordsUseCase) GetByRecordId(ctx context.Context, recordId string) (*model.UserFixedRecord, error) {
	uid, err := model.UserFromContext(ctx)
	if err != nil {
		aelog.Errorf(ctx, "Failed get uid from context: %v", err)
		return nil, err
	}

	fixedRecordModel, err := u.userFixedRecordRepo.GetById(ctx, uid, recordId)

	return fixedRecordModel, err
}

func (u *fixedRecordsUseCase) GetAllByPictureFecthedFlag(ctx context.Context, isPictureFetched bool) ([]*model.UserFixedRecord, error) {
	uid, err := model.UserFromContext(ctx)
	if err != nil {
		aelog.Errorf(ctx, "Failed get uid from context: %v", err)
		return nil, err
	}

	fixedRecordModels, err := u.userFixedRecordRepo.GetAllByPictureFecthedFlag(ctx, uid, isPictureFetched)

	return fixedRecordModels, err
}

func (u *fixedRecordsUseCase) Update(ctx context.Context, recordId string, userFixedRecordModel *model.UserFixedRecord) (*model.UserFixedRecord, error) {
	uid, err := model.UserFromContext(ctx)
	if err != nil {
		aelog.Errorf(ctx, "Failed get uid from context: %v", err)
		return nil, err
	}

	fixedRecordModel, err := u.userFixedRecordRepo.GetById(ctx, uid, recordId)
	if err != nil {
		aelog.Errorf(ctx, "Failed get fixed record by id: %+v", err)
		return nil, err
	}
	if fixedRecordModel == nil {
		aelog.Errorf(ctx, "not exists fixed record: %s", recordId)
		return nil, err
	}

	// update.
	fixedRecordModel, err = u.userFixedRecordRepo.Update(ctx, uid, recordId, userFixedRecordModel)

	return fixedRecordModel, err
}
