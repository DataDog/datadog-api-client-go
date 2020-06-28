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
// SloListResponse A response with one or more service level objective.
type SloListResponse struct {
	// An array of service level objective objects.
	Data []ServiceLevelObjective `json:"data,omitempty"`
	// An array of error messages. Each endpoint documents how/whether this field is used.
	Errors []string `json:"errors,omitempty"`
}
