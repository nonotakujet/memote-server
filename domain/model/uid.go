package model

import (
	"context"

	"github.com/pkg/errors"
)

const uidKey contextKey = "uid"

type contextKey string

type UID struct {
	ID string
}

func NewUID(uidString string) *UID {
	uid := &UID{
		ID: uidString,
	}
	return uid
}

func ContextWithUID(parent context.Context, uid *UID) context.Context {
	return context.WithValue(parent, uidKey, uid)
}

func UserFromContext(ctx context.Context) (*UID, error) {
	v := ctx.Value(uidKey)
	uid, ok := v.(*UID)
	if !ok {
		return nil, errors.New("uid not found")
	}
	return uid, nil
}
