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
// SyntheticsTestsTriggerResult Object describing the results from triggering a Synthetic test.
type SyntheticsTestsTriggerResult struct {
	// Synthetic test trigger result id to poll.
	ResultId string `json:"result_id,omitempty"`
	// Public id of test that was triggered.
	PublicId string `json:"public_id,omitempty"`
	// Location of test result
	Location int `json:"location,omitempty"`
}
