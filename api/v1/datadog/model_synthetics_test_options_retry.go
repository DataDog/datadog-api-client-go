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
// SyntheticsTestOptionsRetry Object describing the retry strategy to apply to a Synthetic test. By default, there is a 300ms wait before retrying a test that has failed.
type SyntheticsTestOptionsRetry struct {
	// The amount of location that needs to fail for the test to be retried.
	Count int64 `json:"count,omitempty"`
	// The interval over which the amount of location needed to fail for the test to be retried.
	Interval float64 `json:"interval,omitempty"`
}
