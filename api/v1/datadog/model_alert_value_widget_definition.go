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
// AlertValueWidgetDefinition Alert values are query values showing the current value of the metric in any monitor defined on your system.
type AlertValueWidgetDefinition struct {
	// ID of the alert to use in the widget.
	AlertId string `json:"alert_id"`
	// Number of decimal to show. If not defined, will use the raw value.
	Precision int64 `json:"precision,omitempty"`
	TextAlign WidgetTextAlign `json:"text_align,omitempty"`
	// Title of the widget.
	Title string `json:"title,omitempty"`
	TitleAlign WidgetTextAlign `json:"title_align,omitempty"`
	// Size of value in the widget.
	TitleSize string `json:"title_size,omitempty"`
	Type AlertValueWidgetDefinitionType `json:"type"`
	// Unit to display with the value.
	Unit string `json:"unit,omitempty"`
}
