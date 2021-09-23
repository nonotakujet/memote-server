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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A RecordsApiController binds http requests to an api service and writes the service results to the http response
type RecordsApiController struct {
	service RecordsApiServicer
}

// NewRecordsApiController creates a default api controller
func NewRecordsApiController(s RecordsApiServicer) Router {
	return &RecordsApiController{service: s}
}

// Routes returns all of the api route for the RecordsApiController
func (c *RecordsApiController) Routes() Routes {
	return Routes{ 
		{
			"PostRecords",
			strings.ToUpper("Post"),
			"/records",
			c.PostRecords,
		},
	}
}

// PostRecords - post records
func (c *RecordsApiController) PostRecords(w http.ResponseWriter, r *http.Request) {
	recordViewModel := &RecordViewModel{}
	if err := json.NewDecoder(r.Body).Decode(&recordViewModel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PostRecords(r.Context(), *recordViewModel)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}