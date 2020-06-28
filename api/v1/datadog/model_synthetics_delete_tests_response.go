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
// SyntheticsDeleteTestsResponse Response object for deleting Synthetic tests.
type SyntheticsDeleteTestsResponse struct {
	// Array of objects containing a deleted Synthetic test ID with the associated deletion timestamp.
	DeletedTests []SyntheticsDeleteTestsResponseDeletedTests `json:"deleted_tests,omitempty"`
}
