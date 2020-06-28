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
// SyntheticsAssertion Object describing the assertions type, their associated operator, which property they apply , and upon which target.
type SyntheticsAssertion struct {
	Operator SyntheticsAssertionOperator `json:"operator"`
	// The associated assertion property.
	Property string `json:"property,omitempty"`
	// Target to apply the assertion to. It can be a string, a number, or a Date.
	Target map[string]interface{} `json:"target,omitempty"`
	Type SyntheticsAssertionType `json:"type"`
}
