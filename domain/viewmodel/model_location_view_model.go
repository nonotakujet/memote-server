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

type LocationViewModel struct {

	Lat float64 `json:"lat"`

	Long float64 `json:"long"`

	Time time.Time `json:"time"`
}
