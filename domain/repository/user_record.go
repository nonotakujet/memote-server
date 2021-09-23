package repository

import (
	"context"

	"github.com/nonotakujet/memote-server/domain/model"
)

type UserRecord interface {
	Create(context.Context, *model.UID, *model.UserRecord) (*model.UserRecord, error)
}
