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
// SyntheticsTestPauseStatus Define whether you want to start (`live`) or pause (`paused`) a Synthetic test.
type SyntheticsTestPauseStatus string

// List of SyntheticsTestPauseStatus
const (
	LIVE SyntheticsTestPauseStatus = "live"
	PAUSED SyntheticsTestPauseStatus = "paused"
)
