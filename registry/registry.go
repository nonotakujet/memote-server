package registry

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/nonotakujet/memote-server/domain/repository"
	"github.com/nonotakujet/memote-server/infra/persistence"
)

type Repository interface {
	NewUserPositionRepository() repository.UserPosition
}

type repositoryImpl struct {
	db                *persistence.DB
	userPositionoRepo repository.UserPosition
}

func NewRepository(ctx context.Context, client *firestore.Client) Repository {
	db := persistence.NewDB(client)
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) NewUserPositionRepository() repository.UserPosition {
	if r.userPositionoRepo == nil {
		r.userPositionoRepo = persistence.NewUserPositionRepository(r.db)
	}
	return r.userPositionoRepo
}
