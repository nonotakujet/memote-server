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

// A FixedRecordsApiController binds http requests to an api service and writes the service results to the http response
type FixedRecordsApiController struct {
	service FixedRecordsApiServicer
}

// NewFixedRecordsApiController creates a default api controller
func NewFixedRecordsApiController(s FixedRecordsApiServicer) Router {
	return &FixedRecordsApiController{service: s}
}

// Routes returns all of the api route for the FixedRecordsApiController
func (c *FixedRecordsApiController) Routes() Routes {
	return Routes{ 
		{
			"GetFixedRecord",
			strings.ToUpper("Get"),
			"/fixed_records/{recordId}",
			c.GetFixedRecord,
		},
		{
			"GetFixedRecords",
			strings.ToUpper("Get"),
			"/fixed_records",
			c.GetFixedRecords,
		},
		{
			"UpdateFixedRecord",
			strings.ToUpper("Put"),
			"/fixed_records/{recordId}",
			c.UpdateFixedRecord,
		},
	}
}

// GetFixedRecord - get fixed record
func (c *FixedRecordsApiController) GetFixedRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recordId := params["recordId"]
	
	result, err := c.service.GetFixedRecord(r.Context(), recordId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetFixedRecords - get fixed records
func (c *FixedRecordsApiController) GetFixedRecords(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	isFixed, err := parseBoolParameter(query.Get("isFixed"))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	result, err := c.service.GetFixedRecords(r.Context(), isFixed)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateFixedRecord - update fixed record
func (c *FixedRecordsApiController) UpdateFixedRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recordId := params["recordId"]
	
	fixedRecordViewModel := &FixedRecordViewModel{}
	if err := json.NewDecoder(r.Body).Decode(&fixedRecordViewModel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.UpdateFixedRecord(r.Context(), recordId, *fixedRecordViewModel)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
