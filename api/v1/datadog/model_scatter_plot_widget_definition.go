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
// ScatterPlotWidgetDefinition The scatter plot visualization allows you to graph a chosen scope over two different metrics with their respective aggregation.
type ScatterPlotWidgetDefinition struct {
	// List of groups used for colors.
	ColorByGroups []string `json:"color_by_groups,omitempty"`
	Requests ScatterPlotWidgetDefinitionRequests `json:"requests"`
	Time WidgetTime `json:"time,omitempty"`
	// Title of your widget.
	Title string `json:"title,omitempty"`
	TitleAlign WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize string `json:"title_size,omitempty"`
	Type ScatterPlotWidgetDefinitionType `json:"type"`
	Xaxis WidgetAxis `json:"xaxis,omitempty"`
	Yaxis WidgetAxis `json:"yaxis,omitempty"`
}
