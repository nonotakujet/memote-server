package viewmodel

import (
	"time"
)

type LocationViewModel struct {
	Lat string `json:"lat"`

	Long int32 `json:"long"`

	Time time.Time `json:"time"`
}
