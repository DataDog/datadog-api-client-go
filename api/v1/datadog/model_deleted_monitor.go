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
// DeletedMonitor Response from the delete monitor call.
type DeletedMonitor struct {
	// ID of the deleted monitor.
	DeletedMonitorId int64 `json:"deleted_monitor_id,omitempty"`
}
