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
// HeatMapWidgetRequest Updated heat map widget.
type HeatMapWidgetRequest struct {
	ApmQuery LogQueryDefinition `json:"apm_query,omitempty"`
	EventQuery EventQueryDefinition `json:"event_query,omitempty"`
	LogQuery LogQueryDefinition `json:"log_query,omitempty"`
	NetworkQuery LogQueryDefinition `json:"network_query,omitempty"`
	ProcessQuery ProcessQueryDefinition `json:"process_query,omitempty"`
	// Widget query.
	Q string `json:"q,omitempty"`
	RumQuery LogQueryDefinition `json:"rum_query,omitempty"`
	Style WidgetStyle `json:"style,omitempty"`
}
