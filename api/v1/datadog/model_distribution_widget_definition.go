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
// DistributionWidgetDefinition The Distribution visualization is another way of showing metrics aggregated across one or several tags, such as hosts. Unlike the heat map, a distribution graph’s x-axis is quantity rather than time.
type DistributionWidgetDefinition struct {
	LegendSize WidgetLegendSize `json:"legend_size,omitempty"`
	// Array of one request object to display in the widget.  See the dedicated [Request JSON schema documentation](https://docs.datadoghq.com/dashboards/graphing_json/request_json)  to learn how to build the `REQUEST_SCHEMA`.
	Requests []DistributionWidgetRequest `json:"requests"`
	// Whether or not to display the legend on this widget.
	ShowLegend bool `json:"show_legend,omitempty"`
	Time WidgetTime `json:"time,omitempty"`
	// Title of the widget.
	Title string `json:"title,omitempty"`
	TitleAlign WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize string `json:"title_size,omitempty"`
	Type DistributionWidgetDefinitionType `json:"type"`
}
