package model

import "time"

// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserFixedRecordLocation struct {
	Name      string    `firestore:"name,omitempty"`
	Latitude  float64   `firestore:"latitude,omitempty"`
	Longitude float64   `firestore:"longitude,omitempty"`
	Pictures  []string  `firestore:"pictures"`
	StartTime time.Time `firestore:"startTime,omitempty"`
	EndTime   time.Time `firestore:"endTime,omitempty"`
}

// https://firebase.google.cn/docs/firestore/manage-data/add-data?hl=ja#custom_objects
type UserFixedRecord struct {
	Id                string                    `firestore:"id,omitempty"`
	MainTitle         string                    `firestore:"mainTitle,omitempty"`
	MainPicture       string                    `firestore:"mainPicture,omitempty"`
	Locations         []UserFixedRecordLocation `firestore:"locations,omitempty"`
	IsPictureFetched  bool                      `firestore:"isPictureFetched,omitempty"`
	Message           string                    `firestore:"message,omitempty"`
	EmotionType       int32                     `firestore:"emotion_type,omitempty"`
	EmotionLevel      int32                     `firestore:"emotion_level,omitempty"`
	LastRecommendedAt time.Time                 `firestore:"lastRecommendedAt,omitempty"`
	CreatedAt         time.Time                 `firestore:"createdAt,omitempty"`
}
