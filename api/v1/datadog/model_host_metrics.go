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
// HostMetrics Host Metrics collected.
type HostMetrics struct {
	// The percent of CPU used (everything but idle).
	Cpu float64 `json:"cpu,omitempty"`
	// The percent of CPU spent waiting on the IO (not reported for all platforms).
	Iowait float64 `json:"iowait,omitempty"`
	// The system load over the last 15 minutes.
	Load float64 `json:"load,omitempty"`
}
