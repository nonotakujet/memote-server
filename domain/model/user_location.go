package model

import "time"

// UserPosition holds about user position snapshot
// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserLocation struct {
	Latitude  float64   `firestore:"latitude,omitempty"`
	Longitude float64   `firestore:"longitude,omitempty"`
	Time      time.Time `firestore:"time,omitempty"`
}
