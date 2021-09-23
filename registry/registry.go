package registry

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/nonotakujet/memote-server/domain/repository"
	"github.com/nonotakujet/memote-server/infra/persistence"
)

type Repository interface {
	NewUserPositionRepository() repository.UserPosition
	NewUserRecordRepository() repository.UserRecord
}

type repositoryImpl struct {
	db                *persistence.DB
	userPositionoRepo repository.UserPosition
	userRecordRepo    repository.UserRecord
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

func (r *repositoryImpl) NewUserRecordRepository() repository.UserRecord {
	if r.userRecordRepo == nil {
		r.userRecordRepo = persistence.NewUserRecordRepository(r.db)
	}
	return r.userRecordRepo
}
