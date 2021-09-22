package persistence

import (
	"context"
	"log"

	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

// UserPositionRepository holds user position inteface
type UserPositionRepository struct {
	db *DB
}

// NewUserPositionRepository new user position
func NewUserPositionRepository(db *DB) repository.UserPosition {
	newRepo := &UserPositionRepository{
		db: db,
	}
	return newRepo
}

// Create UserPosition
func (r *UserPositionRepository) Create(ctx context.Context, uid *model.UID, userPosition *model.UserPosition) (*model.UserPosition, error) {
	_, _, err := r.db.client.Collection("users").Doc(uid.ID).Collection("user_position").Add(ctx, userPosition)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return userPosition, err
}
