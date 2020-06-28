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
// UsageNetworkHostsResponse Response containing the number of active NPM hosts for each hour for a given organization.
type UsageNetworkHostsResponse struct {
	// Get hourly usage for NPM hosts.
	Usage []UsageNetworkHostsHour `json:"usage,omitempty"`
}
