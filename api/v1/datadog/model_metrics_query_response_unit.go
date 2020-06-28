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
// MetricsQueryResponseUnit Object containing the metric unit family, scale factor, name, and short name.
type MetricsQueryResponseUnit struct {
	// Unit family, allows for conversion between units of the same family, for scaling.
	Family string `json:"family,omitempty"`
	// Unit name
	Name string `json:"name,omitempty"`
	// Plural form of the unit name.
	Plural string `json:"plural,omitempty"`
	// Factor for scaling between units of the same family.
	ScaleFactor float64 `json:"scale_factor,omitempty"`
	// Abbreviation of the unit.
	ShortName string `json:"short_name,omitempty"`
}
