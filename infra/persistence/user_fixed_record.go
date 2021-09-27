package persistence

import (
	"context"
	"log"

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
		log.Fatalf("Failed userFixedRecord getById: %v", err)
		return nil, err
	}

	var userFixedRecord = model.UserFixedRecord{}
	if err := ss.DataTo(&userFixedRecord); err != nil {
		log.Fatalf("userFixedRecord parse error: %v", err)
		return nil, err
	}

	return &userFixedRecord, nil
}

// get All UserRecord
func (r *UserFixedRecordRepository) GetAll(ctx context.Context, uid *model.UID) ([]*model.UserFixedRecord, error) {
	dss, err := r.db.client.Collection("users").Doc(uid.ID).Collection("fixedRecords").Documents(ctx).GetAll()

	if err != nil {
		log.Fatalf("Failed userFixedRecord getAll: %v", err)
		return nil, err
	}

	userFixedRecords := make([]*model.UserFixedRecord, len(dss))
	for i, ss := range dss {
		var userFixedRecord = model.UserFixedRecord{}
		if err := ss.DataTo(&userFixedRecord); err != nil {
			log.Fatalf("userFixedRecord parse error: %v", err)
			return nil, err
		}
		userFixedRecords[i] = &userFixedRecord
	}

	return userFixedRecords, nil
}
