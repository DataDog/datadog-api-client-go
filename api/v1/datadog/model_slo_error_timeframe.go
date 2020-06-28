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
// SloErrorTimeframe The timeframe of the threshold associated with this error or \"all\" if all thresholds are affected.
type SloErrorTimeframe string

// List of SLOErrorTimeframe
const (
	SEVEN_DAYS SloErrorTimeframe = "7d"
	THIRTY_DAYS SloErrorTimeframe = "30d"
	NINETY_DAYS SloErrorTimeframe = "90d"
	ALL SloErrorTimeframe = "all"
)
