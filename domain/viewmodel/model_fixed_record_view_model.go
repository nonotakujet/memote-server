/*
 * memote
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package viewmodel

import (
	"time"
)

type FixedRecordViewModel struct {

	Id string `json:"id"`

	MainTitle string `json:"mainTitle"`

	MainPicture string `json:"mainPicture"`

	Locations []StayedLocationViewModel `json:"locations"`

	CreatedAt time.Time `json:"createdAt"`
}