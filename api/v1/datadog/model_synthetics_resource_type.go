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
// SyntheticsResourceType Document type to apply an assertion against.
type SyntheticsResourceType string

// List of SyntheticsResourceType
const (
	DOCUMENT SyntheticsResourceType = "document"
	STYLESHEET SyntheticsResourceType = "stylesheet"
	FETCH SyntheticsResourceType = "fetch"
	IMAGE SyntheticsResourceType = "image"
	SCRIPT SyntheticsResourceType = "script"
	XHR SyntheticsResourceType = "xhr"
	OTHER SyntheticsResourceType = "other"
)
