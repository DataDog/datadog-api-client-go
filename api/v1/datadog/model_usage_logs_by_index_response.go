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
// UsageLogsByIndexResponse Response containing the number of indexed logs for each hour and index for a given organization.
type UsageLogsByIndexResponse struct {
	// An array of objects regarding hourly usage of logs by index response.
	Usage []UsageLogsByIndexHour `json:"usage,omitempty"`
}
