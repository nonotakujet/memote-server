package persistence

import (
	"context"

	"github.com/DeNA/aelog"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

// UserFixedRecordRepository holds user position inteface
type UserFixedRecordRepository struct {
	db *DB
}

// NewUserRecordRepository new user position
func NewUserFixedRecordRepository(db *DB) repository.UserFixedRecord {
	newRepo := &UserFixedRecordRepository{
		db: db,
	}
	return newRepo
}

func (r *UserFixedRecordRepository) GetById(ctx context.Context, uid *model.UID, recordId string) (*model.UserFixedRecord, error) {
	ss, err := r.db.client.Collection("users").Doc(uid.ID).Collection("fixedRecords").Doc(recordId).Get(ctx)

	if err != nil {
		aelog.Errorf(ctx, "Failed userFixedRecord getById: %v", err)
		return nil, err
	}

	var userFixedRecord = model.UserFixedRecord{}
	if err := ss.DataTo(&userFixedRecord); err != nil {
		aelog.Errorf(ctx, "userFixedRecord parse error: %v", err)
		return nil, err
	}

	return &userFixedRecord, nil
}

// get all by fetched flag
func (r *UserFixedRecordRepository) GetAllByPictureFecthedFlag(ctx context.Context, uid *model.UID, isPictureFetched bool) ([]*model.UserFixedRecord, error) {
	dss, err := r.db.client.Collection("users").Doc(uid.ID).Collection("fixedRecords").Where("isPictureFetched", "==", isPictureFetched).Documents(ctx).GetAll()

	if err != nil {
		aelog.Errorf(ctx, "Failed userFixedRecord getAll: %v", err)
		return nil, err
	}

	userFixedRecords := make([]*model.UserFixedRecord, len(dss))
	for i, ss := range dss {
		var userFixedRecord = model.UserFixedRecord{}
		if err := ss.DataTo(&userFixedRecord); err != nil {
			aelog.Errorf(ctx, "userFixedRecord parse error: %v", err)
			return nil, err
		}
		userFixedRecords[i] = &userFixedRecord
	}

	return userFixedRecords, nil
}

// get All UserRecord
func (r *UserFixedRecordRepository) GetAll(ctx context.Context, uid *model.UID) ([]*model.UserFixedRecord, error) {
	dss, err := r.db.client.Collection("users").Doc(uid.ID).Collection("fixedRecords").Documents(ctx).GetAll()

	if err != nil {
		aelog.Errorf(ctx, "Failed userFixedRecord getAll: %v", err)
		return nil, err
	}

	userFixedRecords := make([]*model.UserFixedRecord, len(dss))
	for i, ss := range dss {
		var userFixedRecord = model.UserFixedRecord{}
		if err := ss.DataTo(&userFixedRecord); err != nil {
			aelog.Errorf(ctx, "userFixedRecord parse error: %v", err)
			return nil, err
		}
		userFixedRecords[i] = &userFixedRecord
	}

	return userFixedRecords, nil
}

func (r *UserFixedRecordRepository) Update(ctx context.Context, uid *model.UID, recordId string, userFixedRecordModel *model.UserFixedRecord) (*model.UserFixedRecord, error) {
	_, err := r.db.client.Collection("users").Doc(uid.ID).Collection("fixedRecords").Doc(recordId).Set(ctx, userFixedRecordModel)
	if err != nil {
		aelog.Errorf(ctx, "update fixedRecords failed: %v", err)
		return nil, err
	}
	return userFixedRecordModel, err
}
