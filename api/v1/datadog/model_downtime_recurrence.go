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
// DowntimeRecurrence An object defining the recurrence of the downtime.
type DowntimeRecurrence struct {
	// How often to repeat as an integer. For example, to repeat every 3 days, select a type of `days` and a period of `3`.
	Period int32 `json:"period,omitempty"`
	// The type of recurrence. Choose from `days`, `weeks`, `months`, `years`.
	Type string `json:"type,omitempty"`
	// The date at which the recurrence should end as a POSIX timestamp. `until_occurences` and `until_date` are mutually exclusive.
	UntilDate *int64 `json:"until_date,omitempty"`
	// How many times the downtime is rescheduled. `until_occurences` and `until_date` are mutually exclusive.
	UntilOccurrences *int32 `json:"until_occurrences,omitempty"`
	// A list of week days to repeat on. Choose from `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat` or `Sun`. Only applicable when type is weeks. First letter must be capitalized.
	WeekDays []string `json:"week_days,omitempty"`
}
