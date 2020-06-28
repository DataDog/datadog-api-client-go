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
// LogsCategoryProcessorCategories Object describing the logs filter.
type LogsCategoryProcessorCategories struct {
	Filter LogsFilter `json:"filter,omitempty"`
	// Value to assign to the target attribute.
	Name string `json:"name,omitempty"`
}
