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
// UsageHostsResponse Host usage response.
type UsageHostsResponse struct {
	// An array of objects related to host usage.
	Usage []UsageHostHour `json:"usage,omitempty"`
}
