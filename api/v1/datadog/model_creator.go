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
// Creator Object describing the creator of the shared element.
type Creator struct {
	// Email of the creator.
	Email string `json:"email,omitempty"`
	// Handle of the creator.
	Handle string `json:"handle,omitempty"`
	// Name of the creator.
	Name string `json:"name,omitempty"`
}
