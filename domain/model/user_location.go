package model

// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserLocation struct {
	Geohash   string   `firestore:"geohash,omitempty"`
	Latitude  float64  `firestore:"latitude,omitempty"`
	Longitude float64  `firestore:"longitude,omitempty"`
	Name      string   `firestore:"name,omitempty"`
	RecordIds []string `firestore:"recordIds,omitempty"`
}
