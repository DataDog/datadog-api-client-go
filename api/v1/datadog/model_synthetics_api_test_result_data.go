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
// SyntheticsApiTestResultData Object containing results for your Synthetic API test.
type SyntheticsApiTestResultData struct {
	Cert SyntheticsSslCertificate `json:"cert,omitempty"`
	ErrorCode SyntheticsErrorCode `json:"errorCode,omitempty"`
	// The API test error message.
	ErrorMessage string `json:"errorMessage,omitempty"`
	EventType SyntheticsTestProcessStatus `json:"eventType,omitempty"`
	// The API test HTTP status code.
	HttpStatusCode int64 `json:"httpStatusCode,omitempty"`
	// Request header object used for the API test.
	RequestHeaders map[string]map[string]interface{} `json:"requestHeaders,omitempty"`
	// Response body returned for the API test.
	ResponseBody string `json:"responseBody,omitempty"`
	// Response headers returned for the API test.
	ResponseHeaders map[string]map[string]interface{} `json:"responseHeaders,omitempty"`
	// Global size in byte of the API test response.
	ResponseSize int64 `json:"responseSize,omitempty"`
	Timings SyntheticsTiming `json:"timings,omitempty"`
}
