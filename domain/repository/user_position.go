package repository

import (
	"context"

	"github.com/nonotakujet/memote-server/domain/model"
)

type UserPosition interface {
	Create(context.Context, *model.UserPosition) (*model.UserPosition, error)
}
