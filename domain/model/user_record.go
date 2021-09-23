package model

import "time"

// UserPosition holds about user position snapshot
// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserRecord struct {
	Id        string         `firestore:"id,omitempty"`
	Locations []UserLocation `firestore:"locations,omitempty"`
	CreatedAt time.Time      `firestore:"createdAt,omitempty"`
}
