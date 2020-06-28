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
// OrganizationListResponse Response with the list of organizations.
type OrganizationListResponse struct {
	// Array of organization objects.
	Orgs []Organization `json:"orgs,omitempty"`
}
