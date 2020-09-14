package model

import "time"

// UserPosition holds about user position snapshot
type UserPosition struct {
	Latitude  int64
	Longitude int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
