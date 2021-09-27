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
	"context"
	"net/http"
)



// RecommendedRecordsApiRouter defines the required methods for binding the api requests to a responses for the RecommendedRecordsApi
// The RecommendedRecordsApiRouter implementation should parse necessary information from the http request, 
// pass the data to a RecommendedRecordsApiServicer to perform the required actions, then write the service results to the http response.
type RecommendedRecordsApiRouter interface { 
	GetRecommendedRecords(http.ResponseWriter, *http.Request)
}
// RecordsApiRouter defines the required methods for binding the api requests to a responses for the RecordsApi
// The RecordsApiRouter implementation should parse necessary information from the http request, 
// pass the data to a RecordsApiServicer to perform the required actions, then write the service results to the http response.
type RecordsApiRouter interface { 
	PostRecords(http.ResponseWriter, *http.Request)
}


// RecommendedRecordsApiServicer defines the api actions for the RecommendedRecordsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type RecommendedRecordsApiServicer interface { 
	GetRecommendedRecords(context.Context, float64, float64) (ImplResponse, error)
}


// RecordsApiServicer defines the api actions for the RecordsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type RecordsApiServicer interface { 
	PostRecords(context.Context, RecordViewModel) (ImplResponse, error)
}
