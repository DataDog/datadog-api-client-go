/*
 * Datadog API Collection
 *
 * Collection of all Datadog Public endpoints.
 *
 * API version: 1.0
 * Contact: support@datadoghq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog
// CheckCanDeleteMonitorResponseData Wrapper object with the list of monitor IDs.
type CheckCanDeleteMonitorResponseData struct {
	// An array of of Monitor IDs that can be safely deleted.
	Ok []int64 `json:"ok,omitempty"`
}
