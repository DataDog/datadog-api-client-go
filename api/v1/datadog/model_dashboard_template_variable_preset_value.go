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
// DashboardTemplateVariablePresetValue Template variables saved views.
type DashboardTemplateVariablePresetValue struct {
	// The name of the variable.
	Name string `json:"name,omitempty"`
	// The value of the template variable within the saved view.
	Value string `json:"value,omitempty"`
}
