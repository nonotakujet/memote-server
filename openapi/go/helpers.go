/*
 * memote
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package viewmodel

//Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse {
		Code: code,
		Body: body,
	}
}
