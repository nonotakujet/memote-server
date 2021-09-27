package model

import "time"

// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserRecordLocation struct {
	Latitude  float64   `firestore:"latitude,omitempty"`
	Longitude float64   `firestore:"longitude,omitempty"`
	Time      time.Time `firestore:"time,omitempty"`
}

// UserPosition holds about user position snapshot
// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserRecord struct {
	Id        string               `firestore:"id,omitempty"`
	Locations []UserRecordLocation `firestore:"locations,omitempty"`
	CreatedAt time.Time            `firestore:"createdAt,omitempty"`
}
