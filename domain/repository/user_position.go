package repository

import "github.com/nonotakujet/memote-server/domain/model"

type UserPosition interface {
	Create(*model.UserPosition) (*model.UserPosition, error)
}
