package registry

import (
	"github.com/nonotakujet/memote-server/domain/repository"
	"github.com/nonotakujet/memote-server/infra/persistence"
)

type Repository interface {
	NewUserPositionRepository() repository.UserPosition
}

type repositoryImpl struct {
	userPositionoRepo repository.UserPosition
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) NewUserPositionRepository() repository.UserPosition {
	if r.userPositionoRepo == nil {
		r.userPositionoRepo = persistence.NewUserPositionRepository()
	}
	return r.userPositionoRepo
}
