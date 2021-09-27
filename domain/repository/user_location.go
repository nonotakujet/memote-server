package repository

import (
	"context"

	"github.com/nonotakujet/memote-server/domain/model"
)

type UserLocation interface {
	GetNearBy(context.Context, *model.UID, string) ([]*model.UserLocation, error)
}
