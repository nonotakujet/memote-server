package repository

import (
	"context"

	"github.com/nonotakujet/memote-server/domain/model"
)

type UserFixedRecord interface {
	GetById(context.Context, *model.UID, string) (*model.UserFixedRecord, error)
	GetAllByPictureFecthedFlag(context.Context, *model.UID, bool) ([]*model.UserFixedRecord, error)
	GetAll(context.Context, *model.UID) ([]*model.UserFixedRecord, error)
	Update(context.Context, *model.UID, string, *model.UserFixedRecord) (*model.UserFixedRecord, error)
}
