package persistence

import (
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

// UserPositionRepository holds user position inteface
type UserPositionRepository struct {
}

// NewUserPositionRepository new user position
func NewUserPositionRepository() repository.UserPosition {
	newRepo := &UserPositionRepository{}
	return newRepo
}

// Create UserPosition
func (r *UserPositionRepository) Create(userPosition *model.UserPosition) (*model.UserPosition, error) {
	// todo : nonomura - implement db access.
	return userPosition, nil
}
