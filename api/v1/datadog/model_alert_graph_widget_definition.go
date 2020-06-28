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
// AlertGraphWidgetDefinition Alert graphs are timeseries graphs showing the current status of any monitor defined on your system.
type AlertGraphWidgetDefinition struct {
	// ID of the alert to use in the widget.
	AlertId string `json:"alert_id"`
	Time WidgetTime `json:"time,omitempty"`
	// The title of the widget.
	Title string `json:"title,omitempty"`
	TitleAlign WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize string `json:"title_size,omitempty"`
	Type AlertGraphWidgetDefinitionType `json:"type"`
	VizType WidgetVizType `json:"viz_type"`
}
