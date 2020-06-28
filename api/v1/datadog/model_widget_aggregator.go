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
// WidgetAggregator Aggregator used for the request.
type WidgetAggregator string

// List of WidgetAggregator
const (
	AVERAGE WidgetAggregator = "avg"
	LAST WidgetAggregator = "last"
	MAXIMUM WidgetAggregator = "max"
	MINIMUM WidgetAggregator = "min"
	SUM WidgetAggregator = "sum"
)
