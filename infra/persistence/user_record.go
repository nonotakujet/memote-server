package persistence

import (
	"context"

	"github.com/DeNA/aelog"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

// UserRecordRepository holds user position inteface
type UserRecordRepository struct {
	db *DB
}

// NewUserRecordRepository new user position
func NewUserRecordRepository(db *DB) repository.UserRecord {
	newRepo := &UserRecordRepository{
		db: db,
	}
	return newRepo
}

// Create UserRecord
func (r *UserRecordRepository) Create(ctx context.Context, uid *model.UID, userRecord *model.UserRecord) (*model.UserRecord, error) {
	_, err := r.db.client.Collection("users").Doc(uid.ID).Collection("records").Doc(userRecord.Id).Create(ctx, userRecord)
	if err != nil {
		aelog.Errorf(ctx, "create records failed: %v", err)
	}
	return userRecord, err
}
