package model

import "time"

// UserPosition holds about user position snapshot
// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserPosition struct {
	Latitude  int64     `firestore:"latitude,omitempty"`
	Longitude int64     `firestore:"longitude,omitempty"`
	CreatedAt time.Time `firestore:"createdAt,omitempty"`
	UpdatedAt time.Time `firestore:"updatedAt,omitempty"`
}
