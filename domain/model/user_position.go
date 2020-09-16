package model

import "time"

// UserPosition holds about user position snapshot
type UserPosition struct {
	Latitude  int64     `json: "latitude"`
	Longitude int64     `json: "longitude"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
}
