package repository

import (
	"context"

	"github.com/nonotakujet/memote-server/domain/model"
)

type UserFixedRecord interface {
	GetById(context.Context, *model.UID, string) (*model.UserFixedRecord, error)
	GetAll(context.Context, *model.UID) ([]*model.UserFixedRecord, error)
}
