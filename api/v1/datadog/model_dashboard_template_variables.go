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
// DashboardTemplateVariables Template variable.
type DashboardTemplateVariables struct {
	// The default value for the template variable on dashboard load.
	Default *string `json:"default,omitempty"`
	// The name of the variable.
	Name string `json:"name"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down.
	Prefix *string `json:"prefix,omitempty"`
}
