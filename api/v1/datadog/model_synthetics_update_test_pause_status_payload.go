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
// SyntheticsUpdateTestPauseStatusPayload Object to start or pause an existing Synthetic test.
type SyntheticsUpdateTestPauseStatusPayload struct {
	NewStatus SyntheticsTestPauseStatus `json:"new_status,omitempty"`
}
