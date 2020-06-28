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
// LogsIndex Object describing a Datadog Log index.
type LogsIndex struct {
	// The number of log events you can send in this index per day before you are rate-limited.
	DailyLimit int64 `json:"daily_limit,omitempty"`
	// An array of exclusion objects. The logs are tested against the query of each filter, following the order of the array. Only the first matching active exclusion matters, others (if any) are ignored.
	ExclusionFilters []LogsExclusion `json:"exclusion_filters,omitempty"`
	Filter LogsFilter `json:"filter"`
	// A boolean stating if the index is rate limited, meaning more logs than the daily limit have been sent. Rate limit is reset every-day at 2pm UTC.
	IsRateLimited bool `json:"is_rate_limited,omitempty"`
	// The name of the index.
	Name string `json:"name,omitempty"`
	// The number of days before logs are deleted from this index.
	NumRetentionDays int64 `json:"num_retention_days,omitempty"`
}
